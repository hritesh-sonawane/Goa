package main

import "fmt"

// O(n*2^n) time | O(n*2^n) space
func Powerset(array []int) [][]int {
	subsets := [][]int{{}}
	for _, ele := range array {
		length := len(subsets)
		for i := 0; i < length; i++ {
			currentSubset := subsets[i]
			newsubset := append([]int{}, currentSubset...)
			newsubset = append(newsubset, ele)
			subsets = append(subsets, newsubset)
		}
	}
	return subsets
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(Powerset(arr))
}