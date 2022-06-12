package api

import (
	"Redrock/models"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"math"
	"strconv"
	"strings"
	"time"
)

func (c *Client) NMatches() {
	//开新局
	t := time.Now().Unix() % 10000
	fmt.Println(t)
	room := &Room{
		Rid: strconv.Itoa(int(t)),
	}
	room.Client = append(room.Client, c)
	// 将房间加入到全局
	On.Room[strconv.Itoa(int(t))] = room

	msg := []byte("房间号是" + strconv.Itoa(int(t)))
	c.con.WriteMessage(websocket.TextMessage, msg)
	go room.Competition()
}

func (c *Client) Enter() {
	s := "1,输入房间号\n" +
		"2,自由匹配\n" +
		"3,返回上一级"

	c.con.WriteMessage(websocket.TextMessage, []byte(s))
	_, msg, _ := c.con.ReadMessage()
	choose := msg[0]
	switch choose {
	case byte('1'):
		c.EnterRoom()
		break
	case byte('2'):
		break
	case 3:
		break
	default:
	}
}

func (c *Client) EnterRoom() {
	err := c.con.WriteMessage(websocket.TextMessage, []byte("请输入房间号"))
	if err != nil {
		fmt.Println(err)
	}
	t, msg, err := c.con.ReadMessage()
	fmt.Println(t)
	if err != nil {
		c.con.Close()
		delete(On.Client, strconv.Itoa(c.info.Uid))
	}

	//
	if v, ok := On.Room[string(msg)]; ok {
		// 这下才算进入房间
		fmt.Println(ok)
		if len(On.Room[string(msg)].Client) == 1 {
			v.Client = append(v.Client, c)
			// 对手信息相互加入
			v.Client[0].Competition = c
			c.Competition = v.Client[0]
			fmt.Println("算是进入房间了")
		}
	}
}
func (r *Room) Competition() {
	//等到用户都到了
	for len(r.Client) == 1 {
	}
	fmt.Println("用户都到达了")
	//双方用户到达
	//开始对决
	//据说需要进行准备是吧
	r.Ready()

}

func (r *Room) Ready() {
	state := "1,进入准备\n" +
		"2,退出准备"
	r.Client[0].Fal = false
	r.Client[1].Fal = false
	go func() {
		for r.Client[0].Fal != true || r.Client[1].Fal != true {
			r.Client[0].con.WriteMessage(websocket.TextMessage, []byte(state))
			_, msg, _ := r.Client[0].con.ReadMessage()
			if string(msg) == "1" {
				r.Client[0].Fal = true
			} else {
				r.Client[0].Fal = false
			}
		}
	}()
	for r.Client[0].Fal != true || r.Client[1].Fal != true {
		r.Client[1].con.WriteMessage(websocket.TextMessage, []byte(state))
		_, msg, _ := r.Client[1].con.ReadMessage()
		if string(msg) == "1" {
			r.Client[1].Fal = true
		} else {
			r.Client[1].Fal = false
		}
	}
	r.Game()

}

// 写了个将死的情况 但是 一直判别失败

func (r *Room) Game() {
	//初始化返回棋盘 以及初始化用户侧信息
	var s []rune
	m := ChessTable()
	r.Info = m
	s = r.ReturnDate(m)
	table := NewLogic(m)
	//开始对局
	r.Table = table
	// 发送心跳 开始游戏
	r.Client[0].con.WriteMessage(websocket.PingMessage, nil)
	r.Client[1].con.WriteMessage(websocket.PingMessage, nil)

	r.Client[0].con.WriteMessage(websocket.TextMessage, []byte(string(s)))
	r.Client[1].con.WriteMessage(websocket.TextMessage, []byte(string(s)))
	//
	//r.Client[0].Fal = false
	r.Client[0].Room = r
	r.Client[1].Room = r

	for {
		r.Client[0].con.WriteMessage(websocket.TextMessage, []byte(string(s)))
		_, msg, _ := r.Client[0].con.ReadMessage()
		for {
			//判断用户直到输入正确信息
			err := r.Junge(msg, 0)
			if err == nil {
				break
			}
			r.Client[0].con.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			_, msg, _ = r.Client[0].con.ReadMessage()
		}
		//b, _ := WinOr(0, r)
		s = r.ReturnDate(m)
		//if b {
		//	s = []rune("你输了")
		//}
		fmt.Println(s, string(m[0][0]))
		r.Client[1].con.WriteMessage(websocket.TextMessage, []byte(string(s)))
		_, msg, _ = r.Client[1].con.ReadMessage()
		for {
			err := r.Junge(msg, 1)
			if err == nil {
				break
			}
			r.Client[1].con.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			_, msg, _ = r.Client[1].con.ReadMessage()
		}
		//b, _ = WinOr(1, r)

		s = r.ReturnDate(m)
	}
}

