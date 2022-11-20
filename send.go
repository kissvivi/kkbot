package kkbot

import (
	"encoding/json"
	"fmt"
	"github.kissvivi.kkbot/util"
	"net/http"
)

//type message struct {
//	RoomID    int    `json:"room_id"`
//	Message   string `json:"message"`
//	MessageID string `json:"msg_id"`
//}

//// Send keep taking messages from the channel output,
//// send messages to the live room according to roomID on MissEvan.
//func Send(ctx context.Context, output <-chan string, room *models.Room) {
//	for {
//		select {
//		case <-ctx.Done():
//			return
//		case msg := <-output:
//			if msg != "" {
//				send(msg, room)
//			}
//		}
//	}
//}

func Send(msg string) {
	_url := "https://fm.missevan.com/api/chatroom/message/send"

	data, _ := json.Marshal(message{
		RoomID:    868857167,
		Message:   msg,
		MessageID: util.MessageID(),
	})

	header := http.Header{}
	header.Set("content-type", "application/json; charset=UTF-8")
    token:="token=63593d95b3e102581c4899b8%7C3b5d4cad765566d3b2a9544d5812176c%7C1666792853%7Cfbc2ab95b9f2bcb3; buvid3=ED15C3A3-3E0B-08AA-FA6D-367FCFB81D6883497infoc; b_nut=1666793284; buvid4=D13C4122-3B5A-56C4-07BC-E97C47BC46C283497-022102622-9TO+hWhxQHPBVhCkwMXptgnPD6cyNzTRnn7nizjHfux88+95K89E/w%3D%3D; buvid_fp=4f0533d6bf66ff724674dfe102a5fd40; Hm_lvt_91a4e950402ecbaeb38bd149234eb7cc=1666792637,1667200919; FM_SESS=20221031|78jc96g9j36sockp3nq69uvin; FM_SESS.sig=r3CX_GYJ3TsXC_JXx88e2XETbhI; b_lsid=1F6105103F_1842CEB94C8; MSESSID=ftogo1d6rna5qoi7krjomsabg2; Hm_lpvt_91a4e950402ecbaeb38bd149234eb7cc=1667204804"
	req := util.NewRequest(_url, header,token , data)
	body, err := req.Post()
	if err != nil {
		util.Error("send message failed", err)
		return
	}
	str := string(body)
	if str == `{"code":500150022,"info":"聊天内容含有违规信息"}` {
		Send("检测到屏蔽词，消息被屏蔽了哦～")
		return
	}
	//zap.S().Debug(room.Log(str, nil))
}


// SendMessage send a private message to a user according to uid.
func SendMessage(uid int, content string) (ret []byte, err error) {
	_url := "https://www.missevan.com/mperson/sendmessage"

	data := []byte(fmt.Sprintf("user_id=%d&content=%s", uid, content))

	header := http.Header{}
	header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

	token:="token=63593d95b3e102581c4899b8%7C3b5d4cad765566d3b2a9544d5812176c%7C1666792853%7Cfbc2ab95b9f2bcb3; buvid3=ED15C3A3-3E0B-08AA-FA6D-367FCFB81D6883497infoc; b_nut=1666793284; buvid4=D13C4122-3B5A-56C4-07BC-E97C47BC46C283497-022102622-9TO+hWhxQHPBVhCkwMXptgnPD6cyNzTRnn7nizjHfux88+95K89E/w%3D%3D; buvid_fp=4f0533d6bf66ff724674dfe102a5fd40; Hm_lvt_91a4e950402ecbaeb38bd149234eb7cc=1666792637,1667200919; FM_SESS=20221031|78jc96g9j36sockp3nq69uvin; FM_SESS.sig=r3CX_GYJ3TsXC_JXx88e2XETbhI; b_lsid=1F6105103F_1842CEB94C8; MSESSID=ftogo1d6rna5qoi7krjomsabg2; Hm_lpvt_91a4e950402ecbaeb38bd149234eb7cc=1667204804"
	req := util.NewRequest(_url, header,token , data)
	ret, err = req.Post()
	return
}