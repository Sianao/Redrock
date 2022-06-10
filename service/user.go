package service

import (
	"Redrock/dao"
	"Redrock/models"
	"errors"
	"fmt"
	"log"
)

func Register(usr models.Userinfo) error {
	var tmp string
	result := dao.DB.Select("username").Where("username=?", usr.Username).Find(&tmp)
	fmt.Println(tmp)
	if tmp == "" {
		result = dao.DB.Create(&usr)
		log.Println(result.RowsAffected)
		fmt.Println(result.Error)
		return nil
	}
	err := errors.New("该用户名已被注册 换个试试吧")
	return err

}
