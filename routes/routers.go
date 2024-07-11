package routes

import (
	controllers "gin-air-template/controller"
	"gin-air-template/middleware"
	"github.com/gin-gonic/gin"
)

// InitRoutes 初始化路由
func InitRoutes(router *gin.Engine) {
	// without jwt auth
	router.GET("/ping", controllers.Ping)

	// with jwt auth
	vGroup := router.Group("/api")
	vGroup.Use(middleware.JWTMiddleware())
	vGroup.GET("/profile", func(context *gin.Context) {
		//placeholder
	})
}
