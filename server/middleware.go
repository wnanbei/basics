package server

import (
	"io"
	"time"

	"github.com/galaxy-toolkit/server/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

// NewLoggerHandler 创建服务日志中间件
func NewLoggerHandler(conf config.Server, outputWriter io.Writer) fiber.Handler {
	loggerConfig := logger.Config{
		Next:         nil,
		Format:       "[${time}] ${status} - ${latency} ${method} ${path} ${queryParams}\n",
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
