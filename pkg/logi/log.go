package logi

import (
	watsapplog "go.mau.fi/whatsmeow/util/log"
	gormlogger "gorm.io/gorm/logger"
)

type Log interface {
	GetGormLogger() gormlogger.Interface
	GetWhatsappLogger() watsapplog.Logger

	Fatalf(msg string, args ...any)
	Panicf(msg string, args ...any)
	Warnf(msg string, args ...any)
	Errorf(msg string, args ...any)
	Infof(msg string, args ...any)
	Debugf(msg string, args ...any)

	Debug(msg string, fields ...Field)

	ErrorfNWrap(err error, format string, opts ...any) error
}

type LogType int

const (
	TypeUberZap LogType = iota
	TypeLogrus
)
