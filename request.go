package kkbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/secriy/missevan/live"
)

type message struct {
	RoomID    int64  `json:"room_id"`
	Message   string `json:"message"`
	MessageID string `json:"msg_id"`
}

func sendLiveMessage(rid int64, msg, cookie string) error {
	_url := "https://fm.missevan.com/api/chatroom/message/send"

	// Must use the JSON marshal to generate the data,
	// otherwise the data will contain unescaped characters.
	data, _ := json.Marshal(message{
		RoomID:    rid,
		Message:   msg,
		MessageID: live.MessageID(),
	})

	client := &http.Client{}
	req, err := http.NewRequest("POST", _url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("origin", "https://www.missevan.com")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.53")
	req.Header.Set("content-type", "application/json; charset=UTF-8")
	req.Header.Set("cookie", cookie)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ret, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	str := string(ret)
	if str == `{"code":500020001,"info":"用户未登录或登录已过期，请重新登录"}` {
		return errors.New("login expired")
	}
	if str == `{"code":500150022,"info":"聊天内容含有违规信息"}` {
		return errors.New("message illegal")
	}
	return nil
}

func baseCookie() string {
	_url := "https://fm.missevan.com/api/user/info"

	resp, err := http.Get(_url)
	if err != nil {
		panic(err)
		return ""
	}
	cookie := strings.Builder{}
	for _, v := range resp.Header.Values("set-cookie") {
		cookie.WriteString(v)
	}
	return cookie.String()
}
