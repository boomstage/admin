package util

import (
	"log"
	"os"
	"path"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// don't use hlog, use Zerolog
var Zerolog zerolog.Logger

func InitLog() (zerolog.Logger, error) {
	// 可定制的输出目录。
	logFilePath := "./logs/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Println(err.Error())
		return Zerolog, err
	}

	// 将文件名设置为日期
	logFileName := "server.log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return Zerolog, err
		}
	}

	// 提供压缩和删除
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    100, // 一个文件最大可达多少M。
		MaxBackups: 16,  // 最多同时保存 多少 个文件。
		MaxAge:     10,  // 一个文件最多可以保存 多少 天。
		// Compress:   true, // 用 gzip 压缩。
	}

	// consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	// zerolog.MultiLevelWriter(consoleWriter, lumberjackLogger)

	logger := hertzZerolog.New(
		hertzZerolog.WithOutput(lumberjackLogger),
		hertzZerolog.WithLevel(hlog.LevelInfo),
		hertzZerolog.WithTimestamp(),
		hertzZerolog.WithCaller(), // don't use hlog, use Zerolog
	)

	if IsDisableFileLog() {
		logger.SetOutput(zerolog.ConsoleWriter{Out: os.Stdout})
	} else {
		logger.SetOutput(lumberjackLogger)
	}

	hlog.SetLogger(logger)

	Zerolog = logger.Unwrap()
	return Zerolog, nil
}

func IsDisableFileLog() bool {
	return viper.GetString("env") == "local"
}
