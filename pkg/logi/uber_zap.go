package logi

import (
	"context"
	"fmt"
	"time"

	watsapplog "go.mau.fi/whatsmeow/util/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

type zapLogger struct {
	*zap.Logger
	level       zapcore.Level
	enumVerbose EnumVerbose
}

func SetZap(level EnumVerbose) {
	current = newDevZap(level, zap.AddCallerSkip(2))
}

func newDevZap(level EnumVerbose, opts ...zap.Option) *zapLogger {
	callerEncoder := zapcore.ShortCallerEncoder
	//timeEncoder := zapcore.ISO8601TimeEncoder
	//callerEncoder = zapcore.FullCallerEncoder
	timeEncoder := customTimeEncoder
	if level == DebugLevel || level == TraceLevel {
		callerEncoder = zapcore.FullCallerEncoder
	}

	defaultLogLevel := level.ToZapLevel()
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(defaultLogLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:    "time",
			LevelKey:   "level",
			NameKey:    "lg",
			CallerKey:  "caller",
			MessageKey: "msg",
			//StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     timeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   callerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	lg, err := cfg.Build(opts...)
	if err != nil {
		panic(err)
	}

	if lg == nil {
		panic("logger nil")
	}
	defer func() { _ = lg.Sync() }()

	return &zapLogger{
		Logger:      lg,
		level:       defaultLogLevel,
		enumVerbose: level,
	}
}

func (l *zapLogger) Debug(msg string, fields ...Field) {
	l.Logger.Debug(msg, FieldsToZap(fields)...)
}

func (l *zapLogger) Warnf(msg string, args ...any) {
	l.Logger.Warn(fmt.Sprintf(msg, args...))
}

func (l *zapLogger) Errorf(msg string, args ...any) {
	l.Logger.Error(fmt.Sprintf(msg, args...))
}

func (l *zapLogger) Infof(msg string, args ...any) {
	l.Logger.Info(fmt.Sprintf(msg, args...))
}

func (l *zapLogger) Debugf(msg string, args ...any) {
	l.Logger.Debug(fmt.Sprintf(msg, args...))
}

func (l *zapLogger) Fatalf(msg string, args ...any) {
	l.Logger.Fatal(fmt.Sprintf(msg, args...))
}

func (l *zapLogger) Panicf(msg string, args ...any) {
	l.Logger.Panic(fmt.Sprintf(msg, args...))
}

func (l *zapLogger) Info(ctx context.Context, msg string, args ...any) {
	l.Logger.Info(fmt.Sprintf(msg, args...))
}

func (l *zapLogger) Warn(ctx context.Context, msg string, args ...any) {
	l.Logger.Warn(fmt.Sprintf(msg, args...))
}

func (l *zapLogger) Error(ctx context.Context, msg string, args ...any) {
	l.Logger.Error(fmt.Sprintf(msg, args...))
}

func (l *zapLogger) ErrorfNWrap(err error, format string, opts ...any) error {
	l.Logger.Error(fmt.Sprintf(format, opts...), zap.Error(err))
	return err
}

func (l *zapLogger) GetGormLogger() gormlogger.Interface {
	return &zapGorm{
		Logger:      l.Logger,
		level:       l.level,
		enumVerbose: l.enumVerbose,
	}
}

func (l *zapLogger) GetWhatsappLogger() watsapplog.Logger {
	return &zapWhatsapp{
		Logger:      l.Logger,
		level:       l.level,
		enumVerbose: l.enumVerbose,
	}
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("15:04:05.99"))
}

// WhatsappLogger

type zapWhatsapp struct {
	*zap.Logger
	level       zapcore.Level
	enumVerbose EnumVerbose
}

func (z *zapWhatsapp) Warnf(msg string, args ...interface{}) {
	z.Warn(fmt.Sprintf(msg, args...))
}

func (z *zapWhatsapp) Errorf(msg string, args ...interface{}) {
	z.Error(fmt.Sprintf(msg, args...))
}

func (z *zapWhatsapp) Infof(msg string, args ...interface{}) {
	z.Info(fmt.Sprintf(msg, args...))
}

func (z *zapWhatsapp) Debugf(msg string, args ...interface{}) {
	z.Debug(fmt.Sprintf(msg, args...))
}

func (z *zapWhatsapp) Sub(module string) watsapplog.Logger {
	z.Logger = z.With(zap.String("module", module))
	return z
}
