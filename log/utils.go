package log

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// GroupAttrs 将多个 Attr 组合成一个 Attr
func GroupAttrs(key string, attrs ...slog.Attr) slog.Attr {
	return slog.Attr{Key: key, Value: slog.GroupValue(attrs...)}
}

// ArgsToAttrs 将 fiber url 的参数转换为 slog.Attr
func ArgsToAttrs(args *fiber.Args) []slog.Attr {
	var attrs []slog.Attr
	args.VisitAll(func(key, value []byte) {
		attrs = append(attrs, slog.String(string(key), string(value)))
	})

	return attrs
}

// HeaderToAttrs 将 fiber header 的参数转换为 slog.Attr
func HeaderToAttrs(header *fasthttp.RequestHeader) []slog.Attr {
	var attrs []slog.Attr
	header.VisitAll(func(key, value []byte) {
		attrs = append(attrs, slog.String(string(key), string(value)))
	})

	return attrs
}
