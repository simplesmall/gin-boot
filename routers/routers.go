package routers

import (
	"gin-boot/common/middleware"
	"gin-boot/common/middleware/jwt"
	"gin-boot/config"
	apiv1 "gin-boot/controller/api/v1"
	"gin-boot/model/student"
	"github.com/gin-gonic/gin"
	_ "gin-boot/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer() {
	server:=gin.Default()
	// 初始化数据库连接
	config.InitConnect()

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
		v2:=api.Group("v2")
		{
			v2.GET("/login",jwt.JWTAuthMiddleware(),jwt.JWTHandler)
			v2.POST("/auth",jwt.AuthHandler)
		}
		v3:=api.Group("v3")
		{
			v3.POST("/student",student.CreateStudent())
			v3.GET("/student/:ID",student.GetStudent())
			v3.GET("/class/:ID",student.GetClass())
		}
	}
	server.Run(":8090")
}