package global

import (
	"backend/pkg/cache"
	"context"
	"database/sql"
	"log"
	"sync"
	"time"
)

var (
	// 任務註冊用
	JobSchedulerFuncCache cache.ILocalDataCache
	// 任務設定(from DB)
	// [cronId(int), *cron.Cron]
	JobSchedulerSetting sync.Map

	// 排程任務開關
	JobSchedulerSwitch bool
)

func calperformancereport(ctx context.Context, db *sql.DB, dateType string, startTime, endTime time.Time) {
	if startTime != endTime {
		// 計算盈虧報表
		errCode, tmps := CalPerformanceReportRecord(db, dateType, startTime, endTime)
		if errCode != 0 {
			return
		}

		dataCount := len(tmps)

		//如果有資料
		if dataCount > 0 {
			// 刪除原數據
			err := DelPerformanceReportRecord(db, dateType, startTime, endTime)
			if err != nil {
				return
			}
			// insert into 新數據
			err = InsertPerformanceReportRecord(db, dateType, tmps)
			if err != nil {
				return
			}
		}
	}
}

func updateJobSchedulerLastSyncDate(ctx context.Context, db *sql.DB, jobId, syncDate string) {
	query := `UPDATE job_scheduler
				SET last_sync_date = $1 , "update_time" = now()
				WHERE id = $2`
	result, err := db.Exec(query, syncDate, jobId)
	if err != nil {
		return
	}
	if count, err := result.RowsAffected(); count != 1 {
		log.Printf("updateJobSchedulerLastSyncDate: query exec failed, query = %v, err = %v", query, err)
	}
}

// func calGamereport(ctx context.Context, db *sql.DB, dateType string, startTime, endTime time.Time) {

// 	if startTime.Before(endTime) {
// 		// 計算盈虧報表 in db
// 		ReFreshMVGameStatRecord(ctx, db, dateType, startTime, endTime)
// 	}
// }
