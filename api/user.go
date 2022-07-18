package api

import (
	"Redrock/dao"
	"Redrock/models"
	"Redrock/service"
	"encoding/json"
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

	var user models.Userinfo
	con, err := Upgrade.Upgrade(c.Writer, c.Request, nil)
	_, msg, err := con.ReadMessage()
	err = json.Unmarshal(msg, &user)
	if err != nil {
		return
	}
	m, err := dao.Register(user)
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
	var user models.Userinfo
	con, _ := Upgrade.Upgrade(c.Writer, c.Request, nil)
	_, msg, _ := con.ReadMessage()
	user.Uid = time.Now().Nanosecond() % 100
	json.Unmarshal(msg, &user)
	_, err := service.Login(user)
	var re models.Response
	var m []byte
	if err != nil {
		re.State = 0
		re.Msg = err.Error()
		m, _ = json.Marshal(&re)
		con.WriteMessage(websocket.TextMessage, m)
		con.Close()
		return
	} else {
		re.State = 2
		re.Msg = "登录成功"
		m, _ = json.Marshal(&re)
	}
	client := &Client{con: con, info: user}
	err = client.Menu()
	if err != nil {
		return
	}
}
