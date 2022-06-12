package hander

import (
	"Redrock/models"
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"
)

func Readmsg(con *websocket.Conn) {
	for {
		t, message, err := con.ReadMessage()
		if err != nil {
			con.Close()
			return
		}
		if t == websocket.PingMessage {
			con.WriteControl(websocket.PongMessage, message, time.Now().Add(time.Minute))
		}
		var m models.Message
		json.Unmarshal(message, &m)
		if m.MessageType == "private" {

		}

	}

}
