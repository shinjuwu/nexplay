package job

import (
	"backend/pkg/cache"
	"backend/pkg/redis"
	"context"
	"database/sql"
	"log"

	"github.com/rfyiamcool/cronlib"
)

/*
1. 創建 job 並指定工作
2. 修改、刪除 job
3. 監控已創建之 job
*/

// type JobTriggerFunc func(*sql.DB, redis.IRedisCliect, []string)
// type JobAfterTriggerFunc func(*sql.DB, string, string)

type JobScheduler struct {
	clib    *cronlib.CronSchduler
	joblist cache.ILocalDataCache // [int, IJobScheduler]
}

func NewJobScheduler() *JobScheduler {
	return &JobScheduler{
		clib:    cronlib.New(),
		joblist: cache.NewLocalDataCache(),
	}
}

func (p *JobScheduler) CreateJob(ctx context.Context, db *sql.DB, rdb redis.IRedisCliect, id, spec, info, lastSyncDate string,
	isEnabled bool, execLimit int, callback func(context.Context, *sql.DB, redis.IRedisCliect, []string)) (string, error) {
	job, err := newIJob(ctx, p.clib, db, rdb, id, spec, info, lastSyncDate, isEnabled, execLimit, callback)

	p.joblist.Add(job.Id(), job)

	return job.Id(), err
}

// strat all event already setting
func (p *JobScheduler) Start() {
	p.clib.Start()
	// print job info
	p.joblist.GetInstance().Range(func(key, value any) bool {
		temp := value.(IJob)
		log.Printf("JobScheduler Start() job id is: %s,cron spec: %s, cron info: %s", temp.Id(), temp.Spec(), temp.Info())
		return true
	})

	// p.clib.Wait()
}

// strat all event already setting
func (p *JobScheduler) Stop() error {
	tmps := p.joblist.GetInstance()

	tmps.Range(func(k, v interface{}) bool {
		// key := k.(int)  // cron id
		val := v.(IJob) // cron interface object
		val.Quit()

		// log.Printf("Stop() job id: %d, cron info: %s", key, val.Info())
		return true
	})

	return nil
}

func (p *JobScheduler) Lists() []map[string]interface{} {
	tmps := p.joblist.GetInstance()

	rts := make([]map[string]interface{}, 0)
	tmps.Range(func(k, v interface{}) bool {
		val := v.(IJob) // cron interface object
		tmp := make(map[string]interface{}, 0)
		tmp["id"] = val.Id()
		tmp["spec"] = val.Spec()
		tmp["info"] = val.Info()
		tmp["exec_count"] = val.ExecCount()
		tmp["exec_limit"] = val.ExecLimit()
		tmp["last_sync_date"] = val.LastSyncDate()
		tmp["is_enabled"] = val.IsEnabled()
		rts = append(rts, tmp)
		return true
	})
	return rts
}
