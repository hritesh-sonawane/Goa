package main

import "fmt"

// O(n) time | O(n) space
func SunsetViews(buildings []int, direction string) []int {
	candidateBuildings := make([]int, 0)

	startIdx := len(buildings) - 1
	step := -1
	if direction == "EAST" {
		startIdx = 0
		step = 1
	}

	idx := startIdx
	for idx >= 0 && idx < len(buildings) {
		buildingHeight := buildings[idx]

		for len(candidateBuildings) > 0 && buildings[candidateBuildings[len(candidateBuildings)-1]] <= buildingHeight {
			candidateBuildings = candidateBuildings[:len(candidateBuildings)-1]
		}

		candidateBuildings = append(candidateBuildings, idx)
		idx += step
	}

	if direction == "WEST" {
		reverse(candidateBuildings)
	}

	return candidateBuildings
}

func reverse(array []int) {
	l := len(array) - 1
	for i := 0; i < len(array)/2; i++ {
		array[i], array[l-i] = array[l-i], array[i]
	}
}

func main() {
	buildings := []int{3, 5, 4, 4, 3, 1, 3, 2}
	direction := "EAST"
	fmt.Println(SunsetViews(buildings, direction))
}