package redis

import (
	"github.com/galaxy-toolkit/server/config"
	"github.com/redis/go-redis/v9"
)

// New 根据配置生成 Postgres 数据库实例
func New(conf config.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password,
		DB:       conf.Database,
	})
}
