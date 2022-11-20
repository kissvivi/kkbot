package util

import (
	"fmt"
	"testing"
)

func TestPinyin(t *testing.T) {
	py,_:=Pinyin("别问我为什么")
	fmt.Println(fmt.Sprintf("欢迎 @%s 今晚的星星和月亮都像你，遥不可及地好看。♥₍๐•ᴗ•๐₎♥　\n",py))
}
