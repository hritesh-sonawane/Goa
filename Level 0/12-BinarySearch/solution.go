package main

import "fmt"

// O(log(n)) time | O(1) space
func BinarySearch(array []int, target int) int {
	return helper(array, target, 0, len(array)-1)
}

func helper(array []int, target, left, right int) int {
	for left <= right {
		middle := (left + right) / 2
		potentialMatch := array[middle]
		if target == potentialMatch {
			return middle
		} else if target < potentialMatch {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	arr := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	target := 33
	fmt.Println(BinarySearch(arr, target))
}