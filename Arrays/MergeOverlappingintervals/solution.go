package main

import (
	"fmt"
	"sort"
)

// O(nlog(n)) time | O(n) space
func MergeOverlappingIntervals(intervals [][]int) [][]int {
	sortedIntervals := make([][]int, len(intervals))
	copy(sortedIntervals, intervals)
	sort.Slice(sortedIntervals, func(i, j int) bool {
		return sortedIntervals[i][0] < sortedIntervals[j][0]
	})

	mergedIntervals := make([][]int, 0)
	currentInterval := sortedIntervals[0]
	mergedIntervals = append(mergedIntervals, currentInterval)

	for _, nextInterval := range sortedIntervals {
		currentIntervalEnd := currentInterval[1]
		nextIntervalStart, nextIntervalEnd := nextInterval[0], nextInterval[1]

		if currentIntervalEnd >= nextIntervalStart {
			currentInterval[1] = max(currentIntervalEnd, nextIntervalEnd)
		} else {
			currentInterval = nextInterval
			mergedIntervals = append(mergedIntervals, currentInterval)
		}
	}
	return mergedIntervals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {4, 7}, {6, 8}, {9, 10}}
	fmt.Println(MergeOverlappingIntervals(intervals))
}