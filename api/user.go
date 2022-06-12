package api

import (
	"Redrock/models"
	"Redrock/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var Upgrade = websocket.Upgrader{
	// 跨域 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Register(c *gin.Context) {
	var u models.Userinfo
	con, _ := Upgrade.Upgrade(c.Writer, c.Request, nil)
	_, msg, _ := con.ReadMessage()
	json.Unmarshal(msg, &u)
	fmt.Println(u)
	m, err := service.Register(u)
	var r models.Response
	if err != nil {
		r.State = 0
		r.Msg = err.Error()
		ms, _ := json.Marshal(&r)
		con.WriteMessage(websocket.TextMessage, ms)
	}
	r.State = 2
	r.Msg = m
	ms, _ := json.Marshal(&r)
	con.WriteMessage(websocket.TextMessage, ms)
	return
}

// 登录成功 建立连接

func Login(c *gin.Context) {
	var usr models.Userinfo
	con, _ := Upgrade.Upgrade(c.Writer, c.Request, nil)
	//_, msg, _ := con.ReadMessage()
	usr.Uid = time.Now().Nanosecond() % 100
	//json.Unmarshal(msg, &usr)
	//login, err := service.Login(usr)
	//var re models.Response
	//var m []byte
	//if err != nil {
	//	re.State = 0
	//	re.Msg = err.Error()
	//	m, _ = json.Marshal(&re)
	//	con.WriteMessage(websocket.TextMessage, m)
	//	con.Close()
	//	return
	//} else {
	//	re.State = 2
	//	re.Msg = "登录成功"
	//	m, _ = json.Marshal(&re)
	//}
	//con.WriteMessage(websocket.TextMessage, [])
	vcons := &Client{con: con, info: usr}
	err := vcons.Menu()
	if err != nil {
		con.Close()
		return
	}
}
