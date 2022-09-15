package router

import (
	"gitgub.com/xww2652008969/QJ-cloud/api"
	"gitgub.com/xww2652008969/QJ-cloud/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (receiver *UserRouter) InitUserRouter(group *gin.RouterGroup) {
	api := api.ApigroupApp.Userapi
	user := group.Group("/user")
	{
		user.POST("register", api.UserRegister)
		user.POST("login", api.UserLogin)
	}
	user.Use(middleware.JwtAuth())
	{
		user.POST("update", api.UserUpdate)
	}
}
