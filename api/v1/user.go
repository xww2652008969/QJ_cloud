package v1

import (
	"fmt"
	"gitgub.com/xww2652008969/QJ-cloud/global"
	"gitgub.com/xww2652008969/QJ-cloud/model"
	"gitgub.com/xww2652008969/QJ-cloud/model/Response"
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
	fmt.Println(u.Password)
	if err != nil {
		global.QjLog.Print(err)
		//  传入参数有错
		Response.ResponseRrror(nil, "传参有误", c)
		return
	}
	fmt.Println(len(u.Password))
	if len(u.Name) == 0 || len(u.Password) == 0 || len(u.Email) == 0 { //判断传参是否正确
		Response.ResponseRrror(nil, "传参有误", c)
		return
	}
	user.Name = u.Name
	user.Email = u.Email
	user.Password = u.Password
	err = service.ServiceGroup.Gorm.UserGormapi.UserRegister(user)
	if err != nil {
		//  返回错误
		global.QjLog.Print(err)
		Response.ResponseRrror(nil, err.Error(), c)
		return
	} else {
		//返回正确
		Response.ResponseSuccess(nil, "ok", c)
		return
	}
}
func (userapi *UserApi) UserLogin(c *gin.Context) {
	var u request.REqUserLogin
	var user model.User
	err := c.ShouldBindJSON(&u)
	if err != nil {
		global.QjLog.Print(err)
		Response.ResponseRrror(nil, "传参有误", c)
		return
	}
	if (len(u.Name) == 0 && len(u.Email) == 0) || len(u.Password) == 0 {
		Response.ResponseRrror(nil, "传参有误", c)
		return
	}
	user.Name = u.Name
	user.Email = u.Email
	user.Password = u.Password
	fmt.Println(user)
	err, token := service.ServiceGroup.Gorm.UserGormapi.UserLogin(user)
	if err != nil {
		Response.ResponseRrror(nil, "密码错误", c)
		return
	}
	var resuser Response.ResUserLogin
	resuser.Token = token
	Response.ResponseSuccess(resuser, "ok", c)
}
