package main

import (
	"fmt"
	"sort"
)

// O(nlog(n)) time | O(1) space
func TwoNumberSum(arr []int, target int) []int {
	sort.Ints(arr)
	left, right := 0, len(arr)-1
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == target {
			return []int{arr[left], arr[right]}
		} else if currentSum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}

func main() {
	testArr := []int{3, 5, -4, 8, 11, 1, -1, 6} 
	fmt.Println(TwoNumberSum(testArr, 10))
}