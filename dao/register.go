package dao

import (
	"Redrock/models"
	"log"
)

func Register(usr models.Userinfo) (string, error) {
	result := DB.Create(&usr)
	if result.Error != nil {
		log.Println(result.Error)
		return "", result.Error
	}
	return "注册成功", nil

}
