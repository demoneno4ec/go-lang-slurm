package main

import "testing"

func Test_isSorted(t *testing.T) {
	tests := []struct {
		name string
		ww   []string
		want bool
	}{
		{"sorted", []string{"a", "ab", "b", "bc", "c", "cd"}, true},
		{"sorted2", []string{"a", "aawerf", "ab", "adwerb", "b", "bb", "bc", "c", "cd", "foo", "fooagra"}, true},
		{"sorted_single", []string{"zorro"}, true},
		{"nil_expected false", nil, false},
		{"unsorted в середине", []string{"a", "ab", "b", "cc", "bc", "c", "cd"}, false},
		{"unsorted2 в нескольких местах", []string{"a", "ab", "b", "bc", "c", "cd", "foo", "bar", "abcd"}, false},
		{"unsorted по-русски", []string{"А", "и", "Б", "сидели", "на", "трубе"}, false},
		// TODO: можно добавить тестовых случаев на свое усмотрение
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSorted(tt.ww); got != tt.want {
				t.Errorf("isSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
