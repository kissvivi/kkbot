package game

import (
	"fmt"
	kkbot "github.kissvivi.kkbot"
	"math/rand"
	"strconv"
	"time"
)

const (
	WIN_MIN = 0
	WIN_MAX = 100
	STATUS_GAMEING = 1
	STATUS_OVER = 2
)

type NumberBombI interface {
	Gaming(msg kkbot.FmMessage, win int)
	Over()
	Win(msg kkbot.FmMessage)
	Next()
}

type NumberBomb struct {
	winNum int
	status int
	nowMin int
	nowMax int
}

func (n NumberBomb) GetWinNum() int {
	return n.winNum
}

func NewNumberBomb() *NumberBomb {
	var winNum, status int
	rand.Seed(time.Now().Unix())
	winNum = rand.Intn(100)
	fmt.Println(fmt.Sprintf("构建游戏结构体%d,%d", winNum, status))
	return &NumberBomb{status: 1, winNum: winNum, nowMin: 0, nowMax: 100}
}

func (n NumberBomb) NewWinNum() {
	rand.Seed(time.Now().Unix())
	n.winNum = rand.Intn(100)
	fmt.Println(fmt.Sprintf("重新随机-----【%d】", n.winNum))
}
func (n NumberBomb) Start() {
	//kkbot.Send(fmt.Sprintf("-----游戏开始------！"))
	fmt.Println(fmt.Sprintf("-----游戏开始------！"))
	n.status = STATUS_GAMEING
}

func (n NumberBomb) Gaming(msg kkbot.FmMessage) {
	var (
		winNumber int
		status    int
		nowMin int
		nowMax int
	)
	winNumber = n.winNum
	fmt.Println(fmt.Sprintf("Gaming------------，%s，%d", msg.User.Username, winNumber))
	fmt.Println(fmt.Sprintf("游戏内开始检验正确与否，%s，%d", msg.User.Username, winNumber))
	fmt.Println(fmt.Sprintf("目前的游戏状态%d", status))
	if status == STATUS_GAMEING {
		intMsg, err := strconv.Atoi(msg.Message)
		if err != nil {
			fmt.Println(err)
		} else if intMsg < 100 && intMsg > 0 {
			fmt.Println(fmt.Sprintf("发的是数字，进入游戏[%s]_[%d]", msg.User.Username, intMsg))
			if winNumber == intMsg {
				fmt.Println("猜对了")
				n.Win(msg)
			}
			if winNumber > intMsg {
				fmt.Println(fmt.Sprintf("@%s,猜小了哦！答案在[%d]-[%d]之间", msg.User.Username, nowMin, nowMax))
				n.nowMin = intMsg
				//kkbot.Send(fmt.Sprintf("@%s,猜小了哦！答案在[%d]-[%d]之间", msg.User.Username, nowMin, nowMax))
			}
			if winNumber < intMsg {
				fmt.Println(fmt.Sprintf("@%s,猜大了哦！答案在[%d]-[%d]之间", msg.User.Username, nowMin, nowMax))
				n.nowMax = intMsg
				//kkbot.Send(fmt.Sprintf("@%s,猜大了哦！答案在[%d]-[%d]之间", msg.User.Username, nowMin, nowMax))
			}
		}

	}
}

func (n NumberBomb) Over() {
	//kkbot.Send(fmt.Sprintf("------游戏结束------！"))
	fmt.Println(fmt.Sprintf("------游戏结束------！"))
	n.status = STATUS_OVER
	n.reset()
}

func (n NumberBomb) Win(msg kkbot.FmMessage) {
	//kkbot.Send(fmt.Sprintf("@%s,恭喜你猜对了,正确答案是：%d", msg.User.Username, n.winNum))
	fmt.Println(fmt.Sprintf("@%s,恭喜你猜对了,正确答案是：%d", msg.User.Username, n.winNum))
	n.reset()
}

func (n NumberBomb) Next() {
	fmt.Println(fmt.Sprintf("游戏继续---%d", n.winNum))
}

func (n NumberBomb)reset()  {
	n.NewWinNum()
	n.nowMin = WIN_MIN
	n.nowMax = WIN_MAX
}
