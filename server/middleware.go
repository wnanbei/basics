package server

import (
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/redis/v2"

	"github.com/wnanbei/basics/config"
	"github.com/wnanbei/basics/constant"
	basicCtx "github.com/wnanbei/basics/context"
	"github.com/wnanbei/basics/log"
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

// Logger variables
const (
	TagTime              = "time"
	TagPid               = "pid"
	TagReferer           = "referer"
	TagProtocol          = "protocol"
	TagIP                = "ip"
	TagIPs               = "ips"
	TagHost              = "host"
	TagMethod            = "method"
	TagPath              = "path"
	TagURL               = "url"
	TagUA                = "ua"
	TagLatency           = "latency"
	TagStatus            = "status"
	TagQueryStringParams = "query"
	TagReqBody           = "req_body"
	TagReqHeaders        = "req_headers"
	TagResBody           = "res_body"
	TagError             = "error"
	TagTraceID           = "trace_id"
)

// NewSlogMiddleware creates a new slog middleware handler
func NewSlogMiddleware() fiber.Handler {
	// Set PID once
	pid := strconv.Itoa(os.Getpid())

	// Set variables
	var (
		once       sync.Once
		errHandler fiber.ErrorHandler
	)

	// instead of analyzing the template inside(handler) each time, this is done once before
	// and we create several slices of the same length with the functions to be executed and fixed parts.
	// templateChain, logFunChain, err := buildLogFuncChain(&cfg, createTagMap(&cfg))
	// if err != nil {
	// 	panic(err)
	// }

	// Return new handler
	return func(c *fiber.Ctx) error {
		// Set error handler once
		once.Do(func() {
			// override error handler
			errHandler = c.App().ErrorHandler
		})

		// Set latency start time
		startTime := time.Now()

		// Handle request, store err for logging
		chainErr := c.Next()

		// Manually call error handler
		var errStr string
		if chainErr != nil {
			errStr = chainErr.Error()
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError) //nolint:errcheck // TODO: Explain why we ignore the error here
			}
		}

		slog.LogAttrs(
			c.UserContext(), slog.LevelInfo, "request info",
			slog.String(TagPid, pid),
			slog.String(TagReferer, c.Get("Referer")),
			slog.String(TagProtocol, c.Protocol()),
			slog.String(TagIP, c.IP()),
			slog.String(TagIPs, c.Get(fiber.HeaderXForwardedFor)),
			slog.String(TagHost, c.Hostname()),
			slog.String(TagMethod, c.Method()),
			slog.String(TagPath, c.Path()),
			slog.String(TagURL, c.OriginalURL()),
			slog.String(TagUA, c.Get(fiber.HeaderUserAgent)),
			slog.Int64(TagLatency, time.Since(startTime).Microseconds()),
			slog.Int(TagStatus, c.Response().StatusCode()),
			slog.String(TagReqBody, string(c.Body())),
			slog.String(TagError, errStr),
			slog.String(TagResBody, string(c.Response().Body())),
			log.GroupAttrs(TagReqHeaders, log.HeaderToAttrs(&c.Request().Header)...),
			log.GroupAttrs(TagQueryStringParams, log.ArgsToAttrs(c.Request().URI().QueryArgs())...),
		)

		return nil
	}
}

const (
	// HeaderTraceID is the header key for trace id
	HeaderTraceID = "X-Request-ID"
)

// NewTraceMiddleware creates a new trace middleware handler
func NewTraceMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get id from request, else we generate one
		traceID := c.Get(HeaderTraceID)
		if traceID == "" {
			traceID = utils.UUID()
		}

		// Set new id to response header
		c.Set(HeaderTraceID, traceID)

		// Add the request ID to locals
		ctx := basicCtx.WithTraceID(c.Context(), traceID)
		c.SetUserContext(ctx)

		return c.Next()
	}
}
