package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	basicCtx "github.com/wnanbei/basics/context"
)

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
