package singleton

import "sync"

type Singleton[T any] struct {
	instance *T
	once     *sync.Once
	newer    func() *T
}

func (s *Singleton[T]) Get() *T {
	if s.instance == nil {
		s.once.Do(func() {
			s.instance = s.newer()
		})
	}
	return s.instance
}

func NewSingleton[T any](newer func() *T) *Singleton[T] {
	return &Singleton[T]{
		once:  &sync.Once{},
		newer: newer,
	}
}
