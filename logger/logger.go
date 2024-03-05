package logger

import (
	"log"
	"os"
)

var (
	L = log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
)

func SetLogger(l *log.Logger) {
	L = l
}
