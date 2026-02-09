package job_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/rfyiamcool/cronlib"
	"github.com/robfig/cron/v3"
)

func TestCron(t *testing.T) {
	c := cron.New(cron.WithSeconds())
	i := 1
	c.AddFunc("*/1 * * * * *", func() {
		fmt.Println("exec every second", i)
		i++
	})
	c.Start()
	log.Println(len(c.Entries()))
	time.Sleep(time.Minute * 5)

	// log.Println(c.Entries())
}

// func TestCronManager(t *testing.T) {

// 	tmp := job.NewJobScheduler()

// 	tmp.CreateJob("", "0 49 * * * */4", "test1", true, nilfunc)

// 	err := tmp.Start()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	time.Sleep(time.Minute * 5)

// 	// log.Println(c.Entries())
// }

func TestCronlib(t *testing.T) {

	cron := cronlib.New()
	job, err := cronlib.NewJobModel(
		"*/1 * * * * *",
		func() {
			log.Println(123456)
		},
	)
	if err != nil {
		panic(err.Error())
	}

	err = cron.Register("123", job)
	if err != nil {
		panic(err.Error())
	}

	cron.Start()
	cron.Wait()
}
