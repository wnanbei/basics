package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

// changeConfigLock 编辑配置文件锁，避免并发修改配置
var changeConfigLock sync.RWMutex

// Load 初始化读取配置
func Load[C any](path string, c *C) error {
	if path == "" {
		path = defaultConfigPath("config.json")
	}

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
	if path == "" {
		path = defaultConfigPath("config.json")
	}

	// 读取配置
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 解析到 struct
	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	slog.Info("config read completed", slog.Any("config", *c))

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

// defaultConfigPath 获取默认的配置文件路径，为 go.mod 文件所处的位置
func defaultConfigPath(filename string) string {
	if filename == "" {
		filename = "config.json"
	}
	return filepath.Join(goModDir(""), filename)
}

// goModDir 逐层往上寻找 go.mod 文件，并返回他的父级文件夹的地址.
func goModDir(currentDir string) string {
	if currentDir == "" {
		currentDir, _ = os.Getwd()
	}

	if currentDir == string(os.PathSeparator) {
		return ""
	}

	f := filepath.Join(currentDir, "go.mod")
	_, err := os.Stat(f)
	if err == nil {
		// 找到了.
		return filepath.Dir(f)
	}

	if os.IsNotExist(err) {
		return goModDir(filepath.Dir(currentDir))
	}
	// 权限错误 或者没有找到 返回空
	return ""
}
