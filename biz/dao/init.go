package dao

import (
	"fmt"
	"strings"
	"time"

	"github.com/boomstage/admin/biz/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	DBM           *sqlx.DB
	DBS           *sqlx.DB
	DBAdm         *sqlx.DB
	DataCenterDBM *sqlx.DB
	DataCenterDBS *sqlx.DB

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

func GetDB(options *Options) *sqlx.DB {
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
	DBM = sqlx.MustConnect("mysql", Conf.MySQL.Base.MasterDSN)
	setDBClient(DBM)
}

func setDBClient(cl *sqlx.DB) {
	cl.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	cl.SetMaxOpenConns(128)
	cl.SetMaxIdleConns(128)
	cl.SetConnMaxLifetime(20 * time.Minute)
}
