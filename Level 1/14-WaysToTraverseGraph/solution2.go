package main

import "fmt"

// O(n*m) time | O(n*m) space
func NumberOfWaysToTraverseGraph(width int, height int) int {
	numberOfWays := make([][]int, height+1)
	for i := range numberOfWays {
		numberOfWays[i] = make([]int, width+1)
	}

	for widthIdx := 1; widthIdx < width+1; widthIdx++ {
		for heightIdx := 1; heightIdx < height+1; heightIdx++ {
			if widthIdx == 1 || heightIdx == 1 {
				numberOfWays[heightIdx][widthIdx] = 1
			} else {
				waysLeft := numberOfWays[heightIdx][widthIdx-1]
				waysUp := numberOfWays[heightIdx-1][widthIdx]
				numberOfWays[heightIdx][widthIdx] = waysLeft + waysUp
			}
		}
	}
	return numberOfWays[height][width]
}

func main() {
	fmt.Println(NumberOfWaysToTraverseGraph(4, 3))
}