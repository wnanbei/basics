package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/wnanbei/basics/config"
)

// NewLimiterHandler 创建限流器中间件
func NewLimiterHandler(conf config.Server) fiber.Handler {
	limiterConfig := limiter.Config{
		Max:        conf.GlobalLimiterMax,
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
