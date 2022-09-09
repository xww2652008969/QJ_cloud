package v1

import (
	"gitgub.com/xww2652008969/QJ-cloud/global"
	"gitgub.com/xww2652008969/QJ-cloud/model"
	"gitgub.com/xww2652008969/QJ-cloud/model/request"
	"gitgub.com/xww2652008969/QJ-cloud/service"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func (userapi *UserApi) UserRegister(c *gin.Context) {
	var u request.ReqUseRegister
	var user model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		global.QjLog.Print(err)
		//  传入参数有错
	}
	user.Name = u.Name
	user.Email = u.Emali
	user.Password = u.Emali
	err = service.ServiceGroup.Gorm.UserGormapi.UserRegister(user)
	if err != nil {
		//  返回错误
		global.QjLog.Print(err)
	} else {
		//返回正确
	}
}
func (userapi *UserApi) UserLogin(c *gin.Context) {
	var u request.REqUserLogin
	var user model.User
	err := c.ShouldBindJSON(&u)
	user.Name = u.Name
	if err != nil {
		global.QjLog.Print(err)
		//传入参数有误
	}
	err = service.ServiceGroup.Gorm.UserGormapi.UserLogin(user)
}
