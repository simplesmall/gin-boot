package main

import (
	"gin-boot/config"
	"gin-boot/routers"
)

// @title gin-boot系统接口
// @version 1.0
// @description gin-boot系统接口监控服务后端API接口文档

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8090
// @BasePath
func main() {
	defer config.CloseDB()
	routers.InitServer()
}
