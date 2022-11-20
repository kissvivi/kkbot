package kkbot

import (
	"fmt"
	"testing"
)

func TestSendMessage(t *testing.T) {
	// 20935629 xbc
	// 4004102 xm
	cc,_:=SendMessage(4004102,"test--hahaha")
	fmt.Println(cc)
}
