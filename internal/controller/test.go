package controller

import (
	"encoding/json"
	"fmt"
	"main/internal/app"
	"main/internal/exceptions"
	"main/internal/types"

	"go.uber.org/zap"
)

func Getting(c *app.Context) {
	c.Logger.Info("", zap.String("Method", c.Request.Method))

	*c.Body = map[string]any{"method": "GET"}
}

func Posting(c *app.Context) {
	var data types.TestPostReqData
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		c.Logger.Warn(fmt.Sprintf("Test Posting decode data error, %+v\n", err))
		c.Error(exceptions.HttpError(400))
		return
	}
	c.Logger.Info("Posting")

	// 赋值 body
	*c.Body = &data
}

func Putting(c *app.Context) {
	c.Logger.Info("Putting")
}

func Deleting(c *app.Context) {
	c.Logger.Info("Deleting")

	*c.Body = map[string]any{"method": "DELETE"}
}

func Patching(c *app.Context) {
	c.Logger.Info("Patching")
	*c.Body = map[string]any{"method": "PATCH"}
}
func GetError(c *app.Context) {
	c.Logger.Info("GetError")
	// c.Body = &gin.H{"method": "PUT"}
	// c.Error(exceptions.ErrNotFound)
	// // log.Fatal("抛出异常")
	// panic(fmt.Errorf("doSomething failed:"))
}
