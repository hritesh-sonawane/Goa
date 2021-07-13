package main

import (
	"fmt"
	"math"
)

// O(nm) time | O((min(n, m)) space
func LevenshteinDistance(a, b string) int {
	small, big := a, b
	if len(a) > len(b) {
		big, small = small, big
	}
	evenEdits := make([]int, len(small)+1)
	oddEdits := make([]int, len(small)+1)
	var prevEdits, currEdits []int
	for i := range evenEdits {
		evenEdits[i] = i
		oddEdits[i] = math.MinInt32
	}
	for i := 1; i < len(big)+1; i++ {
		if i%2 == 1 {
			currEdits, prevEdits = oddEdits, evenEdits
		} else {
			currEdits, prevEdits = evenEdits, oddEdits
		}
		currEdits[0] = i
		for j := 1; j < len(small)+1; j++ {
			if big[i-1] == small[j-1] {
				currEdits[j] = prevEdits[j-1]
			} else {
				currEdits[j] = 1 + min(prevEdits[j-1], prevEdits[j], currEdits[j-1])
			}
		}
	}
	if len(big)%2 == 0 {
		return evenEdits[len(small)]
	}
	return oddEdits[len(small)]
}

func min(args ...int) int {
	curr := args[0]
	for _, num := range args {
		if curr > num {
			curr = num
		}
	}
	return curr
}

func main() {
	fmt.Println(LevenshteinDistance("abc", "yabd"))
}