package util

import (
	"time"

	"github.com/fatih/color"
)

func Info(s string) {
	color.Cyan("[KKBot %s] %s\n", time.Now().Format("15:04:05"), s)
}

func Error(s string, err error) {
	color.Red("[KKBot %s] %s: %e\n", time.Now().Format("15:04:05"), s, err)
}
