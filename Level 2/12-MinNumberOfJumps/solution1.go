package main

import "fmt"

// O(n^2) time | O(n) space | naive approach
func MinNumberOfJumps(array []int) int {
	jumps := make([]int, len(array))
	for i := range jumps {
		jumps[i] = -1
	}
	jumps[0] = 0
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[j] >= i-j && (jumps[i] == -1 || jumps[j]+1 < jumps[i]) {
				jumps[i] = jumps[j] + 1
			}
		}
	}
	return jumps[len(array)-1]
}

func main() {
	array := []int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3}
	fmt.Println(MinNumberOfJumps(array))
}