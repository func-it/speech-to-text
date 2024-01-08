package logi

import (
	context "context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

func GetGormLogger() gormlogger.Interface {
	return current.GetGormLogger()
}

type zapGorm struct {
	*zap.Logger
	level       zapcore.Level
	enumVerbose EnumVerbose
}

func (l *zapGorm) Info(ctx context.Context, s string, i ...interface{}) {
	l.Logger.Info(fmt.Sprintf(strings.ReplaceAll(s, "\n", ""), i...))
}

func (l *zapGorm) Warn(ctx context.Context, s string, i ...interface{}) {
	l.Logger.Warn(fmt.Sprintf(strings.ReplaceAll(s, "\n", ""), i...))
}

func (l *zapGorm) Error(ctx context.Context, s string, i ...interface{}) {
	l.Logger.Error(fmt.Sprintf(strings.ReplaceAll(s, "\n", ""), i...))
}

func (l *zapGorm) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	slowThreshold := 200 * time.Millisecond
	IgnoreRecordNotFoundError := false
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.level >= zap.ErrorLevel && (!errors.Is(err, gormlogger.ErrRecordNotFound) || !IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Logger.Error("SQL error: " + sql)
		} else {
			l.Logger.Error("SQL error"+sql, zap.Int64("rows_affected", rows))
		}
	case elapsed > slowThreshold && slowThreshold != 0 && l.level <= zap.WarnLevel:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		if rows == -1 {
			l.Logger.Warn("SQL slow: "+sql, zap.String("slow", slowLog))
		} else {
			l.Logger.Warn("SQL slow: "+sql, zap.String("slow", slowLog), zap.Int64("rows_affected", rows))
		}
	case l.level >= zap.InfoLevel:
		sql, rows := fc()
		if rows == -1 {
			l.Logger.Info("SQL success: " + sql)
		} else {
			l.Logger.Info("SQL success: "+sql, zap.Int64("rows_affected", rows))
		}
	}
}

func (l *zapGorm) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return l
}
