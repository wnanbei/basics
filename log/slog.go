package log

import (
	"log/slog"
	"path/filepath"

	"github.com/wnanbei/basics/config"
)

// InitLogger 初始化日志
func Init(logconf config.Log) {
	logger, err := New(logconf)
	if err != nil {
		panic(err)
	}

	slog.SetDefault(logger)
}

// New 创建日志实例
func New(logConf config.Log) (*slog.Logger, error) {
	writer, err := NewWrite(logConf)
	if err != nil {
		return nil, err
	}

	replace := func(groups []string, a slog.Attr) slog.Attr {
		// Remove the directory from the source's filename.
		if a.Key == slog.SourceKey {
			a.Value = slog.StringValue(filepath.Base(a.Value.String()))
		}
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
