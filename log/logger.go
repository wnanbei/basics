package log

import (
	"context"
	"io"

	"github.com/galaxy-toolkit/server/config"
	"golang.org/x/exp/slog"
)

// Logger 日志
var Logger *slog.Logger
var LoggerWriter io.Writer

// InitLogger 初始化日志
func InitLogger(conf config.Log) {
	LoggerWriter = Writer(conf)

	logger, err := New(conf, LoggerWriter)
	if err != nil {
		panic(err)
	}
	Logger = logger
}

// Debug 日志
func Debug(ctx context.Context, msg string, args ...interface{}) {
	args = append(args, "requestid", ctx.Value("requestid"))
	Logger.DebugCtx(ctx, msg, args...)
}

// Info 日志
func Info(ctx context.Context, msg string, args ...interface{}) {
	args = append(args, "requestid", ctx.Value("requestid"))
	Logger.InfoCtx(ctx, msg, args...)
}

// Warn 日志
func Warn(ctx context.Context, msg string, args ...interface{}) {
	args = append(args, "requestid", ctx.Value("requestid"))
	Logger.WarnCtx(ctx, msg, args...)
}

// Error 日志
func Error(ctx context.Context, msg string, args ...interface{}) {
	args = append(args, "requestid", ctx.Value("requestid"))
	Logger.ErrorCtx(ctx, msg, args...)
}
