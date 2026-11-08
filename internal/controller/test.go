package controller

import (
	"encoding/json"
	"fmt"
	"main/internal/app"
	"main/internal/exceptions"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Getting(c *app.Context) {
	c.Log().Info("", zap.String("Method", c.Request.Method))

	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func Posting(c *app.Context) {
	c.Log().Info("Posting")
	// name := c.PostForm("name")
	// age := c.PostForm("age")
	age := c.PostFormMap("age")
	body, err := c.GetRawData()
	if err != nil {
		fmt.Println("未获取 data")
	}
	var data interface{}
	json.Unmarshal(body, &data)
	fmt.Println("data: ", data)
	fmt.Println("age: ", age)
	fmt.Println("age: ", reflect.TypeOf(age))
	// c.JSON(http.StatusOK, gin.H{"method": "POST"})
	c.Data(200, "application/json", body)
}

func Putting(c *app.Context) {
	c.Log().Info("Putting")
	c.JSON(http.StatusOK, gin.H{"method": "PUT"})
}

func Deleting(c *app.Context) {
	c.Log().Info("Deleting")
	c.JSON(http.StatusOK, gin.H{"method": "DELETE", "test": "6"})
}

func Patching(c *app.Context) {
	c.Log().Info("Patching")
	c.JSON(http.StatusOK, gin.H{"method": "PATCH"})
}
func GetError(c *app.Context) {
	c.Log().Info("GetError")
	c.Error(exceptions.ErrNotFound)

}
