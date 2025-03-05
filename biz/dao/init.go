package dao

import (
	"fmt"
	"github.com/boomstage/admin/biz/model"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBM *gorm.DB
	DBS *gorm.DB

	// RedisClt *redis.Client
	RedisClt *redis.ClusterClient
)

var Conf *model.Config

var (
	User *UserDao
)

func Init() {
	InitConf()
	initMySQLClients()

	User = InitUser()

}

func GetDB(options *Options) *gorm.DB {
	if options.FromMaster {
		return DBM
	}
	return DBS
}

func InitConf() {
	var conf *model.Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	Conf = conf
}

func initMySQLClients() {
	var err error
	DBM, err = gorm.Open(mysql.Open(Conf.MySQL.Base.MasterDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
