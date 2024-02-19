package graceful

import (
	"context"
	"syscall"
	"testing"
	"time"
)

func TestGraceful(t *testing.T) {
	Add("test1", func() {
		t.Log("test1 callback")
	})

	Run("test2", func(ctx context.Context) {
		t.Log("test2 run")
	}, nil)

	Run("test3", func(ctx context.Context) {
		t.Log("test3 run")
	}, func() {
		t.Log("test3 cleaner")
	})

	go func() {
		time.Sleep(time.Second)
		sig <- syscall.SIGINT
	}()

	Wait()
}
