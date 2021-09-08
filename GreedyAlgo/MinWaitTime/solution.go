package main

import (
	"fmt"
	"sort"
)

// O(nlogn) time | O(1) space - where n is the no. of queries
func MinWaitTime(queries []int) int {
	sort.Ints(queries)   // [1, 2, 2, 3, 6]

	waitTime := 0
	for idx, val := range queries {
		queriesLeft := len(queries) - (idx + 1)
		waitTime += val * queriesLeft
	}
	return waitTime   // 4 + 6 + 4 + 3 = 17
}

func main() {
	q := []int{3, 2, 1, 2, 6}
	fmt.Println(MinWaitTime(q))
}