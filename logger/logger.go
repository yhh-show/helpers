package logger

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
)

func SetLogger(l *log.Logger) {
	if l == nil {
		return
	}
	logger = l
}

func Print(v ...any) {
	logger.Print(v...)
}

func Printf(format string, v ...any) {
	logger.Printf(format, v...)
}

func Println(v ...any) {
	logger.Println(v...)
}

func Fatal(v ...any) {
	logger.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	logger.Fatalf(format, v...)
}

func Fatalln(v ...any) {
	logger.Fatalln(v...)
}

func Panic(v ...any) {
	logger.Panic(v...)
}

func Panicf(format string, v ...any) {
	logger.Panicf(format, v...)
}

func Panicln(v ...any) {
	logger.Panicln(v...)
}
