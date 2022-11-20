package kkbot

import (
	"fmt"
	"github.kissvivi.kkbot/util"
)

func HandleFmMessage(msg FmMessage) {
	switch msg.Type {
	case TypeMember:
		handleMember(msg)
	case TypeMessage:
		HandleMessage(msg)
	}
}

// ------------------------------------------------------------------- //

// handleMember 处理类型为 Member 的消息
func handleMember(msg FmMessage) {
	switch msg.Event {
	case EventJoinQueue:
		handleMemberJoinQueue(msg)
	}
}

// HandleMessage 处理类型为 Message 的消息
func HandleMessage(msg FmMessage) {
	switch msg.Event {
	case EventNew:
		handleMessageNewQueue(msg)
	}
}

// ------------------------------------------------------------------- //

var MsgCh = make(chan FmMessage)

func handleMessageNewQueue(msg FmMessage) {
	//for _, v := range msg.Message {
	//	name := v
	//	if name == "" {
	//		// 匿名用户
	//		util.Info("匿名用户进入直播间")
	//	} else {
	//		util.Info(fmt.Sprintf("用户 @%s 进入直播间", name))
	//	}
	//}

	MsgCh <- msg

	util.Info(msg.Message)
}

// ------------------------------------------------------------------- //

func handleMemberJoinQueue(msg FmMessage) {
	for _, v := range msg.Queue {
		name := v.Username
		if name == "" {
			// 匿名用户
			util.Info("匿名用户进入直播间")
			//Send(fmt.Sprintf("欢迎 匿名大佬进入直播间"))
			Send(fmt.Sprintf("⚠️检测到一个匿名的小耳朵闯进了直播间，或许也闯进了我的心(^///^)​​"))
		} else {
			pyname, _ := util.Pinyin(name)
			Send(fmt.Sprintf("🎉 欢迎 @%s 今晚的星星和月亮都像你，遥不可及地好看。♥₍๐•ᴗ•๐₎♥　\n 👏 欢迎 [%s] 😽😽😽", name, pyname))
			util.Info(fmt.Sprintf("用户 @%s 进入直播间", name))
		}
	}
}
