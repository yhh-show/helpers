package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	Writter io.Writer = os.Stderr
)

func Println(v ...any) {
	if Writter == nil {
		return
	}
	fmt.Fprintln(Writter, append([]interface{}{time.Now().Format(time.RFC3339)}, v...))
}
