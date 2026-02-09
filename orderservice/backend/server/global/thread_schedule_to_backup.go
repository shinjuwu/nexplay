package global

import (
	"backend/pkg/database"
	"backend/pkg/logger"
	"fmt"
	"log"
	"sync"
	"time"

	table_model "backend/server/table/model"
)

/*
	1. 備份表名、備份時間長短，備份開關皆在DB中設定
	2. 備份時間固定在每天的凌晨1點(everyday at 1:00 am)
	3. 備份時間以天為單位、加入時區(時區填0或-1效果一樣)
	4. 每天檢查一次，依序執行
*/

func nextBackupTime(day int, hour int) time.Time {
	now := time.Now().UTC()
	nextBackup := now.AddDate(0, 0, day)
	nextBackup = time.Date(nextBackup.Year(), nextBackup.Month(), nextBackup.Day(), hour, 0,
		0, 0, nextBackup.Location())

	return nextBackup
}

// dataKeepingDay: 7 (7天以外備份)
// execBackupHour: 1 (凌晨1點)
func T_execScheduleToBackup(logger *logger.RuntimeGoLogger, idb database.IDatabase, ScheduleToBackups []*table_model.ScheduleToBackup) {

	// 執行次數
	execCount := len(ScheduleToBackups)

	logger.Printf("T_execScheduleToBackup has %d process", execCount)
	nextBackup := nextBackupTime(1, 0)
	now := time.Now().UTC()

	timeOffset := nextBackup.Sub(now)
	// timeOffset := time.Until(nextBackupTime(0, 0))
	c := time.NewTimer(timeOffset)
	for {
		<-c.C

		var wg sync.WaitGroup
		taskCount := len(ScheduleToBackups)
		wg.Add(taskCount)

		taskCh := make(chan struct{}, 1)

		for _, v := range ScheduleToBackups {
			if v.IsEnabled {
				go _execScheduleToBackup(logger, idb, &wg, taskCh, v.Id, v.DataKeepingTable, v.DataKeepingDay)
			}
		}

		// start to precess task
		taskCh <- struct{}{}

		wg.Wait()

		logger.Println("所有任務已完成")
	}
}

func _execScheduleToBackup(logger *logger.RuntimeGoLogger, idb database.IDatabase, wg *sync.WaitGroup, taskCh chan struct{}, taskId string, dataKeepingTable string, dataKeepingDay int) {

	defer wg.Done()

	<-taskCh

	logger.Printf("backup start, taskId: %s", taskId)

	db := idb.GetDB(DB_IDX_ORDER)

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// 備份資料表
	backupDataKeepingTable := dataKeepingTable + "_backup"

	// 計算7天前的時間
	sevenDaysAgo := nextBackupTime(-dataKeepingDay, 0)

	// 創建備份表（僅在表不存在時創建）
	createTableQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (LIKE %s INCLUDING CONSTRAINTS)", backupDataKeepingTable, dataKeepingTable)
	_, err = tx.Exec(createTableQuery)
	if err != nil {
		tx.Rollback()
		logger.Printf("createTableQuery: %s, CREATE TABLE has error: %v", createTableQuery, err)
	}

	insertQuery := fmt.Sprintf(
		`INSERT INTO %s
		SELECT *
		FROM %s
		WHERE bet_time < $1;`,
		backupDataKeepingTable, dataKeepingTable)

	_, err = tx.Exec(insertQuery, sevenDaysAgo)
	if err != nil {
		tx.Rollback()
		logger.Printf("insertQuery: %s, INSERT INTO has error: %v", insertQuery, err)
	}

	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE bet_time < $1", dataKeepingTable)

	_, err = tx.Exec(deleteQuery, sevenDaysAgo)
	if err != nil {
		tx.Rollback()
		logger.Printf("deleteQuery: %s, DELETE has error: %v", deleteQuery, err)
	}

	err = tx.Commit()
	if err != nil {
		logger.Printf("Commit() has error: %v", err)
	}

	logger.Printf("backup completed, taskId: %s", taskId)

	// update last exec time
	updateLastExecTimeQuery := `UPDATE schedule_to_backup SET last_exec_time=now() WHERE id=$1;`
	_, err = db.Exec(updateLastExecTimeQuery, taskId)
	if err != nil {
		logger.Info("updateLastExecTimeQuery exec has error, err: %v", err)
	}

	if len(taskCh) == 0 {
		taskCh <- struct{}{}
	}
}
