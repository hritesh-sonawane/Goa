package main

import (
	"fmt"
)

// O(2^(n+m)) time | O(n+m) space | n: width, m: height
func NumberOfWaysToTraverseGraph(width int, height int) int {
	if width == 1 || height == 1 {
		return 1
	}
	return NumberOfWaysToTraverseGraph(width-1, height) + NumberOfWaysToTraverseGraph(width, height-1)
}

func main() {
	fmt.Println(NumberOfWaysToTraverseGraph(4, 3))
}