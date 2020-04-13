package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go-mall/lib/config/env"
)

var logger *zap.SugaredLogger

// Debug debug log
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf info format log
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Info info log
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof info format log
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warn warn log
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf warn format log
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error error log
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf error format log
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Panic panic log
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf panic format log
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Fatal log message, and call os.Exist.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf log format message, and call os.Exist.
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Build  build the logger.
func Build() {
	options := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}
	ads := make([]zapcore.Core, len(adapters))
	for i, adapter := range adapters {
		ads[i] = adapter()
	}
	core := zapcore.NewTee(ads...)
	logger = zap.New(core, options...).Sugar()
}

// DefaultEncoderConfig default profile for zapcore.EncoderConfig.
var DefaultEncoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
	EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
}

// DefaultLevelEnablerFunc implement zapcore.LevelEnabler.
//  the minimal level for development is Debug, and others is Info.
func DefaultLevelEnablerFunc() zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		if env.IsDevelopment() {
			return level >= zapcore.DebugLevel
		}
		return level >= zapcore.InfoLevel
	}
}

// Instance is a function create a new logger adapter.
type Instance func() zapcore.Core

var adapters []Instance

// Register registers a logger adapter.
func Register(adapter ...Instance) {
	if adapter == nil {
		panic("logger: Register adapter is nil")
	}

	adapters = append(adapters, adapter...)
}
