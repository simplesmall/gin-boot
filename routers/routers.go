package routers

import (
	"gin-boot/common/middleware"
	apiv1 "gin-boot/controller/api/v1"
	"github.com/gin-gonic/gin"
	_ "gin-boot/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer() {
	server:=gin.Default()
	// 配置swagger
	server.Use(middleware.Cors())
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 路由分组
	api:=server.Group("api")
	{
		v1:=api.Group("v1")
		{
			v1.GET("/login",apiv1.Login())
			v1.GET("/notfound",apiv1.NotFoundPage())
		}
	}
	server.Run(":8090")
}