package log

import (
	"log/slog"

	"github.com/wnanbei/basics/config"
)

// InitLogger 初始化日志
func Init(logconf config.Log) error {
	logger, err := New(logconf)
	if err != nil {
		return err
	}

	slog.SetDefault(logger)
	return nil
}

// New 创建日志实例
func New(logConf config.Log) (*slog.Logger, error) {
	writer, err := NewWrite(logConf)
	if err != nil {
		return nil, err
	}

	replace := func(groups []string, a slog.Attr) slog.Attr {
		return a
	}

	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       logConf.Level,
		ReplaceAttr: replace,
	}
	handler := NewMetaHandler(slog.NewJSONHandler(writer, &opts))

	return slog.New(handler), nil
}
