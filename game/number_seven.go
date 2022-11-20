package game

import (
	"fmt"
	kkbot "github.kissvivi.kkbot"
	"github.kissvivi.kkbot/util"
	"strconv"
	"strings"
)
var (
	_IsWinSeven,_StatusSeven bool
	_WinUserMsg              kkbot.FmMessage
	_StartNum                int
)
type NumberSeven struct {
	IsWin bool
	Status     bool
	WinUserMsg kkbot.FmMessage
	StartNum   int
}

func NewNumberSeven() *NumberSeven {
	//_StartNum = 1
	//_StatusSeven = false
	//_StartNum = 1
	reset()
	return &NumberSeven{StartNum:1}
}

func (n NumberSeven) Start() {
	kkbot.Send(fmt.Sprintf("[数7-游戏开始]"))
	n.Status = true
	_StatusSeven = true
}

func (n NumberSeven) Gaming(msg kkbot.FmMessage) {

	util.Info("[数7-游戏内部开始]")
	var (
		winFlagInt int
		winFlagString string
	)

	winFlagInt, err := strconv.Atoi(msg.Message)
	if err != nil {
		fmt.Println(err)
	}
	if _StatusSeven{

		if winFlagInt != _StartNum+1{
			kkbot.Send(fmt.Sprintf("@%s[输入不规范数字,接受惩罚吧！！]",msg.User.Username))
			n.Win(msg)
		}else if winFlagInt % 7 == 0 || strings.Contains(winFlagString, "7"){
			n.IsWin = true
			_IsWinSeven = true
			n.WinUserMsg = msg
			_WinUserMsg = msg
			n.Win(msg)
		}else if winFlagString == "pass"{
			n.Next()
		}
		_StartNum++
	}

}

func (n NumberSeven) Over() {
	kkbot.Send(fmt.Sprintf("[数7-游戏结束]"))
	n.Status = false
	_StatusSeven = false
	reset()
}

func (n NumberSeven) Win(msg kkbot.FmMessage) {
	fmt.Println("win")
	kkbot.Send(fmt.Sprintf("@%s[恭喜你,输入了7相关数字[%s],接受惩罚吧！！]",msg.User.Username,msg.Message))
	reset()
	//panic("implement me")
}

func (n NumberSeven) Next() {
	fmt.Println("继续游戏，下一位")
	kkbot.Send(fmt.Sprintf("[拍桌子-继续游戏]"))
}

func reset()  {
	//_StatusSeven = false
	_IsWinSeven = false
	_StartNum = 1
}

//var _ GameI = (*NumberSeven)(nil)
