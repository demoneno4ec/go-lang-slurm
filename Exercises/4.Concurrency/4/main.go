package main

import "fmt"

var (
	a = []int{1, 3, 5, 7, 9}
	b = []int{2, 4, 6, 8, 10}
	c = []int{}
)

func main() {
	realMain(a, b, &c)
	fmt.Println(c)
}

func realMain(a, b []int, c *[]int) {
	// TODO: код писать здесь
}
