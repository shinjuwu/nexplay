package job

import (
	"backend/pkg/redis"
	"context"
	"database/sql"

	"github.com/rfyiamcool/cronlib"
)

/*
1. 創建 job 並指定工作
2. 修改、刪除 job
3. 監控已創建之 job
*/

type IJob interface {
	Id() string
	Spec() string
	Info() string
	LastSyncDate() string
	ExecCount() int
	ExecLimit() int
	IsEnabled() bool
	Run()
	Quit()
}

// type JobSchduleActiveCallback func(*sql.DB, []string)

const (
	EventType_JobSchduleActiveCallback = iota
)

func newIJob(ctx context.Context, clib *cronlib.CronSchduler, db *sql.DB, rdb redis.IRedisCliect, id, spec, info, lastSyncDate string,
	isEnabled bool, execLimit int, callback func(context.Context, *sql.DB, redis.IRedisCliect, []string)) (IJob, error) {

	return newjobV2(ctx, clib, db, rdb, id, spec, info, lastSyncDate, isEnabled, execLimit, callback)
}
