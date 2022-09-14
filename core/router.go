package core

import (
	"gitgub.com/xww2652008969/QJ-cloud/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	SystemRouter := router.Routerapp
	Groups := r.Group("")
	{
		SystemRouter.User.InitUserRouter(Groups)

	}
	return r
}
