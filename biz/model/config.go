package model

type Config struct {
	Port       int64       `mapstructure:"port"`
	Env        string      `mapstructure:"env"`
	MetricPort int64       `mapstructure:"metric_port"`
	MySQL      *MySQLConf  `mapstructure:"mysql"`
	Redis      *RedisConf  `mapstructure:"redis"`
	Twilio     *TwilioConf `mapstructure:"twilio"`
	JWT        *JWTConf    `mapstructure:"jwt"`
	AWS        *AWSConf    `mapstructure:"aws"`
	File       *FileConf   `mapstructure:"file"`
	Lark       *LarkConf   `mapstructure:"lark"`
}

type AWSConf struct {
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	Bucket          string `mapstructure:"bucket"`
}

type MySQLConf struct {
	Base struct {
		MasterDSN string `mapstructure:"master_dsn"`
		SlaveDSN  string `mapstructure:"slave_dsn"`
	} `mapstructure:"base"`
	Admin struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"admin"`
	DataCenter struct {
		MasterDSN string `mapstructure:"master_dsn"`
		SlaveDSN  string `mapstructure:"slave_dsn"`
	} `mapstructure:"datacenter"`
}

type LarkConf struct {
	AppId     string `mapstructure:"app_id"`
	AppSecret string `mapstructure:"app_secret"`
}

type RedisConf struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type TwilioConf struct {
	AccountSid string `mapstructure:"account_sid"`
	AuthToken  string `mapstructure:"auth_token"`
}

type JWTConf struct {
	Secrets map[UserSource]string `mapstructure:"secrets"`
}
