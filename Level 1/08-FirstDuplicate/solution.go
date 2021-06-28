package main

import "fmt"

// O(n^2) time | O(1) space
func FirstDuplicate(array []int) int {
	minIdx := len(array)
	for i := 0; i < len(array); i++ {
		value := array[i]
		for j := i+1; j < len(array); j++ {
			valToCompare := array[j]
			if value == valToCompare {
				minIdx = min(minIdx, j)
			}
		}
	}
	if minIdx == len(array) {
		return -1
	}
	return array[minIdx]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{2, 1, 5, 2, 3, 3, 4}
	fmt.Println(FirstDuplicate(arr))
}