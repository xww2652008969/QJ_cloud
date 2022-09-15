package middleware

import (
	"gitgub.com/xww2652008969/QJ-cloud/model/Response"
	"gitgub.com/xww2652008969/QJ-cloud/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.GetHeader("token")
		uuid, err := utils.Parsetoken(t)
		if err != nil {
			c.Abort()
			Response.ResponseRrror(nil, "验证失败", c)
		}
		c.Set("uuid", uuid) //将uuid储存该请求中
		c.Next()
	}
}
