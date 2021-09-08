package main

import "fmt"

// Naive approach
// O(n) time | O(1) space
func IndexEqualsValue(array []int) int {
	for index, value := range array {
		if index == value {
			return index
		}
	}
	return -1
}

func main() {
	array := []int{-5, -3, 0, 3, 4, 5, 9}
	fmt.Println(IndexEqualsValue(array))
}