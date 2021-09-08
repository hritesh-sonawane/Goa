package main

import "fmt"

// Best: O(n^2) time | O(1) space
// Average: O(n^2) time | O(1) space
// Worst: O(n^2) time | O(1) space
func SelectionSort(array []int) []int {
	idx := 0
	for idx < len(array)-1 {
		smallestIdx := idx
		for i := idx+1; i < len(array); i++ {
			if array[smallestIdx] > array[i] {
				smallestIdx = i
			}
		}
		array[idx], array[smallestIdx] = array[smallestIdx], array[idx]
		idx += 1
	}
	return array
}

func main() {
	arr := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(SelectionSort(arr))
}