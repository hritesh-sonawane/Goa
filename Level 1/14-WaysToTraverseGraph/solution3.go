package main

import "fmt"

// O(n+m) time | O(1) space
func NumberOfWaysToTraverseGraph(width int, height int) int {
	xDistanceToCorner := width - 1
	yDistanceToCorner := height - 1

	// Number of permutations of right and down movements
	// is the no. of ways to reach the bottom right corner
	numerator := factorial(xDistanceToCorner + yDistanceToCorner)
	denominator := factorial(xDistanceToCorner) * factorial(yDistanceToCorner)
	return numerator / denominator
}

func factorial(num int) int {
	res := 1
	for n := 2; n < num+1; n++ {
		res *= n
	}
	return res
}

func main() {
	fmt.Println(NumberOfWaysToTraverseGraph(5, 5))
}