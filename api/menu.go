package api

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
)

const menu = "1,新建房间\n" +
	"2,加入房间\n" +
	"3,观战\n" +
	"4,退出"

var (
	ChooseNewMatch   = byte('1')
	ChooseEnterMatch = byte('2')
	ChooseWatchMatch = byte('3')
	CHooseQuit       = byte('4')
)

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
	case ChooseNewMatch:
		c.NMatches()
		break
	case ChooseEnterMatch:
		c.Enter()
		break
	case ChooseWatchMatch:
		c.Watch()
		break
	case CHooseQuit:
		c.con.Close()
		err := errors.New("退出")
		return err
	default:

	}
	return nil
}
