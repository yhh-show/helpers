package safego

import (
	"fmt"
	"github.com/yhh-show/helpers/errs"
)

// Go 防止panic的goroutine
func Go(fn func()) chan any {
	ch := make(chan any, 1)

	go func() {
		defer func() {
			r := recover()
			ch <- r
			if r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				errs.Report(err, "safego.Go panic")
			}
		}()

		fn()
	}()

	return ch
}
