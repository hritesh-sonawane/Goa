package main

import (
	"fmt"
	"math"
)

// O(n) time | O(n) space
func MaximizeExpression(array []int) int {
	if len(array) < 4 {
		return 0
	}

	maxOfA := []int{array[0]}
	maxOfAMinusB := []int{math.MinInt32}
	maxOfAMinusBPlusC := []int{math.MinInt32, math.MinInt32}
	maxOfAMinusBPlusCMinusD := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	for idx := 1; idx < len(array); idx++ {
		currentMax := max(maxOfA[idx-1], array[idx])
		maxOfA = append(maxOfA, currentMax)
	}

	for idx := 1; idx < len(array); idx++ {
		currentMax := max(maxOfAMinusB[idx-1], maxOfA[idx-1] - array[idx])
		maxOfAMinusB = append(maxOfAMinusB, currentMax)
	}

	for idx := 2; idx < len(array); idx++ {
		currentMax := max(maxOfAMinusBPlusC[idx-1], maxOfAMinusB[idx-1] + array[idx])
		maxOfAMinusBPlusC = append(maxOfAMinusBPlusC, currentMax)
	}

	for idx := 3; idx < len(array); idx++ {
		currentMax := max(maxOfAMinusBPlusCMinusD[idx-1], maxOfAMinusBPlusC[idx-1] - array[idx])
		maxOfAMinusBPlusCMinusD = append(maxOfAMinusBPlusCMinusD, currentMax)
	}

	return maxOfAMinusBPlusCMinusD[len(maxOfAMinusBPlusCMinusD)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	array := []int{3, 6, 1, -3, 2, 7}
	fmt.Println(MaximizeExpression(array))
}