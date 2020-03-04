package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 日志记录器
var Logger *zap.SugaredLogger

// InitLogger 初始化日志记录器
func InitLogger() error {
	logf, err := os.OpenFile(Config.Foxpot.LogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	cfg := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	var lvl zapcore.Level
	switch Config.Foxpot.LogLevel {
	case "DEBUG":
		lvl = zapcore.DebugLevel
	case "INFO":
		lvl = zapcore.InfoLevel
	case "WARN":
		lvl = zapcore.WarnLevel
	case "ERROR":
		lvl = zapcore.ErrorLevel
	case "PANIC":
		lvl = zapcore.PanicLevel
	case "FATAL":
		lvl = zapcore.FatalLevel
	default:
		lvl = zapcore.InfoLevel
	}
	Logger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(cfg), zapcore.AddSync(logf), lvl), zap.AddCaller()).Sugar()
	return nil
}
