package v1

import (
	"gin-boot/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func NotFoundPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"data":response.ResponseStructure.NotFound(response.ResponseStructure{}),
		})
	}
}