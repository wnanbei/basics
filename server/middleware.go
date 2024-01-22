package server

import (
	"io"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/redis/v2"
	"github.com/spf13/viper"
	"github.com/wnanbei/basics/config"
	"github.com/wnanbei/basics/constant"
)

// NewLoggerHandler 创建服务日志中间件
func NewLoggerHandler(conf config.Server, outputWriter io.Writer) fiber.Handler {
	loggerConfig := logger.Config{
		Next:         nil,
		Format:       "[${time}] ${method} ${path} ${status} ${latency} - ${locals:requestid} - query[${queryParams}] body[${body}]\n",
		TimeFormat:   "2006-01-02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       outputWriter,
	}
	return logger.New(loggerConfig)
}

// NewLimiterHandler 创建限流器中间件
func NewLimiterHandler(conf config.Server) fiber.Handler {
	limiterConfig := limiter.Config{
		Max:        viper.GetInt("server.globalLimiterMax"),
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.SlidingWindow{},
	}
	return limiter.New(limiterConfig)
}

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
