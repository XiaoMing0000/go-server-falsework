package controller

import (
	"encoding/json"
	"fmt"
	"main/internal/app"
	"main/internal/exceptions"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Getting(c *app.Context) {
	c.Log().Info("", zap.String("Method", c.Request.Method))

	c.Body = &gin.H{"method": "GET"}
}

func Posting(c *app.Context) {
	c.Log().Info("Posting")
	body, err := c.GetRawData()
	if err != nil {
		fmt.Println("未获取 data")
	}
	var data gin.H
	json.Unmarshal(body, &data)
	if err != nil {
		// 读取失败（客户端断开、body 过大等）
	} else {
		c.Body = &data
	}
}

func Putting(c *app.Context) {
	c.Log().Info("Putting")
	c.Body = &gin.H{"method": "PUT"}
}

func Deleting(c *app.Context) {
	c.Log().Info("Deleting")
	c.Body = &gin.H{"method": "DELETE"}
}

func Patching(c *app.Context) {
	c.Log().Info("Patching")
	c.Body = &gin.H{"method": "PATCH"}
}
func GetError(c *app.Context) {
	c.Log().Info("GetError")
	c.Body = &gin.H{"method": "PUT"}
	c.Error(exceptions.ErrNotFound)
	// log.Fatal("抛出异常")
	panic(fmt.Errorf("doSomething failed:"))
}
