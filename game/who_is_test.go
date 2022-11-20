package game

import (
	kkbot "github.kissvivi.kkbot"
	"math/rand"
	"testing"
	"time"
)

func TestWhoIS_InitCards(t *testing.T) {
	rand.Seed(time.Now().Unix())
	who:=NewWhoIS()
	who.InitCards("./whois.xlsx")
	user:=kkbot.FmMessage{}.User
	user.UserID = 1
	user.Username = "1号选手"

	user2:=kkbot.FmMessage{}.User
	user2.UserID = 2
	user2.Username = "2号选手"

	user3:=kkbot.FmMessage{}.User
	user3.UserID = 1
	user3.Username = "3号选手"

	user4:=kkbot.FmMessage{}.User
	user4.UserID = 2
	user4.Username = "4号选手"

	time.Sleep(1*time.Second)
	msg1:=kkbot.FmMessage{User: user,Message: "[加入-谁是卧底]"}
	msg2:=kkbot.FmMessage{User: user2,Message: "[加入-谁是卧底]"}
	msgu3:=kkbot.FmMessage{User: user3,Message: "[加入-谁是卧底]"}
	msgu4:=kkbot.FmMessage{User: user4,Message: "[加入-谁是卧底]"}

	who.Join("谁是卧底",msg1)
	who.Join("谁是卧底",msg2)
	who.Join("谁是卧底",msgu3)
	who.Join("谁是卧底",msgu4)

	msg999:=kkbot.FmMessage{User: user,Message: "[开始游戏-谁是卧底]"}
	who.Start(msg999)

	time.Sleep(5*time.Second)
	msg3:=kkbot.FmMessage{User: user,Message: "1号选手"}
	msg4:=kkbot.FmMessage{User: user2,Message: "2号选手"}
	msg33:=kkbot.FmMessage{User: user3,Message: "1号选手"}
	msg34:=kkbot.FmMessage{User: user4,Message: "1号选手"}
	who.Gaming(msg3)
	who.Gaming(msg4)
	who.Gaming(msg33)
	who.Gaming(msg34)

	time.Sleep(5*time.Second)
	msg13:=kkbot.FmMessage{User: user,Message: "卧底>1号选手"}
	msg14:=kkbot.FmMessage{User: user2,Message: "卧底>2号选手"}
	msg133:=kkbot.FmMessage{User: user3,Message: "卧底>3号选手"}
	msg134:=kkbot.FmMessage{User: user4,Message: "卧底>1号选手"}
	who.Gaming(msg13)
	who.Gaming(msg14)
	who.Gaming(msg133)
	who.Gaming(msg134)


}
