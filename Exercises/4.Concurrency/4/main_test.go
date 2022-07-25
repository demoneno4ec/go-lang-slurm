package main

import (
	"reflect"
	"sync"
	"testing"
)

// в этом тесте проверяется только соответсвие результату
// запуск горутин будет проверяться в ревью
func Test_realMain(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			"тест 1-10",
			[]int{1, 3, 5, 7, 9},
			[]int{2, 4, 6, 8, 10},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			"тест 1-2",
			[]int{1},
			[]int{2},
			[]int{1, 2},
		},
		{
			"тест zero",
			[]int{},
			[]int{},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := make([]int, 0, len(tt.a)+len(tt.b))
			wg := &sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				realMain(tt.a, tt.b, &c)
			}()

			if !reflect.DeepEqual(tt.want, c) {
				t.Errorf("realMain() c = %v, want %v", c, tt.want)
			}
		})
	}
}
