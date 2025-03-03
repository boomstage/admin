package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	env     string
	svcName string
)

const (
	EnvTypeLocal string = "local"
	EnvTypeTest  string = "test"
	EnvTypeProd  string = "prod"
)

func EnvInit() {
	initViper()
	env = viper.GetString("env")
	if env != EnvTypeLocal && env != EnvTypeTest && env != EnvTypeProd {
		panic(fmt.Errorf("invalid env: %s", env))
	}
	svcName = viper.GetString("svc_name")
	fmt.Printf("env.Init, env: %s, svc_name: %s\n", env, svcName)
}

func GetEnv() string {
	return env
}

func IsLocal() bool {
	return env == EnvTypeLocal
}

func IsTest() bool {
	return env == EnvTypeTest
}

func IsProd() bool {
	return env == EnvTypeProd
}

func GetSvcName() string {
	return svcName
}

func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func initViper() {
	if viper.ConfigFileUsed() != "" {
		return
	}
	var confFile string
	flag.StringVar(&confFile, "config", "./config.yaml", "config file")
	flag.Parse()
	viper.SetConfigFile(confFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
