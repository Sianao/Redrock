package service

import (
	"Redrock/dao"
	"Redrock/models"
	"errors"
)

func Register(usr models.Userinfo) (string, error) {
	result := dao.DB.Create(&usr)
	if result.Error != nil {
		err := errors.New("注册失败 换个用户名试试吧")
		return "", err
	}
	return "注册成功", nil

}

func Login(userinfo models.Userinfo) (models.Userinfo, error) {
	var pwd models.Userinfo
	result := dao.DB.Model(models.Userinfo{}).Where("username=?", userinfo.Username).Scan(&pwd)
	if result.RowsAffected == 0 {
		err := errors.New("用户信息不存在")
		return models.Userinfo{}, err
	}
	if pwd.Password == userinfo.Password {
		pwd.Password = ""
		return pwd, nil
	} else {
		err := errors.New("密码错误")
		return pwd, err
	}
}
