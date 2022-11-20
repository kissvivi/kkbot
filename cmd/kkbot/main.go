package main

import (
	"fmt"
	kkbot "github.kissvivi.kkbot"
	"github.kissvivi.kkbot/game"
	"github.kissvivi.kkbot/util"
)

var _nowGame string

func main() {

	////初始化config
	////初始化application
	////启动数据库
	//cfg, err := kkbot.InitConfig()
	//if err != nil {
	//	panic(err)
	//}
	////初始化数据库
	//baseDB := db.NewBaseDB("mysql")
	//baseDB.SetConfig(cfg)
	//baseDB.InitDB()

	fmt.Print("[KiKiBot] 请输入直播间 ID: ")
	var rid int64
	_, err := fmt.Scanln(&rid)
	if err != nil {
		util.Error("直播间 ID 不正确", nil)
		return
	}

	ch := make(chan kkbot.FmMessage)

	go kkbot.Connect(ch, rid)
	go func() {
		//var ga = &kkbot.GameNumberBomb{}
		ga := kkbot.NewGameNumberBomb()
		who := game.NewWhoIS()
		for msgg := range kkbot.MsgCh {

			switch msgg.Message {
			case "[开始游戏-数字炸弹]":
				fmt.Println("开始游戏")
				if msgg.User.UserID == 23696295 {
					ga.Start()
					_nowGame = "[数字炸弹]"
				} else {
					kkbot.Send("非管理员无法开始游戏")
				}
			case "[结束游戏-数字炸弹]":
				fmt.Println("结束游戏")
				if msgg.User.UserID == 23696295 {
					ga.Over()

				} else {
					kkbot.Send("非管理员无法结束游戏")
				}

			case "[开始游戏-谁是卧底]":
				if msgg.User.UserID == 23696295 {
					who.InitCards("./game/whois.xlsx")
					who.Start(msgg)
					_nowGame = "[谁是卧底]"
				} else {
					kkbot.Send("非管理员无法结束游戏")
				}
			case "[加入-谁是卧底]":
				who.Join("谁是卧底", msgg)
			case "[谁是卧底-help]":
				who.Help()
			case "[谁是卧底-playerList]":
				who.PlayerList()
			}
			switch _nowGame {
			case "[谁是卧底]":
				who.Gaming(msgg)
			case "[数字炸弹]":
				ga.Gaming(msgg)
			}

		}

	}()

	//go func() {
	//	//var ga = &kkbot.GameNumberBomb{}
	//	gans := game.NewNumberSeven()
	//
	//
	//	for msgg := range ch {
	//
	//		if msgg.Message == "[开始游戏-数7]" {
	//			fmt.Println("开始游戏")
	//			if msgg.User.UserID == 23696295 {
	//				gans.Start()
	//
	//			} else {
	//				kkbot.Send("非管理员无法开始游戏")
	//			}
	//
	//		}
	//		if msgg.Message == "[结束游戏-数7]" {
	//			fmt.Println("结束游戏")
	//			if msgg.User.UserID == 23696295 {
	//				gans.Over()
	//
	//			} else {
	//				kkbot.Send("非管理员无法结束游戏")
	//			}
	//		}
	//
	//		gans.Gaming(msgg)
	//	}
	//
	//}()

	//for  {
	//	select {
	//	case x, ok := <-ch:
	//		if ok{
	//			switch x.Message {
	//			case constants.CmdGameStart:
	//			case constants.CmdNumberBomb:
	//			case constants.CmdGameStart:
	//			case constants.CmdGameStart:
	//			case constants.CmdGameStart:
	//
	//
	//			}
	//		}
	//	}
	//}

	for msg := range ch {
		kkbot.HandleFmMessage(msg)
	}
}
