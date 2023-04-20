package config

import "golang.org/x/exp/slog"

// Config 通用配置模版
type Config struct {
	Server Server `json:"server" mapstructure:"server"` // 服务端配置
	MySQL  MySQL  `json:"mysql" mapstructure:"mysql"`   // MySQL 数据库配置
	Redis  Redis  `json:"redis" mapstructure:"redis"`   // Redis 数据库配置
	Log    Log    `json:"log" mapstructure:"log"`       // Log 日志配置
}

// Server 常用配置
type Server struct {
	Host          string `json:"host" mapstructure:"host"`                     // host
	Port          string `json:"port" mapstructure:"port"`                     // 端口
	Version       string `json:"version" mapstructure:"version"`               // 服务版本
	BasePath      string `json:"base_path" mapstructure:"base_path"`           // 根路径
	Env           Env    `json:"env" mapstructure:"env"`                       // 所属环境
	EnableSwagger bool   `json:"enable_swagger" mapstructure:"enable_swagger"` // 是否开启 swagger 文档
}

// Env 运行环境
type Env string

const (
	DEV  Env = "dev"  // 开发环境
	Test Env = "test" // 测试环境
	Prod Env = "prod" // 生产环境
)

// MySQL 数据库配置
type MySQL struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
	UserName string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

// Redis 数据库配置
type Redis struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

// Log 日志配置
type Log struct {
	Path       string     `json:"path" mapstructure:"path"`               // 日志存放路径
	Filename   string     `json:"filename" mapstructure:"filename"`       // 日志文件名
	Level      slog.Level `json:"level" mapstructure:"level"`             // 日志级别。-4:debug, 0:info, 4:warn, 8:error
	MaxSize    int        `json:"max_size" mapstructure:"max_size"`       // 单个日志文件最大大小，单位：MB
	MaxBackups int        `json:"max_backups" mapstructure:"max_backups"` // 最多保存多少个日志文件
	MaxAge     int        `json:"max_age" mapstructure:"max_age"`         // 最多保存多长时间的日志文件，单位：天
}
