package main

import "fmt"

func main() {
	fmt.Println(isSorted(nil))
}

func isSorted(strings []string) bool {
	lenStrings := len(strings)
	if strings == nil {
		return false
	}

	for index, value := range strings {
		index++
		for index < lenStrings {
			if strings[index] < value {
				return false
			}
			index++
		}
	}

	return true
}
