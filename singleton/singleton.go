package singleton

import (
	"github.com/yhh-show/helpers/graceful"
	"reflect"
	"sync"
)

// Singleton 单例模式
type Singleton[T any] struct {
	instance *T
	once     *sync.Once
	newer    func() *T
	cleaner  func(*T)
}

func (s *Singleton[T]) Get() *T {
	if s.instance == nil {
		s.once.Do(func() {
			s.instance = s.newer()
			graceful.Add("singleton "+reflect.TypeOf(s.instance).String(), func() {
				if s.cleaner != nil {
					s.cleaner(s.instance)
				}
			})
		})
	}
	return s.instance
}

func NewSingleton[T any](newer func() *T, cleaner func(*T)) *Singleton[T] {
	return &Singleton[T]{
		once:    &sync.Once{},
		newer:   newer,
		cleaner: cleaner,
	}
}
