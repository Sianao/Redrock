package service

import (
	"Redrock/dao"
	"Redrock/models"
)

func Login(userinfo models.Userinfo) (string, error) {
	var pwd models.Userinfo
	login, err := dao.Login(pwd)
	if err != nil {
		return "", err
	}
	return login, err
}
