package middleware

import (
	"errors"
	"fmt"
	"main/internal/app"
	"main/internal/exceptions"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Api(c *app.Context) {
	start := time.Now()
	c.Logger.Info(fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL))
	c.Next()

	// 检测异常
	if len(c.Errors) == 0 {
		c.Logger.Debug(fmt.Sprintf("Benchmark took %d nanoseconds", time.Since(start)))
		if !c.Writer.Written() {
			// 自定义相应体结构
			c.JSON(http.StatusOK, gin.H{
				"code":    c.Code,
				"success": c.Success,
				"data":    c.Body,
			})
		}
	} else {
		err := c.Errors.Last().Err
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			if appErr.Status >= 500 {
				c.Logger.Error(appErr.Message, zap.Uint32("code", appErr.Code), zap.Int("status", appErr.Status))
			} else {
				c.Logger.Warn(appErr.Message, zap.Uint32("code", appErr.Code), zap.Int("status", appErr.Status))
			}

			if !c.Writer.Written() {
				c.JSON(appErr.Status, gin.H{
					"code":    appErr.Code,
					"success": appErr.Code != 0,
					"message": appErr.Message,
				})
			}
		} else {
			if !c.Writer.Written() {
				c.JSON(exceptions.ErrInternalServer.Status, gin.H{
					"code":    exceptions.ErrInternalServer.Code,
					"success": false,
					"message": exceptions.ErrInternalServer.Message,
				})
			}
		}
	}

	c.Logger.Info(fmt.Sprintf("Benchmark took %d nanoseconds", time.Since(start)))
}
