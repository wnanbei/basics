package log

import (
	"context"

	"golang.org/x/exp/slog"
)

// Logger 日志
type Logger struct {
	*slog.Logger
}

func (l *Logger) Debug(ctx context.Context, msg string, args ...slog.Attr) {
	l.Logger.LogAttrs(ctx, slog.LevelDebug, msg, AppendAttrs(ctx, args)...)
}

func (l *Logger) Info(ctx context.Context, msg string, args ...slog.Attr) {
	l.Logger.LogAttrs(ctx, slog.LevelInfo, msg, AppendAttrs(ctx, args)...)
}

func (l *Logger) Warn(ctx context.Context, msg string, args ...slog.Attr) {
	l.Logger.LogAttrs(ctx, slog.LevelWarn, msg, AppendAttrs(ctx, args)...)
}

func (l *Logger) Error(ctx context.Context, msg string, args ...slog.Attr) {
	l.Logger.LogAttrs(ctx, slog.LevelError, msg, AppendAttrs(ctx, args)...)
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		Logger: l.Logger.With(args...),
	}
}

func (l *Logger) WithGroup(group string) *Logger {
	return &Logger{
		Logger: l.Logger.WithGroup(group),
	}
}

func AppendAttrs(ctx context.Context, attrs []slog.Attr) []slog.Attr {
	return append(attrs,
		slog.String("requestid", ctx.Value("requestid").(string)),
	)
}
