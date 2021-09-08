package main

import "fmt"

// O(n) time | O(n) space
// n is the total no. of elements in the array
func SpiralTraverse(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}

	result := []int{}
	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1

	for startRow <= endRow && startCol <= endCol {
		for col := startCol; col <= endCol; col++ {
			result = append(result, array[startRow][col])
		}

		for row := startRow + 1; row <= endRow; row++ {
			result = append(result, array[row][endCol])
		}

		for col := endCol - 1; col >= startRow; col-- {
			if startRow == endRow {
				break
			}
			result = append(result, array[endRow][col])
		}

		for row := endRow - 1; row > startRow; row-- {
			if startCol == endCol {
				break
			}
			result = append(result, array[row][startCol])
		}

		startRow++
		endRow--
		startCol++
		endCol--
	}
	return result
}

func main() {
	arr := [][]int{
		{1,  2,  3,  4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10,  9,  8, 7},
	}
	fmt.Println(SpiralTraverse(arr))
}