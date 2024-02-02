package middlewares

import (
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/redis/v2"

	"github.com/wnanbei/basics/config"
	"github.com/wnanbei/basics/constant"
)

// NewSessionStore 创建 session 存储
func NewSessionStore(conf config.Redis) *session.Store {
	store := redis.New(redis.Config{
		Host:      conf.Host,
		Port:      conf.Port,
		Username:  "",
		Password:  conf.Password,
		Database:  conf.Database,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})

	return session.New(session.Config{
		Expiration:        24 * time.Hour,
		Storage:           store,
		KeyLookup:         constant.RedisKeyUserSession,
		CookieDomain:      "",
		CookiePath:        "",
		CookieSecure:      false,
		CookieHTTPOnly:    false,
		CookieSameSite:    "Lax",
		CookieSessionOnly: false,
		KeyGenerator: func() string {
			return "session:" + utils.UUIDv4()
		},
	})
}
