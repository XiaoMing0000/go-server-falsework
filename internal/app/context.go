package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Context struct {
	*gin.Context
	logger  *zap.Logger
	traceId string
}

func (c *Context) TraceId() string {
	return c.traceId
}

// 设置当前请求的 trace id
func (c *Context) SetTraceId(traceId string) {
	c.traceId = traceId
}

// 设置上加文日志
func (c *Context) SetLog(logger *zap.Logger) {
	c.logger = logger
}

// 获取上下文日志
func (c *Context) Log() *zap.Logger {
	return c.logger
}

// 自定义 Context 结构
func Handle(handle func(ctx *Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, ok := c.Get("app_context")
		if !ok {
			ctx := &Context{
				Context: c,
				logger:  zap.NewNop(),
			}
			c.Set("app_context", ctx)
			handle(ctx)
		} else {
			ctx := value.(*Context)
			handle(ctx)
		}
	}
}
