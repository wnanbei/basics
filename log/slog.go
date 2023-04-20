package log

import (
	"os"
	"path/filepath"

	"github.com/galaxy-toolkit/server/config"
	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// New 创建日志实例
func New(log config.Log) (*slog.Logger, error) {
	if err := os.MkdirAll(log.Path, 0777); err != nil {
		return nil, err
	}

	writer := &lumberjack.Logger{
		Filename:   filepath.Join(log.Path, log.Filename),
		MaxSize:    log.MaxSize,
		MaxBackups: log.MaxBackups,
		MaxAge:     log.MaxAge,
	}

	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       log.Level,
		ReplaceAttr: nil,
	}

	return slog.New(opts.NewJSONHandler(writer)), nil
}
