package main

import "fmt"

// Best: O(Nlog(N)) time | O(log(N)) space
// Average: O(N(log(N))) time | O(log(N)) space
// Worst: O(N^2) time | O(log(N)) space
func QuickSort(array []int) []int {
	return helper(array, 0, len(array)-1)
}

func helper(array []int, start, end int) []int {
	if start >= end {
		return array
	}

	pivot := start
	left := start + 1
	right := end

	for right >= left {
		if array[left] > array[right] && array[right] < array[pivot] {
			array[left], array[right] = array[right], array[left]
		}
		if array[left] <= array[pivot] {
			left += 1
		}
		if array[right] >= array[pivot] {
			right -= 1
		}
	}

	array[pivot], array[right] = array[right], array[pivot]

	if right-1-start < end-(right+1) {
		helper(array, start, right-1)
		helper(array, right+1, end)
	} else {
		helper(array, right+1, end)
		helper(array, start, right-1)
	}

	return array
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(QuickSort(array))
}