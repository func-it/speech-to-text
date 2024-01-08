package logi

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
)

var MapLevel = map[EnumVerbose]LevelValue{
	PanicLevel: {logrus.PanicLevel, zapcore.PanicLevel},
	FatalLevel: {logrus.FatalLevel, zapcore.FatalLevel},
	ErrorLevel: {logrus.ErrorLevel, zapcore.ErrorLevel},
	WarnLevel:  {logrus.WarnLevel, zapcore.WarnLevel},
	InfoLevel:  {logrus.InfoLevel, zapcore.InfoLevel},
	DebugLevel: {logrus.DebugLevel, zapcore.DebugLevel},
	TraceLevel: {logrus.TraceLevel, zapcore.DebugLevel},
}

type LevelValue struct {
	Logrus logrus.Level
	Zap    zapcore.Level
}

type EnumVerbose string

const (
	PanicLevel EnumVerbose = "panic"
	FatalLevel EnumVerbose = "fatal"
	ErrorLevel EnumVerbose = "error"
	WarnLevel  EnumVerbose = "warn"
	InfoLevel  EnumVerbose = "info"
	DebugLevel EnumVerbose = "debug"
	TraceLevel EnumVerbose = "trace"
)

func (e *EnumVerbose) ToLogrusLevel() logrus.Level {
	if e == nil {
		return logrus.WarnLevel
	}

	return MapLevel[*e].Logrus
}

func (e *EnumVerbose) ToZapLevel() zapcore.Level {
	if e == nil {
		return zapcore.WarnLevel
	}

	return MapLevel[*e].Zap
}

func ParseLevel(lvl string) (EnumVerbose, error) {
	switch strings.ToLower(lvl) {
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	case "trace":
		return TraceLevel, nil
	}

	return "", fmt.Errorf("not a valid verbose/log Level: %q", lvl)
}

func (e *EnumVerbose) UnmarshalJSON(b []byte) error {
	v, err := ParseLevel(string(b))
	if err != nil {
		return err
	}

	*e = v
	return nil
}

func (e *EnumVerbose) String() string {
	if e == nil || *e == "" {
		return "warn"
	}

	return string(*e)
}

func (e *EnumVerbose) Type() string {
	return "(panic|fatal|error|warn|info|debug|trace)"
}

func (e *EnumVerbose) Set(s string) error {
	if e == nil {
		return fmt.Errorf("unexpected nil value for EnumVerbose")
	}

	if s == "" {
		*e = WarnLevel
		return nil
	}

	v, err := ParseLevel(string(s))
	if err != nil {
		return err
	}

	*e = v

	return nil
}

func EnumVerboseDecodeHook() mapstructure.DecodeHookFuncType {
	// Wrapped in a function call to add optional input parameters (eg. separator)
	return func(f, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(EnumVerbose("")) {
			return data, nil
		}

		s, ok := data.(string)
		if !ok {
			return nil, fmt.Errorf("verbose: unexpected type")
		}

		_, err := ParseLevel(s)
		if err != nil {
			return nil, err
		}

		return data, nil
	}
}
