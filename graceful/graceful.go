package graceful

import (
	"context"
	"github.com/yhh-show/helpers/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/yhh-show/helpers/safego"
)

var (
	sig    = make(chan os.Signal, 1)
	fns    []*fnItem
	locker = &sync.Mutex{}
)

type fnItem struct {
	name string
	fn   func()
}

func Wait() {
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig

	logger.Println("graceful shutdown on signal:", s)

	for _, fn := range fns {
		logger.Println("graceful shutdown:", fn.name)
		fn.fn()
	}

	logger.Println("graceful shutdown done")
}

func Add(name string, fn func()) {
	locker.Lock()
	defer locker.Unlock()

	fns = append([]*fnItem{{name, fn}}, fns...)
}

func Run(name string, runner func(context.Context), cleaner func()) {
	ctx, cancel := context.WithCancel(context.Background())

	time.Sleep(time.Millisecond)

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