func (r *Room) ReturnDate(m [][]rune) []rune {
	var s []rune
	for _, v := range m {
		s = append(s, v...)
	}
	return s
}
func (r *Room) Junge(msg []byte, che int) error {
	var s [4]int
	m := strings.Split(string(msg), " ")
	fmt.Println(len(m), "len")
	for k, v := range m {
		s[k], _ = strconv.Atoi(v)
	}

	_, err := r.IsMy(s, che)
	if err != nil {
		return err
	}
	fmt.Println(s, "int")

	return nil
}
func (r *Room) IsMy(step [4]int, che int) (string, error) {

	// 判断是否动了别人的棋子 以及是否输入的是空格
	if r.Table[step[0]][step[1]].state != che || r.Table[step[0]][step[1]].T == models.Kong {
		if r.Table[step[0]][step[1]].T == models.Kong {
			err := errors.New("该位置没有你的子")
			return "", err
		}
		err := errors.New("不是你的棋子 返回")
		return "", err
	}
	// 判断是否走到自己的兵的地方
	fmt.Println(string(r.Table[step[2]][step[3]].T))
	if r.Table[step[2]][step[3]].state == che && r.Table[step[2]][step[3]].T != models.Kong {
		err := errors.New("干嘛 吃你自己的棋子")
		return "", err
	}
	//判断是否越界
	for _, v := range step {
		if v > 9 || v < 0 {
			err := errors.New("瞎几把输 ")
			return "", err
		}
	}
	//判断走位是否符合逻辑
	switch r.Table[step[0]][step[1]].T {
	case models.Che:
		fmt.Println("Che\n\n")
		err := r.Che(step)
		if err != nil {
			return "", err
		}
		//
		r.Info[step[2]][step[3]] = models.Che
		r.Info[step[0]][step[1]] = models.Kong
		r.Table[step[2]][step[3]].state = che
		r.Table[step[2]][step[3]].T = models.Che
		r.Table[step[0]][step[1]].T = models.Kong
		fmt.Println("hello")
	case models.Ma:
		fmt.Println("Ma\n\n")
		err := r.Ma(step)
		if err != nil {
			return "", err
		}
		r.Info[step[2]][step[3]] = models.Ma
		r.Info[step[0]][step[1]] = models.Kong
		r.Table[step[2]][step[3]].state = che
		r.Table[step[2]][step[3]].T = models.Ma
		r.Table[step[0]][step[1]].T = models.Kong

	case models.Xian:
		fmt.Println("Xian \n\n")
		err := r.Xian(step)
		if err != nil {
			return "", err
		}
		r.Info[step[2]][step[3]] = models.Xian
		r.Info[step[0]][step[1]] = models.Kong
		r.Table[step[2]][step[3]].state = che
		r.Table[step[2]][step[3]].T = models.Xian
		r.Table[step[0]][step[1]].T = models.Kong
	case models.Shi:
		fmt.Println("Shi \n\n")
		err := r.Shi(step)
		if err != nil {
			return "", err
		}
		r.Info[step[2]][step[3]] = models.Shi
		r.Info[step[0]][step[1]] = models.Kong
		r.Table[step[2]][step[3]].state = che
		r.Table[step[2]][step[3]].T = models.Shi
		r.Table[step[0]][step[1]].T = models.Kong
	case models.Jiang:
		fmt.Println("Jiang \n\n")
		err := r.Jiang(step)
		if err != nil {
			return "", err
		}
		r.Info[step[2]][step[3]] = models.Jiang
		r.Info[step[0]][step[1]] = models.Kong
		r.Table[step[2]][step[3]].state = che
		r.Table[step[2]][step[3]].T = models.Jiang
		r.Table[step[0]][step[1]].T = models.Kong
	case models.Pao:
		fmt.Println("Pao \n\n")
		_, err := r.Pao(step)
		if err != nil {
			return "", err
		}

		r.Info[step[2]][step[3]] = models.Pao
		r.Info[step[0]][step[1]] = models.Kong
		r.Table[step[2]][step[3]].state = che
		r.Table[step[2]][step[3]].T = models.Pao
		r.Table[step[0]][step[1]].T = models.Kong
	case models.Bing:
		fmt.Println("Bing \n\n")
		err := r.Bing(step, che)
		if err != nil {
			return "", err
		}
		r.Info[step[2]][step[3]] = models.Bing
		r.Info[step[0]][step[1]] = models.Kong
		r.Table[step[2]][step[3]].state = che
		r.Table[step[2]][step[3]].T = models.Bing
		r.Table[step[0]][step[1]].T = models.Kong
	}
	return "", nil
}
func NewLogic(m [][]rune) (chess [10][10]Chess) {
	for k, v := range m {
		for m, n := range v {
			chess[k][m].T = n
			if k <= 4 {
				chess[k][m].state = 0
			} else {
				chess[k][m].state = 1
			}

		}
	}
	return
}
func ChessTable() [][]rune {

	table := [][]rune{
		[]rune("车马象士将士象马车\n"),
		[]rune("空空空空空空空空空\n"),
		[]rune("空炮空空空空空炮空\n"),
		[]rune("兵空兵空兵空兵空兵\n"),
		[]rune("空空空空空空空空空\n"),
		[]rune("空空空空空空空空空\n"),
		[]rune("兵空兵空兵空兵空兵\n"),
		[]rune("空炮空空空空空炮空\n"),
		[]rune("空空空空空空空空空\n"),
		[]rune("车马象士将士象马车\n"),
	}

	return table
}

