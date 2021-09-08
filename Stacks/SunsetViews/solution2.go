package main

import "fmt"

// Better understandable solution
// O(n) time | O(n) space
func SunsetViews(buildings []int, direction string) []int {
	buildingsWithSunsetViews := make([]int, 0)

	startIdx := len(buildings) - 1
	step := -1
	if direction == "WEST" {
		startIdx = 0
		step = 1
	}

	idx := startIdx
	runningMaxHeight := 0
	for idx >= 0 && idx < len(buildings) {
		buildingHeight := buildings[idx]

		if buildingHeight > runningMaxHeight {
			buildingsWithSunsetViews = append(buildingsWithSunsetViews, idx)
		}

		runningMaxHeight = max(runningMaxHeight, buildingHeight)
		idx += step
	}

	if direction == "EAST" {
		reverse(buildingsWithSunsetViews)
	}

	return buildingsWithSunsetViews
}

func reverse(array []int) {
	l := len(array) - 1
	for i := 0; i < len(array)/2; i++ {
		array[i], array[l-i] = array[l-i], array[i]
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	buildings := []int{3, 5, 4, 4, 3, 1, 3, 2}
	direction := "EAST"
	fmt.Println(SunsetViews(buildings, direction))
}