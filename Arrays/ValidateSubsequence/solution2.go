package main

import "fmt"

// O(n) time | O(1) space
func validate(arr []int, seq []int) bool {
	seqIdx := 0
	for _, value := range arr {
		if seqIdx == len(seq) {
			break
		}
		if value == seq[seqIdx] {
			seqIdx += 1
		}
	}
	return seqIdx == len(seq)
}

func main() {
	arr1 := []int{3, 5, -4, 8, 11, 1, -1, 6} 
	seq1 := []int{8, 11, 6}
	fmt.Println(validate(arr1, seq1))
}