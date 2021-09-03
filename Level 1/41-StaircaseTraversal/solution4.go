package main

import "fmt"

/*

height: 4
maxSteps: 2
currentNumberOfWays: 0->1->2->3->5
waysToTop: {1, 1, 2, 3, 5}
currentHeight: 1->2->3->4
startOfWindow: -2->-1->0->1
endOfWindow: 0->1->2->3

*/

// O(n) time | O(n) space
func StaircaseTraversal(height, maxSteps int) int {
	currentNumberOfWays := 0
	waysToTop := []int{1} // dynamic programming:)

	for currentHeight := 1; currentHeight < height+1; currentHeight++ {
		startOfWindow := currentHeight - maxSteps - 1
		endOfWindow := currentHeight - 1

		if startOfWindow >= 0 {
			currentNumberOfWays -= waysToTop[startOfWindow]
		}

		currentNumberOfWays += waysToTop[endOfWindow]
		waysToTop = append(waysToTop, currentNumberOfWays)
	}

	return waysToTop[height]
}

func main() {
	height := 4
	maxSteps := 2
	fmt.Println(StaircaseTraversal(height, maxSteps))
}