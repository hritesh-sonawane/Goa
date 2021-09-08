package main

import "fmt"

// O(n*k) time | O(n) space
func StaircaseTraversal(height, maxSteps int) int {
	return numberOfWaysToTop(height, maxSteps, map[int]int{0:1, 1: 1})
}

func numberOfWaysToTop(height, maxSteps int, memoize map[int]int) int {
	if ways, found := memoize[height]; found {
		return ways
	}

	numberOfWays := 0
	for step := 1; step < min(maxSteps, height)+1; step++ {
		numberOfWays += numberOfWaysToTop(height-step, maxSteps, memoize)
	}
	memoize[height] = numberOfWays

	return numberOfWays
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := 4
	maxSteps := 2
	fmt.Println(StaircaseTraversal(height, maxSteps))
}