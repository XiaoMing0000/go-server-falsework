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
	c.Log().Info(fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL))
	c.Next()

	// 检测异常
	if len(c.Errors) == 0 {
		c.Log().Debug(fmt.Sprintf("Benchmark took %d nanoseconds", time.Since(start)))
	} else {
		err := c.Errors.Last().Err
		var appErr *exceptions.AppError
		if errors.As(err, &appErr) {
			c.Log().Warn(appErr.Message, zap.Uint32("code", appErr.Code), zap.Int("status", appErr.Status))
			c.JSON(appErr.Status, gin.H{
				"success": false,
				"code":    appErr.Code,
				"message": appErr.Message,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"code":    500,
				"message": "an unexpected error occurred",
			})
		}
	}

	c.Log().Info(fmt.Sprintf("Benchmark took %d nanoseconds", time.Since(start)))
}
