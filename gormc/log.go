package gormc

import (
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormLog struct {
	Level                     logger.LogLevel
	IgnoreRecordNotFoundError bool
	SlowThreshold             time.Duration
}

func (g *GormLog) LogMode(level logger.LogLevel) logger.Interface {
	newLog := *g
	newLog.Level = level
	return &newLog
}

func (g *GormLog) Info(ctx context.Context, format string, data ...any) {
	if g.Level < logger.Info {
		return
	}
	logx.WithContext(ctx).Infof(format, data...)
}

func (g *GormLog) Warn(ctx context.Context, format string, data ...any) {
	if g.Level < logger.Warn {
		return
	}
	logx.WithContext(ctx).Errorf(format, data...)
}

func (g *GormLog) Error(ctx context.Context, format string, data ...any) {
	if g.Level < logger.Error {
		return
	}
	logx.WithContext(ctx).Errorf(format, data...)
}

func (g *GormLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.Level <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && g.Level >= logger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !g.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			logx.WithContext(ctx).WithDuration(elapsed).Errorw(
				err.Error(),
				logx.Field("sql", sql),
			)
		} else {
			logx.WithContext(ctx).WithDuration(elapsed).Errorw(
				err.Error(),
				logx.Field("sql", sql),
				logx.Field("rows", rows),
			)
		}
	case elapsed > g.SlowThreshold && g.SlowThreshold != 0 && g.Level >= logger.Warn:
		sql, rows := fc()
		if rows == -1 {
			logx.WithContext(ctx).WithDuration(elapsed).Sloww(
				"slow sql",
				logx.Field("sql", sql),
			)
		} else {
			logx.WithContext(ctx).WithDuration(elapsed).Sloww(
				"slow sql",
				logx.Field("sql", sql),
				logx.Field("rows", rows),
			)
		}
	case g.Level == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			logx.WithContext(ctx).WithDuration(elapsed).Infow(
				"",
				logx.Field("sql", sql),
			)
		} else {
			logx.WithContext(ctx).WithDuration(elapsed).Infow(
				"",
				logx.Field("sql", sql),
				logx.Field("rows", rows),
			)
		}
	}
}
