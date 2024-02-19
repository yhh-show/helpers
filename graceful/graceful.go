package graceful

import (
	"context"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/yhh-show/helpers/inner/logger"
	"github.com/yhh-show/helpers/safego"
)

var (
	sig    = make(chan os.Signal, 1)
	fns    []*fnItem
	locker = &sync.Mutex{}

	LoggerWritter io.Writer = os.Stderr
)

type fnItem struct {
	name string
	fn   func()
}

func Wait() {
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	logger.Println("graceful shutdown")

	for _, fn := range fns {
		logger.Println("graceful shutdown:", fn.name)
		fn.fn()
	}
}

func Add(name string, fn func()) {
	locker.Lock()
	defer locker.Unlock()

	fns = append([]*fnItem{{name, fn}}, fns...)
}

func Run(name string, runner func(context.Context), cleaner func()) {
	ctx, cancel := context.WithCancel(context.Background())

	safego.Go(func() {
		logger.Println("graceful run:", name)
		runner(ctx)
	})

	Add(name, func() {
		cancel()

		if cleaner != nil {
			cleaner()
		}
	})
}
