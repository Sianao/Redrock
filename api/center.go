package api

import (
	"Redrock/models"
	"github.com/gorilla/websocket"
)

var On Online

type Step struct {
	Location [2]int `json:"location"`
	To       [2]int `json:"to"`
}
type Chess struct {
	T     rune
	state int
}
type Client struct {
	Room *Room
	// 自己的连接
	// 直接互发 在中间进行判断
	con  *websocket.Conn
	Fal  bool
	info models.Userinfo
	//对手的连接
	Competition *Client
}

type Online struct {
	// 保存所有在线的客户端
	Client map[string]*Client
	//所有在线的房间
	Room map[string]*Room
}
type Room struct {
	Msg chan string
	//一个房间 的信息
	Table  [10][10]Chess
	Info   [][]rune
	Step   [2][2]int
	Jia    [2][2]int
	Client []*Client
	//room id
	Rid    string
	Signal int
}
