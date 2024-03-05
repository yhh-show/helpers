package ticker

import (
	"context"
	"time"
)

func Run(ctx context.Context, fn func(), d time.Duration) {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			break
		case <-ticker.C:
			fn()
		}
	}
}
