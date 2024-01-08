package logi

import watsapplog "go.mau.fi/whatsmeow/util/log"

var (
	current Log
)

func Debug(format string, fields ...Field) {
	current.Debug(format, fields...)
}

func Debugf(format string, opts ...any) {
	current.Debugf(format, opts...)
}

func Warnf(format string, opts ...any) {
	current.Warnf(format, opts...)
}

func Errorf(format string, opts ...any) {
	current.Errorf(format, opts...)
}

func Fatalf(format string, opts ...any) {
	current.Fatalf(format, opts...)
}

func Infof(format string, opts ...any) {
	current.Infof(format, opts...)
}

func ErrorfNWrapNReturn(err error, format string, opts ...any) error {
	return current.ErrorfNWrap(err, format, opts...)
}

func ErrorNReturn(err error) error {
	current.Errorf(err.Error())
	return err
}

func GetWhatsappLogger() watsapplog.Logger {
	return current.GetWhatsappLogger()
}
