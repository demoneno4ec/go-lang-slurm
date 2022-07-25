package main

import "testing"

func Test_doubleDetector(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want bool
	}{
		{"тест где все уникально, ожидаем false", []int{1, 2, 3, 4, 5, 6, 7}, false},
		{"тест дубль в самом конце, ожидаем true", []int{1, 2, 3, 4, 5, 6, 7, 1}, true},
		{"тест дубль в середине, ожидаем true", []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}, true},
		{"тест дубль по краям [1, 2, 1], ожидаем true", []int{1, 2, 1}, true},
		{"тест дубль по краям [1, 1], ожидаем true", []int{1, 1}, true},
		{"тест для [1], ожидаем false", []int{1}, false},
		{"тест для nil, ожидаем false", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doubleDetector(tt.src); got != tt.want {
				t.Errorf("doubleDetector() = %v, want %v", got, tt.want)
			}
		})
	}
}
