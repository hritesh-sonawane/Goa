package main

import "fmt"

// O(n^2) time | O(1) space
func TwoNumberSum(arr []int, target int) []int {
	for i := 0; i < len(arr); i++ {
		firstNum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			secondNum := arr[j]
			if firstNum + secondNum == target {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

func main() {
	testArr := []int{3, 5, -4, 8, 11, 1, -1, 6} 
	fmt.Println(TwoNumberSum(testArr, 10))
}