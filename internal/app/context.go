package app

import (
	"fmt"
	"log"
	"main/internal/exceptions"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 自定义 Context 解构
// 注意：增加的字段需要在中间件初始化
type Context struct {
	*gin.Context
	logger  *zap.Logger
	traceId string
	Code    int
	Success bool
	Body    *gin.H
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
		var ctx *Context
		// stack := debug.Stack()
		defer func() {
			if r := recover(); r != nil {
				if ctx.logger != nil {
					ctx.Log().Error(fmt.Sprintf("%+v\n", r))
					// 打印调用栈信息
					// ctx.Log().Error(fmt.Sprintf("[PANIC] %v\n%s", r, stack))
				} else {
					log.Printf("%+v\n", r)
					// 打印调用栈信息
					// log.Printf("[PANIC] %v\n%s", r, stack)
				}
				if !c.Writer.Written() {
					c.Error(exceptions.ErrInternalServer)
				} else {
					c.Abort()
				}
			}
		}()

		if !ok {
			ctx = &Context{
				Context: c,
				logger:  zap.NewNop(),
				Code:    0,
				Success: true,
			}
			c.Set("app_context", ctx)
		} else {
			ctx = value.(*Context)
		}
		handle(ctx)
	}
}
