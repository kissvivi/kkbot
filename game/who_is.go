package game

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	kkbot "github.kissvivi.kkbot"
	"github.kissvivi.kkbot/util"
	"math/rand"
	"strings"
)

var queue = make([]Player,0)
var cardMap = make([]cards,0)
var _WHO_STATUS bool
var _W_PLAYER Player

var _sumL,_nowL,_NOW_LUN int

type WhoIS struct {
}

func NewWhoIS() *WhoIS {
	_WHO_STATUS = true

	return &WhoIS{}
}

func (w WhoIS) Gaming(msg kkbot.FmMessage) {
	//util.Info("[游戏进行中--------]")
	if isUserIn(msg) {
		if queue[_nowL].ID == msg.User.UserID{
			w.Next()
		}
		util.Info(fmt.Sprintf("第[%d]轮进行中",_NOW_LUN))
		if _sumL<=2{
			util.Info("[人数小于3，游戏结束--------]")
			w.Win(msg)
		}
		if strings.Contains(msg.Message, "卧底>"){
			str := strings.Split(msg.Message, ">")
			fmt.Println(fmt.Sprintf("[猜卧底---@%s 猜卧底是 -%s]",msg.User.Username,str[1]))
			wduser:=str[1]
			// 去除空格
			wduser = strings.Replace(wduser, " ", "", -1)
			// 去除换行符
			wduser = strings.Replace(wduser, "\n", "", -1)

			setCount(wduser)
			if wduser == _W_PLAYER.Name {
				util.Info(fmt.Sprintf("[猜卧底--- 恭喜@%s,猜对卧底了,卧底是->@%s]",msg.User.Username,_W_PLAYER.Name))
				w.Win(msg)
			}else {
				w.Next()
			}
		}

	}
}

func (w WhoIS) Over() {
	reset()
	queue = nil
	cardMap = nil
}

func (w WhoIS) Win(msg kkbot.FmMessage) {
	reset()
	w.Next()
}

type cards struct {
	winCard string
	loseCard string
}


// Next 发卡
func (w WhoIS) Next() {

	if _NOW_LUN ==0{
		util.Info("[初始化]-[发卡中]----------")
		winNum1 := rand.Intn(len(cardMap))

		for i, c := range cardMap {
			if i == winNum1 {
				util.Info(fmt.Sprintf("[发卡中]-[胜者卡：%s，败者卡：%s]]",c.winCard,c.loseCard))
			}
		}
		if len(queue)>0{
			winNum2 := rand.Intn(len(queue))
			util.Info(fmt.Sprintf("[发卡中]-[卧底：%v]------",queue[winNum2]))
			_W_PLAYER = queue[winNum2]

			for _, player := range queue {

				if player.ID == _W_PLAYER.ID{
					kkbot.SendMessage(int(player.ID),cardMap[0].loseCard)
				}
			}
		}
	}else{
		if  _nowL < len(queue)-1 {
			util.Info(fmt.Sprintf("@%s,开始猜卧底",queue[_nowL].Name))
			_nowL++
		} else if len(queue)-1 == _nowL{
			ss:=shuxString()
			fmt.Println(fmt.Sprintf("阐述顺序为：%s",ss))
			fmt.Println("第一轮投票结束，开始第二轮阐述")
			_NOW_LUN+=1
			_nowL=0

			//t人
		}else{
			_nowL=0
		}
	}


}
var _shux string

func shuxString() string {
	_shux = ""
	var shux string
	for i, player := range queue {

		shux = fmt.Sprintf("%d-%s ",i+1,player.Name)
		_shux+=shux
	}

	return _shux
}
func (w WhoIS) Start(msg kkbot.FmMessage) {
	if isUserIn(msg) {
		if msg.Message == "[开始游戏-谁是卧底]"{
			ss:=shuxString()
			fmt.Println(fmt.Sprintf("阐述顺序为：%s",ss))
			w.Next()
			_NOW_LUN+=1
			//w.Gaming(msg)
		}

	}
}

func (w WhoIS) AddPlayer(player Player) bool {
	queue = append(queue, player)
	_sumL=len(queue)-1
	return true
}

func (w WhoIS) Join(game string, msg kkbot.FmMessage) {
	player:=Player{msg.User.UserID,msg.User.Username,0}
	if msg.Message == "[加入-谁是卧底]"{
		w.AddPlayer(player)
	}
}

func (w WhoIS) InitCards(path string)  {
	// 首先读excel
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get all the rows in the Sheet2.
	rows, _ := xlsx.GetRows("whois")
	for i, row := range rows {
		// 去掉第一行，第一行是表头
		if i == 0 {
			continue
		}
		var data cards
		for j, colCell := range row {
			// 排除第一列为Null
			if j == 0 && colCell == "Null" {
				continue
			}
			// 第一列即是一级
			if j == 0 && colCell != "Null" {
				data.winCard = colCell
			}
			// 第二列即是二级
			if j == 1 {
				data.loseCard = colCell
			}
		}
		cardMap = append(cardMap, data)
	}
}

func getUserNameByID(uid int64) (res bool){
    res =false
	for _, player := range queue {

		if player.ID == uid{
			res= true
			return
		}else{
			res= false
		}
	}
	return
}

func isUserIn(msg kkbot.FmMessage) (res bool) {
	res = false
	if _WHO_STATUS&&getUserNameByID(msg.User.UserID){
		return true
	}
	return
}

func setCount(name string)  {
	for _, player := range queue {
		if player.Name == name {
			player.Count++
		}
	}
}

func reSet()  {
	_W_PLAYER = Player{}
	_NOW_LUN = 0
	_sumL = 0
	_nowL = 0
}