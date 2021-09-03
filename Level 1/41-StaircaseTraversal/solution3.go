package main

import "fmt"

// O(n*k) time | O(n) space
func StaircaseTraversal(height, maxSteps int) int {
	waysToTop := make([]int, height+1)
	waysToTop[0] = 1
	waysToTop[1] = 1

	for curretHeight := 2; curretHeight < height+1; curretHeight++ {
		step := 1
		for step <= maxSteps && step <= curretHeight {
			waysToTop[curretHeight] = waysToTop[curretHeight] + waysToTop[curretHeight - step]
			step += 1
		}
	}

	return waysToTop[height]
}

func main() {
	height := 4
	maxSteps := 2
	fmt.Println(StaircaseTraversal(height, maxSteps))
}