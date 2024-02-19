package safego

import (
	"errors"
	"testing"
)

func TestGo(t *testing.T) {
	err := errors.New("test3")
	tests := []struct {
		name string
		args func()
		want any
	}{
		{
			name: "test1",
			args: func() {
				panic("test1")
			},
			want: "test1",
		},
		{
			name: "test2",
			args: func() {
			},
			want: nil,
		},
		{
			name: "test3",
			args: func() {
				panic(err)
			},
			want: err,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := <-Go(tt.args)
			if got != tt.want && !errors.Is(got.(error), tt.want.(error)) {
				t.Errorf("Go() = %v, want %v", got, tt.want)
			}
		})
	}
}
