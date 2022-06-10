package api

import (
	"Redrock/models"
	"Redrock/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(context *gin.Context) {
	var r models.Userinfo
	err := context.ShouldBind(&r)
	fmt.Println(r)
	if err != nil {
		log.Println(err)
		return
	}
	err = service.Register(r)
	if err != nil {
		context.JSON(200, err.Error())
		return
	}
	context.JSON(200, nil)
	return

}
