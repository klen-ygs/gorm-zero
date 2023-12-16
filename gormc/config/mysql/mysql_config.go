package mysql

import (
	"errors"
	"fmt"
	"time"

	"github.com/klen-ygs/gorm-zero/gormc"
	"github.com/klen-ygs/gorm-zero/gormc/config"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Conf struct {
	Path          string // 服务器地址
	Port          int    `json:",default=3306"` // 端口
	Config        string `json:",default=\"\""` // default: charset=utf8mb4&parseTime=True&loc=Local 高级配置
	Dbname        string // 数据库名
	Username      string // 数据库用户名
	Password      string // 数据库密码
	MaxIdleConns  int    `json:",default=10"`                               // 空闲中的最大连接数
	MaxOpenConns  int    `json:",default=10"`                               // 打开到数据库的最大连接数
	LogMode       string `json:",default=dev,options=dev|test|prod|silent"` // 是否开启Gorm全局日志
	SlowThreshold int64  `json:",default=1000"`
}

func (m *Conf) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + fmt.Sprintf("%d", m.Port) + ")/" + m.Dbname + "?" + m.GetConnConfig()
}

func (m *Conf) GetGormLogMode() logger.LogLevel {
	return config.OverwriteGormLogMode(m.LogMode)
}

func (m *Conf) GetSlowThreshold() time.Duration {
	return time.Duration(m.SlowThreshold) * time.Millisecond
}
func (m *Conf) GetColorful() bool {
	return true
}

func (m *Conf) GetConnConfig() string {
	if m.Config == "" {
		return "charset=utf8mb4&parseTime=True&loc=Local"
	}
	return m.Config
}

func Connect(m Conf) (*gorm.DB, error) {
	if m.Dbname == "" {
		return nil, errors.New("database name is empty")
	}
	mysqlCfg := mysql.Config{
		DSN: m.Dsn(),
	}
	newLogger := config.NewLogxGormLogger(&m)
	db, err := gorm.Open(mysql.New(mysqlCfg), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	err = db.Use(gormc.OtelPlugin{})
	if err != nil {
		return nil, err
	}

	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(m.MaxIdleConns)
	sqldb.SetMaxOpenConns(m.MaxOpenConns)
	return db, nil

}

func MustConnect(m Conf) *gorm.DB {
	if m.Dbname == "" {
		logx.Must(errors.New("database name is empty"))
	}
	mysqlCfg := mysql.Config{
		DSN: m.Dsn(),
	}
	newLogger := config.NewLogxGormLogger(&m)
	db, err := gorm.Open(mysql.New(mysqlCfg), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logx.Must(err)
	}

	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(m.MaxIdleConns)
	sqldb.SetMaxOpenConns(m.MaxOpenConns)

	return db
}

func ConnectWithConfig(m Conf, cfg *gorm.Config) (*gorm.DB, error) {
	if m.Dbname == "" {
		return nil, errors.New("database name is empty")
	}
	mysqlCfg := mysql.Config{
		DSN: m.Dsn(),
	}
	db, err := gorm.Open(mysql.New(mysqlCfg), cfg)
	if err != nil {
		return nil, err
	}
	err = db.Use(gormc.OtelPlugin{})
	if err != nil {
		return nil, err
	}

	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(m.MaxIdleConns)
	sqldb.SetMaxOpenConns(m.MaxOpenConns)
	return db, nil

}

func MustConnectWithConfig(m Conf, cfg *gorm.Config) *gorm.DB {
	if m.Dbname == "" {
		logx.Must(errors.New("database name is empty"))
	}
	mysqlCfg := mysql.Config{
		DSN: m.Dsn(),
	}
	db, err := gorm.Open(mysql.New(mysqlCfg), cfg)
	if err != nil {
		logx.Must(err)
	}

	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(m.MaxIdleConns)
	sqldb.SetMaxOpenConns(m.MaxOpenConns)

	return db
}
