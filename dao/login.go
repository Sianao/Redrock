package dao

import (
	"Redrock/models"
	"errors"
)

func Login(userinfo models.Userinfo) (string, error) {
	var tmpu models.Userinfo
	result := DB.Model(models.Userinfo{}).Where("username=?", userinfo.Username).Scan(&tmpu)
	if result.RowsAffected == 0 {
		err := errors.New("用户信息不存在")
		return "", err
	}
	if tmpu.Password == userinfo.Password {

		return "", nil
	} else {
		err := errors.New("密码错误")
		return "", err
	}
}
