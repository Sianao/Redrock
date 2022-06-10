package controller

import (
	"Redrock/api"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	r := gin.Default()
	r.POST("/register", api.Register)
	r.Run()
}
