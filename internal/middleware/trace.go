package middleware

import (
	"main/internal/app"
	"main/internal/utils"

	"go.uber.org/zap"
)

const (
	// 客户端可通过这些请求头透传 trace id，优先使用 X-Trace-ID。
	traceIDHeader   = "X-Trace-ID"
	requestIDHeader = "X-Request-ID"

	traceIDField = "trace_id"
)

// Trace 为每个请求建立独立的 trace id 和 logger：
//   - 优先复用客户端传入的 X-Trace-ID / X-Request-ID，非法或缺失时生成新的
//   - trace id 写入 app.Context、响应头
//   - logger 通过 With 派生，只属于当前请求，不同请求之间互不影响
func Trace() func(c *app.Context) {
	return func(c *app.Context) {
		traceID := c.GetHeader(traceIDHeader)
		if !validTraceID(traceID) {
			traceID = c.GetHeader(requestIDHeader)

			if !validTraceID(traceID) {
				traceID = utils.GetUUID()
			}
		}

		c.SetTraceId(traceID)
		c.SetLog(app.Logger().With(zap.String(traceIDField, traceID)))
		c.Header(traceIDHeader, traceID)

		c.Next()
	}
}

// validTraceID 校验外部传入的 trace id，避免超长或非法字符污染日志。
func validTraceID(id string) bool {
	if len(id) == 0 || len(id) > 128 {
		return false
	}
	for _, r := range id {
		switch {
		case r >= 'a' && r <= 'z',
			r >= 'A' && r <= 'Z',
			r >= '0' && r <= '9',
			r == '-', r == '_', r == '.', r == ':':
		default:
			return false
		}
	}
	return true
}
