package main

import (
	"fmt"
	"testing"
)

func Test_createList(t *testing.T) {
	tests := []struct {
		name    string
		sl      []int
		wantLen int
	}{
		{
			"тест 1-2-3",
			[]int{1, 2, 3},
			3,
		},
		{
			"тест 0",
			[]int{0},
			1,
		},
		{
			"тест nil",
			nil,
			0,
		},
		{
			"тест 27 чисел",
			[]int{-10, -8, -8, -5, -5, 7, 7, 8, 8, 8, 9, 9, 9, 10, 10, 11, 12, 15, 20, 29, 39, 49, 533, 555, 577, 599, 800},
			27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createList(tt.sl)
			if length(got) != tt.wantLen {
				t.Errorf("createList length = %d, want %d", length(got), tt.wantLen)
				return
			}

			if eq, reason := checkVal(got, tt.sl); !eq {
				t.Errorf("createList made incorrect list, reason %s", reason)
			}
		})
	}
}

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name      string
		nodes     []int
		val       int
		wantCount int
	}{
		{"тест 1, 2, 2, 2, 3 тогда цифр 1 в кол-ве 1", []int{1, 2, 2, 2, 3, 3}, 1, 1},
		{"тест 1, 2, 2, 2, 3 тогда цифр 2 в кол-ве 1", []int{1, 2, 2, 2, 3, 3}, 2, 1},
		{"тест -10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7 тогда цифр 2 в кол-ве 1", []int{-10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7}, 2, 1},
		{"тест -10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7 тогда цифр 7 в кол-ве 1", []int{-10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7}, 7, 1},
		{"тест -10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7 тогда цифр -10 в кол-ве 1", []int{-10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7}, -10, 1},
		{"тест -10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7 тогда цифр 0 в кол-ве 0", []int{-10, -10, -2, -2, 2, 2, 3, 5, 6, 6, 6, 6, 7, 7}, 0, 0},
		{"тест 0, 0, 0, 0, 0 тогда цифр 0 в кол-ве 1", []int{0, 0, 0, 0, 0}, 0, 1},
		{"тест 5, 5, 5, 5 тогда цифр 5 в кол-ве 1", []int{5, 5, 5, 5}, 5, 1},
		{"тест 1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 4, 5, 5, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9 тогда цифр 5 в кол-ве 1", []int{1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 4, 5, 5, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9}, 5, 1},
		{"nil", nil, 5, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := createList(tt.nodes)
			deleteDuplicates(list)
			if countVal(list, tt.val) != tt.wantCount {
				t.Errorf("after deleteDuplicates countVal(list, %d)=%d, want %d", tt.val, countVal(list, tt.val), tt.wantCount)
			}
		})
	}
}

// helpers

// lastNode returns pointer to last item.
func lastNode(head *ListNode) *ListNode {
	var last *ListNode
	for node := head; node != nil; node = node.Next {
		if node.Next == nil {
			last = node
		}
	}
	return last
}

// countVal returns count of the item with Val=val.
func countVal(head *ListNode, val int) int {
	res := 0
	for node := head; node != nil; node = node.Next {
		if node.Val == val {
			res++
		}
	}
	return res
}

// length returns count of the items.
func length(head *ListNode) int {
	res := 0
	for node := head; node != nil; node = node.Next {
		res++
	}
	return res
}

// checkVal returns flag of equal values in listNode and slice []int.
func checkVal(head *ListNode, sl []int) (bool, string) {
	lenSL := len(sl)
	lenHead := length(head)
	if lenSL != lenHead {
		return false, fmt.Sprintf("heal length=%d <> slice length=%d", lenHead, lenHead)
	}
	i := 0
	for node := head; node != nil; node = node.Next {
		if node.Val != sl[i] {
			return false, fmt.Sprintf("heal[%d]=%d <> sl[%d]=%d", i, node.Val, i, sl[i])
		}
		i++
	}
	return true, ""
}
