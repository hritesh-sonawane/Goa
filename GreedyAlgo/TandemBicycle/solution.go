package main

import (
	"fmt"
	"sort"
)

// O(nlog(n)) time | O(1) space | n => no. of tandem bicycles
func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)

	if !fastest {
		reverseArray(redShirtSpeeds)
	}

	totalSpeed := 0
	for idx := range redShirtSpeeds {
		rider1 := redShirtSpeeds[idx]
		rider2 := blueShirtSpeeds[len(blueShirtSpeeds)-idx-1]
		totalSpeed += max(rider1, rider2)
	}
	return totalSpeed
}

func reverseArray(array []int) {
	var start = 0
	var end = len(array)-1
	for start < end {
		temp := array[start]
		array[start] = array[end]
		array[end] = temp
		start += 1
		end -= 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr1 := []int{5, 5, 3, 9, 2}
	arr2 := []int{3, 6, 7, 2, 1}
	fmt.Println(TandemBicycle(arr1, arr2, true))
}