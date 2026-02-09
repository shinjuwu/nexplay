package global

import (
	"backend/pkg/redis"
	"backend/pkg/utils"
	"context"
	"database/sql"
	"log"
	"time"
)

/*
業績報表 - 代理時段統計資料(record profit)
*/

/* rp_agent_stat_15min
每小時的 0, 15, 30, 45 分執行
*/
func jobRPAgentStat15Minute(ctx context.Context, db *sql.DB, rdb redis.IRedisCliect, args []string) {
	dateType := "15min"
	jobId := args[0]
	lastSyncDate := args[1]
	if lastSyncDate == "" {
		now := time.Now().UTC()
		startTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()-15, 0, 0, now.UTC().Location())
		endTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.UTC().Location())
		log.Printf("jobRPAgentStat15Minute start at %v, endTime at %v", startTime, endTime)
		calperformancereport(ctx, db, dateType, startTime, endTime)
		updateJobSchedulerLastSyncDate(ctx, db, jobId, utils.GetUnsignedTimeUTC(startTime, dateType))
	} else {
		lastSyncTime, err := utils.GetUnsignedTimeUTCFromStr(lastSyncDate, dateType)
		if err != nil {
			return
		}
		now := time.Now().UTC()
		startTime := time.Date(lastSyncTime.Year(), lastSyncTime.Month(), lastSyncTime.Day(), lastSyncTime.Hour(), lastSyncTime.Minute()+15, 0, 0, now.UTC().Location())
		recordTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute()-15, 0, 0, now.UTC().Location())
		endTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.UTC().Location())
		log.Printf("jobRPAgentStat15Minute start at %v, endTime at %v", startTime, endTime)
		calperformancereport(ctx, db, dateType, startTime, endTime)
		updateJobSchedulerLastSyncDate(ctx, db, jobId, utils.GetUnsignedTimeUTC(recordTime, dateType))
	}
}
