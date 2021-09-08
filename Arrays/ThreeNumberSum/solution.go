package main

import (
	"fmt"
	"sort"
)

// O(n^2) time | O(n) space
func ThreeNumberSum(arr []int, target int) [][]int {
	sort.Ints(arr)
	triplets := [][]int{}
	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == target {
				triplets = append(triplets, []int{arr[i], arr[left], arr[right]})
				left += 1
				right -= 1
			} else if sum < target {
				left += 1
			} else if sum > target {
				right -= 1
			}
		}
	}
	return triplets
}

func main() {
	testArr := []int{12, 3, 1, 2, -6, 5, -8, 6} 
	fmt.Println(ThreeNumberSum(testArr, 0))
}