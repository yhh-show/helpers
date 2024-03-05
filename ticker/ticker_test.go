package ticker

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go Run(ctx, func() {
		fmt.Println(time.Now())
	}, time.Second)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
