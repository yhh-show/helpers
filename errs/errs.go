package errs

import "github.com/yhh-show/helpers/logger"

type Reporter func(err error, args ...any)

var (
	reporter Reporter
)

func SetReporter(r Reporter) {
	reporter = r
}

func Report(err error, args ...any) {
	logger.Println(append(args, err)...)
	if reporter != nil {
		reporter(err, args...)
	}
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
