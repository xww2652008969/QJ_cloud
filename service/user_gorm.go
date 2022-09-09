package service

import (
	"errors"
	"fmt"
	"gitgub.com/xww2652008969/QJ-cloud/global"
	"gitgub.com/xww2652008969/QJ-cloud/model"
	"github.com/google/uuid"
)

type UserGormService struct {
}

//用户注册
func (usersevices *UserGormService) UserRegister(user model.User) error {
	var u model.User
	if global.QjDb.Where("name=?", user.Name).First(&u).Error == nil {
		return errors.New("用户注册")
	} else {
		user.Uuid = uuid.Must(uuid.NewRandom())                 //生产uuui
		user.Workaddres = "work" + fmt.Sprintf("%v", user.Uuid) //创建工作目录
		err := global.QjDb.Create(&user).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//用户登录
func (usersevices *UserGormService) UserLogin(user model.User) error {
	var u model.User
	err := global.QjDb.Where("name=? or Emali=? and password=?", user.Name, user.Email, u.Password).First(&u).Error //验证密码
	if err != nil {
		return errors.New("密码错误")
	}
	return nil
}

//用户更新信息
func (usersevices *UserGormService) UserUpdate(user model.User, password string) error {
	var u model.User
	r := global.QjDb.Model(&u).Where("name=? or Emali=? and password=?", user.Name, user.Email, user.Password).Update("password", password)
	if r.RowsAffected != 1 {
		global.QjDb.Begin().Rollback() //回滚
		return errors.New("更新出错")
	} else {
		global.QjDb.Begin().Commit() //提交
		return nil
	}
}

//用户登出
func (usersevices *UserGormService) UserExit() {

}
