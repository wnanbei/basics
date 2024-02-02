package middlewares

import (
	"log/slog"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/wnanbei/basics/log"
)

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
