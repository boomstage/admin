package service

import (
	"github.com/go-co-op/gocron/v2"
)

type CronSvc struct {
	// nothing
}

func InitCron() *CronSvc {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	// dashboard 实时用户数据
	//if _, err = s.NewJob(
	//	gocron.CronJob("*/5 * * * *", false), // every 5 minutes
	//	gocron.NewTask(Dashboard.RealtimeDataCron, context.TODO()),
	//); err != nil {
	//	panic(err)
	//}
	s.Start()
	return &CronSvc{}
}
