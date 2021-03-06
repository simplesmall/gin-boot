package v1

import (
	"gin-boot/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 登录测试接口
// @Description  登录测试接口描述信息
// @Success 200 {string} string    "ok"
// @Router /api/v1/login [get]
// @Tags 登录
// BaseURL /
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		res:=response.ResponseStructure{
			Code: 0,
			Msg:  "Hello login!!",
			Data: nil,
		}
		c.JSON(http.StatusOK,gin.H{
			"data":res,
		})
	}
}

// @Summary notfound测试接口
// @Description  notfound测试接口描述信息
// @Success 200 {string} string    "ok"
// @Router /api/v1/notfound [get]
// @Tags NotFound
// BaseURL /
func NotFoundPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"data":response.ResponseStructure.NotFound(response.ResponseStructure{}),
		})
	}
}