//  通过对该步骤判断是否将死
// 自救 拜托 真的写不粗来

func Move(che int, r *Room) bool {
	op := 1 - che
	var mo Room
	var mun [4]int
	mo = *r
	// 对每个元素尝试移动 再判断 当存在不被将死的情况 返回
	for k, v := range mo.Table {
		for i, j := range v {
			if j.state == che {
				mun[0] = k
				mun[1] = i
				switch j.T {
				case models.Che:
					//横移
					mun[2] = k
					mo = *r
					for i := 0; i <= 8; i++ {
						mo = *r
						mun[3] = i
						//尝试移动
						// 返回 err 说明不能移动
						_, err := mo.IsMy(mun, che)
						if err != nil {
							continue
						}
						b, _ := WinOr(op, &mo)
						if b == false {
							return false
						}
					}
					mun[3] = k

					for j := 0; j <= 9; j++ {
						//尝试移动
						mo = *r
						mun[2] = j
						_, err := mo.IsMy(mun, che)
						if err != nil {
							continue
						}
						b, _ := WinOr(op, &mo)
						if b == false {
							return false
						}
					}
				case models.Ma:
					mun[0] = k
					mun[1] = i
					for i := -2; i <= 2; i++ {
						for j := -2; j <= 2; j++ {
							mo = *r
							mun[2] = k + i
							mun[3] = j + i
							_, err := mo.IsMy(mun, che)
							if err != nil {
								continue
							}
							b, _ := WinOr(op, &mo)
							if b == false {
								return false
							}

						}
					}
				case models.Pao:
					mun[2] = k
					for i := 0; i <= 8; i++ {
						mo = *r
						mun[3] = i
						//尝试移动
						// 返回 err 说明不能移动
						_, err := mo.IsMy(mun, che)
						if err != nil {
							continue
						}
						b, _ := WinOr(op, &mo)
						if b == false {
							return false
						}
					}
					mo = *r
					mun[3] = k

					for j := 0; j < 9; j++ {
						//尝试移动
						mun[2] = j
						_, err := mo.IsMy(mun, che)
						if err != nil {
							continue
						}
						b, _ := WinOr(op, &mo)
						if b == false {
							return false
						}
					}
				case models.Shi:
					mun[0] = k
					mun[1] = i
					for l := -1; l <= 1; l++ {
						for j := -1; j <= 1; j++ {
							mo = *r
							if i == 0 || j == 0 {
								continue
							}
							mun[2] = l + k
							mun[3] = j + i
							_, err := mo.IsMy(mun, che)
							if err != nil {
								continue
							}
							b, _ := WinOr(op, &mo)
							if b == false {
								return false
							}
						}

					}
				case models.Jiang:
					mun[0] = k
					mun[1] = i

					for m := -1; m <= 1; m++ {
						for n := -1; n <= 1; n++ {
							if math.Abs(float64(m+n)) == 1 {
								mo = *r
								mun[2] = m
								mun[3] = n
								_, err := mo.IsMy(mun, che)
								if err != nil {
									continue
								}
								b, _ := WinOr(op, &mo)
								if b == false {
									return false
								}
							}

						}

					}
				case models.Bing:
					mun[0] = k
					mun[1] = i

					for m := -1; m <= 1; m++ {
						for n := -1; n <= 1; n++ {
							if math.Abs(float64(m+n)) == 1 {
								mo = *r
								mun[2] = m
								mun[3] = n
								_, err := mo.IsMy(mun, che)
								if err != nil {
									continue
								}
								b, _ := WinOr(op, &mo)
								if b == false {
									return false
								}
							}
						}
					}
				case models.Xian:
					mun[0] = k
					mun[1] = i
					for m := -2; m <= 2; m++ {
						for n := -2; n <= 2; n++ {
							if math.Abs(float64(m))+math.Abs(float64(n)) == 4 {
								mo = *r
								mun[2] = m
								mun[3] = n
								_, err := mo.IsMy(mun, che)
								if err != nil {
									continue
								}
								b, _ := WinOr(op, &mo)
								if b == false {
									return false
								}
							}
						}
					}
				}
			}
		}
	}
	return true
}
func WinOr(che int, r *Room) (bool, string) {
	// 对 输入的信息进行判断
	var mun [4]int
	//我这步到达的地方作为起点

	mun[2] = r.Jia[1-che][0]
	mun[3] = r.Jia[1-che][1]

	//导入对方将的位置看是否可以到达 将的位置
	for k, v := range r.Table {
		for m, n := range v {
			// 判断是否能够直接到达对面的地方 //判断 0  是否赢了  的时候 就是赢了
			if n.state == che && n.T != models.Kong {
				mun[0] = k
				mun[1] = m
				switch n.T {
				case models.Ma:
					err := r.Ma(mun)
					if err == nil {
						//如果能直接到 就对 其另外一个元素进行移动测试
						b := Move(1-che, r)
						if b {
							return true, ""
						}
					}

				case models.Che:
					err := r.Che(mun)
					//ERR=NIL 说明该处 可以到达
					if err == nil {
						b := Move(che, r)
						if b {
							return true, ""
						}
					}
				case models.Bing:
					err := r.Bing(mun, che)
					if err == nil {
						b := Move(che, r)
						if !b {
							return false, ""
						}
					}
				case models.Pao:
					n, err := r.Pao(mun)
					if err != nil && n == 2 {
						b := Move(che, r)
						if !b {
							return false, ""
						}
					}

				}

			}
		}
	}
	return true, ""
}
