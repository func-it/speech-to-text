package logi

import (
	"context"
	"errors"
	"fmt"
	runtime "runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	watsapplog "go.mau.fi/whatsmeow/util/log"
	gormlogger "gorm.io/gorm/logger"
)

type logrusLogger struct {
	logrus *logrus.Logger
}

func (l *logrusLogger) Fatalf(msg string, args ...any) {
	l.logrus.Fatalf(msg, args...)
}

func (l *logrusLogger) Panicf(msg string, args ...any) {
	l.logrus.Panicf(msg, args...)
}

func (l *logrusLogger) Warnf(msg string, args ...any) {
	l.logrus.Warnf(msg, args...)
}

func (l *logrusLogger) Errorf(msg string, args ...any) {
	l.logrus.Errorf(msg, args...)
}

func (l *logrusLogger) Infof(msg string, args ...any) {
	l.logrus.Infof(msg, args...)
}

func (l *logrusLogger) Debugf(msg string, args ...any) {
	l.logrus.Debugf(msg, args...)
}

func (l *logrusLogger) Debug(msg string, fields ...Field) {
	l.logrus.Debugf(msg)
}

func NewLogrus(level EnumVerbose) Log {
	return newLogrus(level)
}

func newLogrus(level EnumVerbose) *logrusLogger {
	lg := logrus.New()
	lg.Level = level.ToLogrusLevel()

	lg.SetReportCaller(true)
	lg.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           time.RFC1123,
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			function = ""
			arr := strings.Split(frame.File, "/")
			if len(arr) > 3 {
				arr = arr[len(arr)-3:]
			}
			file = fmt.Sprintf(" %s:%d", strings.Join(arr, "/"), frame.Line)

			return
		},
	})

	return &logrusLogger{
		logrus: lg,
	}
}

func (l *logrusLogger) GetGormLogger() gormlogger.Interface {
	return nil
}

func (l *logrusLogger) GetWhatsappLogger() watsapplog.Logger {
	return nil
}

func (l *logrusLogger) ErrorfNWrap(err error, format string, opts ...any) error {
	l.logrus.WithError(err).Errorf(format, opts...)
	return err
}

func (l *logrusLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	switch level {
	case gormlogger.Silent:
		l.logrus.Level = logrus.FatalLevel
	case gormlogger.Error:
		l.logrus.Level = logrus.ErrorLevel
	case gormlogger.Warn:
		l.logrus.Level = logrus.WarnLevel
	case gormlogger.Info:
		l.logrus.Level = logrus.InfoLevel
	}
	return l
}

func (l *logrusLogger) Info(ctx context.Context, s string, i ...any) {
	l.logrus.WithContext(ctx).Infof(s, i...)
}

func (l *logrusLogger) Warn(ctx context.Context, s string, i ...any) {
	l.logrus.WithContext(ctx).Warnf(s, i...)
}

func (l *logrusLogger) Error(ctx context.Context, s string, i ...any) {
	l.logrus.WithContext(ctx).Errorf(s, i...)
}

func (l *logrusLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.logrus.Level <= logrus.InfoLevel {
		return
	}

	slowThreshold := 200 * time.Millisecond
	IgnoreRecordNotFoundError := false
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.logrus.Level >= logrus.ErrorLevel && (!errors.Is(err, gormlogger.ErrRecordNotFound) || !IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			//l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			l.logrus.WithContext(ctx).Errorf("SQL error: %s", sql)
		} else {
			//l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			l.logrus.WithContext(ctx).Errorf("SQL error: %s, row affected: %d", sql, rows)
		}
	case elapsed > slowThreshold && slowThreshold != 0 && l.logrus.Level >= logrus.WarnLevel:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		if rows == -1 {
			//l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			l.logrus.WithContext(ctx).Warnf("SQL slow %s : %s", slowLog, sql)
		} else {
			//l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			l.logrus.WithContext(ctx).Warnf("SQL slow %s : %s, row affected: %d", slowLog, sql, rows)
		}
	case l.logrus.Level >= logrus.InfoLevel:
		sql, rows := fc()
		if rows == -1 {
			//l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			l.logrus.WithContext(ctx).Tracef("SQL success: %s", sql)
		} else {
			l.logrus.WithContext(ctx).Tracef("SQL success: %s, row affected: %d", sql, rows)
		}
	}
}

func (l *logrusLogger) Sub(module string) watsapplog.Logger {
	l.logrus.WithField("module", module)
	return l
}
