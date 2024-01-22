package redis

import (
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/wnanbei/basics/config"
)

// New 根据配置生成 Postgres 数据库实例
func New(conf config.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + strconv.Itoa(conf.Port),
		Password: conf.Password,
		DB:       conf.Database,
	})
}
