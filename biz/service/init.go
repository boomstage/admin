package service

import (
	"math/rand"
	"time"
)

var (
	User *UserSvc

	Cron *CronSvc
)

func Init() {
	rand.Seed(time.Now().UnixNano())

	User = InitUser()

	Cron = InitCron() // 最后执行
}
