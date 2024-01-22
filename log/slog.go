package log

import (
	"io"
	"os"
	"path/filepath"

	"github.com/wnanbei/basics/config"
	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Basic *Logger  // Basic 基础日志
var Server *Logger // Server 服务日志

const (
	ServerGroupName = "server" // server 组名
)

// InitLogger 初始化日志
func InitLogger(conf config.Log) {
	LoggerWriter = Writer(conf)

	logger, err := New(conf, LoggerWriter)
	if err != nil {
		panic(err)
	}

	Basic = logger
	Server = Basic.WithGroup(ServerGroupName)
}

// New 创建日志实例
func New(log config.Log, writer io.Writer) (*Logger, error) {
	if err := os.MkdirAll(log.Path, 0777); err != nil {
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
		Level:       log.Level,
		ReplaceAttr: replace,
	}

	return &Logger{slog.New(slog.NewJSONHandler(writer, &opts))}, nil
}

// LoggerWriter 日志写入
var LoggerWriter io.Writer

// Writer 获取滚动写入日志 writer
func Writer(log config.Log) io.Writer {
	return &lumberjack.Logger{
		Filename:   filepath.Join(log.Path, log.Filename),
		MaxSize:    log.MaxSize,
		MaxBackups: log.MaxBackups,
		MaxAge:     log.MaxAge,
	}
}
