package routers

import (
	apiv1 "gin-boot/controller/api/v1"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	server:=gin.Default()
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