package safego

// Go 防止panic的goroutine
func Go(fn func()) chan any {
	ch := make(chan any)

	go func() {
		defer func() {
			ch <- recover()
		}()

		fn()
	}()

	return ch
}
