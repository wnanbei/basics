package log

import (
	"io"
	"os"
	"path/filepath"

	"github.com/wnanbei/basics/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Writer 日志写入位置
func NewWrite(logconf config.Log) (io.Writer, error) {
	switch logconf.Type {
	case config.Stdout:
		return os.Stdout, nil
	case config.File:
		return NewLumberjackWriter(logconf)
	default:
		return os.Stdout, nil
	}
}

// NewLumberjackWriter 获取滚动写入日志 writer
func NewLumberjackWriter(logconf config.Log) (io.Writer, error) {
	if err := os.MkdirAll(logconf.Path, 0777); err != nil {
		return nil, err
	}

	return &lumberjack.Logger{
		Filename:   filepath.Join(logconf.Path, logconf.Filename),
		MaxSize:    logconf.MaxSize,
		MaxBackups: logconf.MaxBackups,
		MaxAge:     logconf.MaxAge,
	}, nil
}
