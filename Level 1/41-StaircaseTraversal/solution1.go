package main

import "fmt"

// O(k^n) time | O(n) space | n: height | k: maxSteps
func StaircaseTraversal(height, maxSteps int) int {
	return numberOfWaysToTop(height, maxSteps)
}

func numberOfWaysToTop(height, maxSteps int) int {
	if height <= 1 {
		return 1
	}

	numberOfWays := 0
	for step := 1; step < min(maxSteps, height)+1; step++ {
		numberOfWays += numberOfWaysToTop(height - step, maxSteps)
	}

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