package logger

import (
	"log"
	"os"
)

var (
	Logger = log.New(os.Stderr, "", log.LstdFlags)
)

func Println(v ...any) {
	if Logger == nil {
		return
	}
	Logger.Println(v...)
}
