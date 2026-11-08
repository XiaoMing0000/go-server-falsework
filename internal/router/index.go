package router

import (
	"main/internal/app"
	"main/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Router(rg *gin.RouterGroup) {
	apiRg := rg.Group("/api", app.Handle(middleware.Trace()), app.Handle(middleware.Api))
	{
		testRouter(apiRg.Group("/test"))
	}
}
