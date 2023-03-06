package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(doubleDetector(nums))
}

func doubleDetector(nums []int) bool {
	lenNums := len(nums)
	for index, value := range nums {
		index++
		for index < lenNums {
			if nums[index] == value {
				return true
			}
			index++
		}
	}

	return false
}
