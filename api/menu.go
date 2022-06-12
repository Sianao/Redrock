package api

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
)

var menu = "1,新建房间\n" +
	"2,加入房间\n" +
	"3,查看信息\n" +
	"4,退出"

func (c *Client) Menu() error {

	err := c.con.WriteMessage(websocket.TextMessage, []byte(menu))
	if err != nil {
		c.con.Close()
		return err
	}
	// 客户端加入
	On.Client[strconv.Itoa(c.info.Uid)] = c
	_, msg, err := c.con.ReadMessage()
	fmt.Println(string(msg))
	if err != nil {
		c.con.Close()
		return err
	}
	choose := msg[0]
	switch choose {
	case byte('1'):
		c.NMatches()
		break
	case byte('2'):
		c.Enter()
		break
	case byte('3'):
		break
	case byte('4'):
		c.con.Close()
		err := errors.New("退出")
		return err
	default:

	}
	return nil
}
