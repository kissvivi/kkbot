package kkbot

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type GameNumberBomb struct {
}


func NewGameNumberBomb() *GameNumberBomb {
	//var winNum,status int
	//rand.Seed(time.Now().Unix())
	//winNum = rand.Intn(100)
	//status = 1
	//fmt.Println(fmt.Sprintf("构建游戏结构体%d,%d",winNum,status))
	_NowMin = 0
	_NowMax = 100
	return &GameNumberBomb{}
}

func (n GameNumberBomb)SetWinNum(){

	rand.Seed(time.Now().Unix())
	_WinNum = rand.Intn(100)
	fmt.Println(fmt.Sprintf("重新随机-----【%d】",_WinNum))
}
var _NowMin,_NowMax,_WinNum,_Status int
func (n GameNumberBomb) Start()  {
	Send(fmt.Sprintf("[数字炸弹-游戏开始]"))
	_Status = 1
}

func (n GameNumberBomb)Gaming(msg FmMessage)  {
	winNumber :=_WinNum
	fmt.Println(fmt.Sprintf("游戏内开始检验正确与否，%s，%d",msg.User.Username,winNumber))
	fmt.Println(fmt.Sprintf("目前的游戏状态%d",_Status))
	if _Status == 1{
		intMsg, err := strconv.Atoi(msg.Message)
		if err != nil {
			fmt.Println(err)
		}else if intMsg< 101 && intMsg>=0{
			fmt.Println(fmt.Sprintf("发的是数字，进入游戏[%s]_[%d]",msg.User.Username,intMsg))
			if winNumber == intMsg  {
				fmt.Println("猜对了")
				n.Win(msg)
			}
			if winNumber > intMsg{
				fmt.Println("猜小了")
				if intMsg > _NowMin{
					_NowMin = intMsg
				}
				Send(fmt.Sprintf("@%s,猜小了哦！答案在[%d]-[%d]之间",msg.User.Username,_NowMin,_NowMax))
			}
			if winNumber < intMsg{
				fmt.Println("猜大了")
				if intMsg < _NowMax{
					_NowMax = intMsg
				}
				Send(fmt.Sprintf("@%s,猜大了哦！答案在[%d]-[%d]之间",msg.User.Username,_NowMin,_NowMax))
			}
		}

	}
}


func (n GameNumberBomb) Over()  {
	Send(fmt.Sprintf("[数字炸弹-游戏结束]"))
	n.SetWinNum()
	_Status = 0
	_NowMax = 100
	_NowMin = 0
}

func (n GameNumberBomb) Win(msg FmMessage)  {
	Send(fmt.Sprintf("@%s,恭喜你猜对了,正确答案是：%d",msg.User.Username,_WinNum))
	n.SetWinNum()
	_NowMax = 100
	_NowMin = 0
}

func (n GameNumberBomb) Next()  {
	//fmt.Println(fmt.Sprintf("游戏继续---%d",n.winNum))
}

