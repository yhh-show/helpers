package errs

import (
	"fmt"
	"github.com/yhh-show/helpers/jsons"
	"github.com/yhh-show/helpers/logger"
	"runtime"
)

type Reporter func(err error, args ...any)

var (
	reporter Reporter
)

func SetReporter(r Reporter) {
	reporter = r
}

func Report(err error, args ...any) bool {
	if err == nil {
		return false
	}

	_, file, line, _ := runtime.Caller(1)
	a := jsons.ToString(args)
	logger.Println(fmt.Sprintf("file: %s:%d", file, line), "| err:", err, "| args:", a)

	if reporter != nil {
		reporter(err, args...)
	}

	return true
}

func Loe(err error, args ...any) bool {
	if err != nil {
		Report(err, args...)
		return true
	}
	return false
}

func Poe(err error, args ...any) bool {
	if err != nil {
		Report(err, args...)
		panic(err)
	}
	return false
}
