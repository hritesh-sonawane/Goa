package main

import "fmt"

// O(log(n)) time | O(1) space
func IndexEqualsValue(array []int) int {
	leftIdx := 0
	rightIdx := len(array) - 1

	for leftIdx <= rightIdx {
		middleIdx := leftIdx + (rightIdx - leftIdx)/2
		middleValue := array[middleIdx]

		if middleValue < middleIdx {
			leftIdx = middleIdx + 1
		} else if middleValue == middleIdx && middleIdx == 0 {
			return middleIdx
		} else if middleValue == middleIdx && array[middleIdx-1] < middleIdx-1 {
			return middleIdx
		} else {
			rightIdx = middleIdx - 1
		}
	}
	return -1
}

func main() {
	array := []int{-5, -3, 0, 3, 4, 5, 9}
	fmt.Println(IndexEqualsValue(array))
}