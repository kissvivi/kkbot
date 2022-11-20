package game

import (
	kkbot "github.kissvivi.kkbot"
)

type Player struct {
	ID   int64
	Name string
	Count int
}
type GameI interface {
	Gaming(msg kkbot.FmMessage)
	Over()
	Win(msg kkbot.FmMessage)
	Next()
	Start(msg kkbot.FmMessage)
	AddPlayer(player Player) bool
	Join(game string, msg kkbot.FmMessage)
}

//
//import (
//	"fmt"
//	"kkbot"
//)
//
//var ga = &kkbot.NumberBomb{}
//
//func GameRun() {
//	ga = kkbot.NewNumberBomb()
//	for msgg := range kkbot.MsgCh {
//		if msgg.Message == "-------开始游戏-数字炸弹--------" {
//			fmt.Println("开始游戏")
//			if msgg.User.UserID == 23696295 {
//				kkbot.Start()
//
//			} else {
//				kkbot.Send("非管理员无法开始游戏")
//			}
//
//		}
//		if msgg.Message == "-------结束游戏-数字炸弹--------" {
//			fmt.Println("结束游戏")
//			if msgg.User.UserID == 23696295 {
//				ga.Over()
//
//			} else {
//				kkbot.Send("非管理员无法结束游戏")
//			}
//		}
//		ga.Gaming(msgg)
//	}
//
//}
