package config

import "golang.org/x/exp/slog"

// Config 通用配置模版
type Config struct {
	Server   Server   `json:"server" mapstructure:"server"`     // 服务端配置
	Log      Log      `json:"log" mapstructure:"log"`           // Log 日志配置
	Database Database `json:"database" mapstructure:"database"` // 数据库配置
}

// Server 服务配置
type Server struct {
	Host          string `json:"host" mapstructure:"host"`                     // host
	Port          string `json:"port" mapstructure:"port"`                     // 端口
	Version       string `json:"version" mapstructure:"version"`               // 服务版本
	BasePath      string `json:"base_path" mapstructure:"base_path"`           // 根路径
	Env           Env    `json:"env" mapstructure:"env"`                       // 所属环境
	EnableSwagger bool   `json:"enable_swagger" mapstructure:"enable_swagger"` // 是否开启 swagger 文档
	Title         string `json:"title" mapstructure:"title"`                   // 服务标题名称，用于 swagger
	Monitor       bool   `json:"monitor" mapstructure:"monitor"`               // 是否开启可视化监控
}

// Env 运行环境
type Env string

const (
	DEV  Env = "dev"  // 开发环境
	Test Env = "test" // 测试环境
	Prod Env = "prod" // 生产环境
)

// Log 日志配置
type Log struct {
	Path       string     `json:"path" mapstructure:"path"`               // 日志存放路径
	Filename   string     `json:"filename" mapstructure:"filename"`       // 日志文件名
	Level      slog.Level `json:"level" mapstructure:"level"`             // 日志级别。-4:debug, 0:info, 4:warn, 8:error
	MaxSize    int        `json:"max_size" mapstructure:"max_size"`       // 单个日志文件最大大小，单位：MB
	MaxBackups int        `json:"max_backups" mapstructure:"max_backups"` // 最多保存多少个日志文件
	MaxAge     int        `json:"max_age" mapstructure:"max_age"`         // 最多保存多长时间的日志文件，单位：天
}

// Database 数据库通用配置模版
type Database struct {
	MySQL    MySQL    `json:"mysql" mapstructure:"mysql"`       // MySQL 数据库配置
	Postgres Postgres `json:"postgres" mapstructure:"postgres"` // Postgres 数据库配置
	Redis    Redis    `json:"redis" mapstructure:"redis"`       // Redis 数据库配置
}

// MySQL 数据库配置
type MySQL struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
	UserName string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

// Postgres 数据库配置
type Postgres struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
	UserName string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

// Redis 数据库配置
type Redis struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Password string `json:"password" mapstructure:"password"`
	Database int    `json:"database" mapstructure:"database"`
}

func defaultConfig() *Config {
	return &Config{
		Server: Server{
			Host:          "127.0.0.1",
			Port:          "9999",
			Version:       "v1.0.0",
			BasePath:      "/",
			Env:           DEV,
			EnableSwagger: true,
			Title:         "server",
			Monitor:       true,
		},
		Log: Log{
			Path:       "logs",
			Filename:   "server.log",
			Level:      slog.LevelDebug,
			MaxSize:    128,
			MaxBackups: 10,
			MaxAge:     30,
		},
		Database: Database{
			MySQL: MySQL{
				Host:     "127.0.0.1",
				Port:     "3306",
				UserName: "root",
				Password: "root",
				Database: "server",
			},
			Postgres: Postgres{
				Host:     "127.0.0.1",
				Port:     "5432",
				UserName: "root",
				Password: "root",
				Database: "server",
			},
			Redis: Redis{
				Host:     "127.0.0.1",
				Port:     6379,
				Password: "",
				Database: 0,
			},
		},
	}
}
