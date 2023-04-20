package config

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

// changeConfigLock 编辑配置文件锁，避免并发修改配置
var changeConfigLock sync.RWMutex

// Load 初始化读取配置
func Load[C any](path string, c *C) error {
	// 读取配置
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 解析到 struct
	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

// LoadAndWatch 初始化读取配置，并监控配置变化
func LoadAndWatch[C any](path string, c *C) error {
	// 读取配置
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 解析到 struct
	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	// 监控配置文件
	viper.OnConfigChange(func(e fsnotify.Event) {
		changeConfigLock.Lock()
		defer changeConfigLock.Unlock()

		if err := viper.Unmarshal(c); err != nil {
			slog.Error("change config failed", err)
		} else {
			slog.Info("config changed", slog.Any("new_config", *c))
		}
	})
	viper.WatchConfig()

	return nil
}
