package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
)

type Config struct {
	Level       string
	Format      string
	OutputPaths []string
	EnableColor bool
}

var zlogger zerolog.Logger

func Init(conf *Config) error {
	// 设置日志级别
	level, err := zerolog.ParseLevel(conf.Level)
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(level)

	// 创建多个写入器
	var writers []io.Writer
	for _, path := range conf.OutputPaths {
		if "stdout" == path {
			writers = append(writers, os.Stdout)
		} else {

			logFile := &lumberjack.Logger{
				Filename:   path, // 日志文件路径
				MaxSize:    5,    // 每个日志文件的最大大小（MB）
				MaxBackups: 5,    // 保留旧日志文件的最大数量
				MaxAge:     30,   // 保留旧日志文件的最大天数
				Compress:   true, // 是否压缩旧日志文件
			}
			defer logFile.Close()
			writers = append(writers, logFile)
		}
	}

	// 设置输出格式
	switch strings.ToLower(conf.Format) {
	case "text":
		consoleWriter := zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "2006-01-02 15:04:05",
			NoColor:    !conf.EnableColor,
		}
		writers = append([]io.Writer{consoleWriter}, writers...)
	}
	writers = append([]io.Writer{zerolog.ConsoleWriter{Out: os.Stderr}}, writers...)

	// 创建全局日志器
	multiWriter := zerolog.MultiLevelWriter(writers...)
	zlogger = zerolog.New(multiWriter).With().Timestamp().Logger()

	// 替换标准库log
	log.Logger = zlogger

	return nil
}
