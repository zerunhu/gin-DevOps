package service

import (
	"errors"
	"gin-DevOps/config"
	"gin-DevOps/model"
	"gin-DevOps/utils"
	"reflect"
)

func Register(u model.User) (err error, userInter model.User) {
	var user model.User
	config.GdoDb.Where("username = ?", u.Username).First(&user)
	if !reflect.DeepEqual(user, model.User{}){
		return errors.New("用户名已注册"), userInter
	}
	u.Password = utils.MD5V([]byte(u.Password))
	err = config.GdoDb.Create(&u).Error
	return err, u
}
