package logi

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field zapcore.Field

func (f Field) ToZap() zap.Field {
	return zap.Field(f)
}

func FieldsToZap(fs []Field) []zap.Field {
	ret := make([]zap.Field, 0, len(fs))
	for _, f := range fs {
		ret = append(ret, f.ToZap())
	}
	return ret
}

func String(k, v string) Field {
	return Field(zap.String(k, v))
}

func Any(k string, v any) Field {
	return Field(zap.Any(k, v))
}
