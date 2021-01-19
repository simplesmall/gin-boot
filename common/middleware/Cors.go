package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

/* 文档注解参考
@Summary 接口概要说明
@Description 接口详细描述信息
@Tags 用户信息   //swagger API分类标签, 同一个tag为一组
@accept json  //浏览器可处理数据类型，浏览器默认发 Accept: *
@Produce  json  //设置返回数据的类型和编码
@Param id path int true "ID"    //url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
@Param name query string false "name"
@Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
@Failure 400 {object} Res {"code":200,"data":null,"msg":""}
@Router /test/{id} [get]    //路由信息，一定要写上
*/