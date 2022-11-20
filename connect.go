package kkbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.kissvivi.kkbot/util"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/secriy/missevan/live"
)

type connection struct {
	conn *websocket.Conn
	mu   *sync.Mutex
}

func Connect(input chan<- FmMessage, rid int64) {
	h := http.Header{}
	h.Add("Pragma", "no-cache")
	h.Add("Origin", "https://fm.missevan.com")
	h.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,ja;q=0.5")
	h.Add("User-Agent", "MissEvanApp/4.8.0 (iOS;15.4.1;iPhone14,5)")
	h.Add("Cache-Control", "no-cache")
	h.Add("Cookie", baseCookie())

	dialer := new(websocket.Dialer)
	conn, resp, err := dialer.Dial(fmt.Sprintf("wss://im.missevan.com/ws?room_id=%d", rid), h)
	if err != nil {
		util.Error("消息解码失败", err)
		return
	}
	defer conn.Close()

	if resp.StatusCode != 101 {
		util.Error(fmt.Sprintf("尝试连接失败 (%d)", resp.StatusCode), nil)
		return
	}

	c := &connection{conn, &sync.Mutex{}}
	joinMsg := fmt.Sprintf(`{"action":"join","uuid":"3b65048c-f7b1-41e3-ae88-da2146d24004","type":"room","room_id":%d}`, rid)
	if err := c.conn.WriteMessage(websocket.TextMessage, []byte(joinMsg)); err != nil {
		util.Error("尝试进入直播间失败", err)
		return
	}

	go heartBeat(c) // Keep connected

	util.Info(fmt.Sprintf("连接直播间 %d 成功", rid))

	for {
		msgType, msgData, err := c.conn.ReadMessage()
		if err != nil {
			util.Error("接收消息失败", err)
			return
		}
		if len(msgData) == 0 {
			continue
		}
		// Parse the message according to the message type,
		// most of the time it is a binary message,
		// so we need to decode first.
		switch msgType {
		case websocket.BinaryMessage:
			data, err := live.BrotliDecompress(msgData)
			if err != nil {
				util.Error("消息解码失败", err)
				continue
			}
			msgData = data
		case websocket.TextMessage:
		}

		if string(msgData) == "❤️" {
			continue
		}

		textMsgs, err := parseMessage(msgData)
		if err != nil {
			util.Error("消息解析失败", err)
			continue
		}
		fmt.Printf("%+v\n",textMsgs)
		for _, textMsg := range textMsgs {
			input <- textMsg
		}
	}
}

func parseMessage(msgData []byte) ([]FmMessage, error) {
	// Make sure the JSON object is an array.
	buf := bytes.Buffer{}
	if msgData[0] != '[' {
		buf.WriteByte('[')
		buf.Write(msgData)
		buf.WriteByte(']')
	} else {
		buf.Write(msgData)
	}
	var textMsgs []FmMessage
	if err := json.Unmarshal(buf.Bytes(), &textMsgs); err != nil {
		return nil, err
	}
	return textMsgs, nil
}

// heartBeat sends a heartbeat to the WebSocket connection every 30 seconds.
func heartBeat(c *connection) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		_ = c.conn.WriteMessage(websocket.TextMessage, []byte("❤️"))
		c.mu.Unlock()
	}
}
