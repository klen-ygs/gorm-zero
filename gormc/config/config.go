package config

import (
	"log"
	"os"
	"time"

	logx "github.com/klen/gorm-zero/gormc"

	"gorm.io/gorm/logger"
)

type GormLogConfigI interface {
	GetGormLogMode() logger.LogLevel
	GetSlowThreshold() time.Duration
	GetColorful() bool
}

func NewLogxGormLogger(cfg GormLogConfigI) logger.Interface {
	return &logx.GormLog{
		Level:                     cfg.GetGormLogMode(),
		IgnoreRecordNotFoundError: true,
		SlowThreshold:             cfg.GetSlowThreshold(),
	}
}

func NewDefaultGormLogger(cfg GormLogConfigI) logger.Interface {
	newLogger := logger.New(
		log.New(os.Stderr, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             cfg.GetSlowThreshold(), // 慢 SQL 阈值
			LogLevel:                  cfg.GetGormLogMode(),   // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  cfg.GetColorful(),      // 禁用彩色打印
		},
	)
	return newLogger
}

func OverwriteGormLogMode(mode string) logger.LogLevel {
	switch mode {
	case "dev":
		return logger.Info
	case "test":
		return logger.Warn
	case "prod":
		return logger.Error
	case "silent":
		return logger.Silent
	default:
		return logger.Info
	}
}
