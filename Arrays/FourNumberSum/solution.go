package main

import "fmt"

// Average: O(n^2) time | O(n^2) space
// Worst: O(n^3) time | O(n^2) space
func FourNumberSum(array []int, target int) [][]int {
	allPairSums := map[int][][]int{}
	quadruplets := [][]int{}

	for i := 1; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			currentSum := array[i] + array[j]
			difference := target - currentSum
			if pairs, found := allPairSums[difference]; found {
				for _, pair := range pairs {
					newQuad := append(pair, array[i], array[j])
					quadruplets = append(quadruplets, newQuad)
				}
			}
		}
		for k := 0; k < i; k++ {
			currentSum := array[i] + array[k]
			allPairSums[currentSum] = append(allPairSums[currentSum], []int{array[k], array[i]})
		}
	}
	return quadruplets
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	target := 10
	fmt.Println(FourNumberSum(arr, target))
}