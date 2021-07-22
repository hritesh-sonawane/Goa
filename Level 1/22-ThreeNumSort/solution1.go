package main

import "fmt"

// Better understandable solution
// O(n) time | O(1) space
func ThreeNumberSort(array []int, order []int) []int {
	firstValue, thirdValue := order[0], order[2]

	firstIdx := 0
	for idx := range array {
		if array[idx] == firstValue {
			array[firstIdx], array[idx] = array[idx], array[firstIdx]
			firstIdx += 1
		}
	}

	thirdIdx := len(array) - 1
	for idx := len(array) - 1; idx >= 0; idx-- {
		if array[idx] == thirdValue {
			array[thirdIdx], array[idx] = array[idx], array[thirdIdx]
			thirdIdx -= 1
		}
	}
	return array
}

func main() {
	arr := []int{1, 0, 0, -1, -1, 0, 1, 1}
	order := []int{0, 1, -1}
	fmt.Println(ThreeNumberSort(arr, order))
}