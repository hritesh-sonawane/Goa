package main

import (
	"fmt"
	"sort"
)

func ClassPhotos(redArr, blueArr []int) bool {
	sort.Slice(redArr, func(i, j int) bool {
		return redArr[i] > redArr[j]
	})
	sort.Slice(blueArr, func(i, j int) bool {
		return blueArr[i] > blueArr[j]
	})

	shirtColorInFirstRow := "BLUE"
	if redArr[0] < blueArr[0] {
		shirtColorInFirstRow = "RED"
	}

	for idx := range redArr {
		redHeight := redArr[idx]
		blueHeight := blueArr[idx]

		if shirtColorInFirstRow == "RED" {
			if redHeight >= blueHeight {
				return false
			}
		} else {
			if blueHeight >= redHeight {
				return false
			}
		}
	}
	return true
}

func main() {
	redArr := []int{5, 8, 1, 3, 4}
	blueArr := []int{6, 9, 2, 4, 5}
	fmt.Println(ClassPhotos(redArr, blueArr))
}