package router

import (
	"main/internal/app"
	"main/internal/controller"

	"github.com/gin-gonic/gin"
)

func testRouter(rg *gin.RouterGroup) *gin.RouterGroup {
	rg.GET("", app.Handle(controller.Getting))
	rg.GET("/error", app.Handle(controller.GetError))
	rg.POST("", app.Handle(controller.Posting))
	rg.PUT("", app.Handle(controller.Putting))
	rg.PATCH("", app.Handle(controller.Patching))
	rg.DELETE("", app.Handle(controller.Deleting))
	return rg
}
