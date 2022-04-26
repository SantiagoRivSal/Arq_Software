package router

import (
	"router/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/controller", controller.Control)
	r.POST("/login", controller.Login)
	return r
}
