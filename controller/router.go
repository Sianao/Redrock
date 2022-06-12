package controller

import (
	"Redrock/api"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

func Entrance() {
	// 入口 两个
	r := gin.Default()
	//两个事件
	r.GET("/register", api.Register)
	r.GET("/login", api.Login)
	r.Run(":9000")

}
