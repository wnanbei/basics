package basicCtx

import "context"

type CtxKeyType string

const (
	MetaKey CtxKeyType = "basics-meta"
)

// WithMeta 设置元数据信息
func WithMeta(ctx context.Context, meta *Meta) context.Context {
	return context.WithValue(ctx, MetaKey, meta)
}

// GetMeta 获取元数据信息
func GetMeta(ctx context.Context) *Meta {
	meta, ok := ctx.Value(MetaKey).(*Meta)
	if !ok {
		return &Meta{}
	}
	return meta
}

// WithTraceID 设置 trace id
func WithTraceID(ctx context.Context, traceID string) context.Context {
	meta := GetMeta(ctx)
	meta.WithTraceID(traceID)
	return WithMeta(ctx, meta)
}

// WithUserID 设置 user id
func WithUserID(ctx context.Context, userID int64) context.Context {
	meta := GetMeta(ctx)
	meta.WithUserID(userID)
	return WithMeta(ctx, meta)
}
