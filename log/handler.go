package log

import (
	"context"
	"log/slog"

	basicCtx "github.com/wnanbei/basics/context"
)

const (
	TraceIDKey = "trace_id"
	UserIDKey  = "user_id"
)

type MetaHandler struct {
	handler slog.Handler
}

func NewMetaHandler(handler slog.Handler) *MetaHandler {
	return &MetaHandler{handler: handler}
}

// Enabled reports whether the handler handles records at the given level.
func (h *MetaHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

// WithAttrs returns a new JSONHandler whose attributes consists
// of h's attributes followed by attrs.
func (h *MetaHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &MetaHandler{handler: h.handler.WithAttrs(attrs)}
}

// WithGroup returns a new JSONHandler whose group consists
// of h's group followed by group.
func (h *MetaHandler) WithGroup(group string) slog.Handler {
	return &MetaHandler{handler: h.handler.WithGroup(group)}
}

// Handle formats its argument Record as a JSON object on a single line.
func (h *MetaHandler) Handle(ctx context.Context, rec slog.Record) error {
	meta := basicCtx.GetMeta(ctx)
	rec.AddAttrs(slog.String(TraceIDKey, meta.TraceID()), slog.Int64(UserIDKey, meta.UserID()))

	return h.handler.Handle(ctx, rec)
}
