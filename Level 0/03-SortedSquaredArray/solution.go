package main

import (
	"fmt"
	"sort"
)

// O(nlogn) time | O(n) space - where 'n' is the length of input array
func SortedSqArray(arr []int) []int {
	sortedsquares := make([]int, len(arr))
	for idx, value := range arr {
		sortedsquares[idx] = value * value
	}
	sort.Ints(sortedsquares)
	return sortedsquares
}

func main() {
	arr1 := []int{1, -2, 3, 4, 5, 6} 
	fmt.Println(SortedSqArray(arr1))
}