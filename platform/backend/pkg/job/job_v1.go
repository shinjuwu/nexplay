package job

import (
	"backend/pkg/redis"
	"context"
	"database/sql"
	"log"
	"sync"

	"github.com/robfig/cron/v3"
)

/*
base on robfig/cron
*/

type jobV1 struct {
	ctx          context.Context
	db           *sql.DB
	rdb          redis.IRedisCliect
	id           string // uuid from db
	cronId       int    // create from after job start
	spec         string
	info         string
	lastSyncDate string
	execCount    int // 目前已執行次數
	execLimit    int // 最大可執行次數
	isEnabled    bool
	jobCron      cron.Cron //*cron.New(cron.WithSeconds())
	eventType    int
	callback     func(context.Context, *sql.DB, redis.IRedisCliect, []string)
	mu           *sync.Mutex
}

/* 創建一個自定義的 cron job 物件
必須要自己創建一個 cron.Cron 才能達到關關單一 job 的功能
*/
func newjobV1(ctx context.Context, clib cron.Cron, db *sql.DB, rdb redis.IRedisCliect, id, spec, info, lastSyncDate string, isEnabled bool, execLimit int, callback func(context.Context, *sql.DB, redis.IRedisCliect, []string)) (IJob, error) {

	job := &jobV1{
		ctx:          ctx,
		db:           db,
		rdb:          rdb,
		id:           id,
		cronId:       -1, // default
		spec:         spec,
		info:         info,
		lastSyncDate: lastSyncDate,
		execCount:    0,
		execLimit:    execLimit,
		isEnabled:    isEnabled,
		eventType:    -1,
		jobCron:      clib,
		callback:     callback,
		mu:           new(sync.Mutex),
	}

	entryId, err := job.jobCron.AddFunc(job.spec, func() {
		job.mu.Lock()

		err := db.QueryRowContext(ctx, `SELECT last_sync_date FROM job_scheduler WHERE id=$1`, job.id).Scan(&job.lastSyncDate)
		if err != nil {
			job.lastSyncDate = ""
		}

		args := []string{}
		args = append(args, job.id)
		args = append(args, job.lastSyncDate)

		if callback != nil {
			callback(job.ctx, job.db, job.rdb, args)
		}

		// execCount, execLimit default is 0
		// 設定 execLimit 為0,  之後永遠不會執行停止
		if job.execLimit > 0 {
			job.execCount++
			if job.execCount == job.execLimit {
				job.Quit()
			}
		}
		job.mu.Unlock()
	})
	job.eventType = EventType_JobSchduleActiveCallback
	job.cronId = int(entryId)

	return job, err
}

func (p *jobV1) Id() string {
	return p.id
}

func (p *jobV1) Spec() string {
	return p.spec
}

func (p *jobV1) Info() string {
	return p.info
}

func (p *jobV1) LastSyncDate() string {
	return p.lastSyncDate
}

func (p *jobV1) ExecCount() int {
	return p.execCount
}

func (p *jobV1) ExecLimit() int {
	return p.execLimit
}

func (p *jobV1) IsEnabled() bool {
	return p.isEnabled
}

func (p *jobV1) Run() {
	log.Printf("Job_runtime Run() job id is: %s,cron spec: %s, cron id: %d, cron info: %s", p.id, p.spec, p.cronId, p.info)
	p.jobCron.Start()
}

func (p *jobV1) Quit() {
	log.Printf("Job_runtime Quit() job id is: %s,cron spec: %s, cron id: %d, cron info: %s", p.id, p.spec, p.cronId, p.info)
	p.jobCron.Stop()
}
