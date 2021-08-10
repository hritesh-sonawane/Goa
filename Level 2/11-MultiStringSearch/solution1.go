package main

import "fmt"

// O(bns) time | O(n) space | n: length of biggest string
// this is the naive approach
func MultiStringSearch(bigString string, smallStrings []string) []bool {
	output := make([]bool, len(smallStrings))
	for i, smallString := range smallStrings {
		output[i] = isInBigString(bigString, smallString)
	}
	return output
}

func isInBigString(bigString, smallString string) bool {
	for i := range bigString {
		if i+len(smallString) > len(bigString) {
			break
		}
		if isInBigStringHelper(bigString, smallString, i) {
			return true
		}
	}
	return false
}

func isInBigStringHelper(bigString, smallString string, startIdx int) bool {
	leftBigIdx := startIdx
	rightBigIdx := startIdx + len(smallString) - 1
	leftSmallIdx := 0
	rightSmallIdx := len(smallString) - 1

	for leftBigIdx <= rightBigIdx {
		if bigString[leftBigIdx] != smallString[leftSmallIdx] || bigString[rightBigIdx] != smallString[rightSmallIdx] {
			return false
		}
		leftBigIdx += 1
		rightBigIdx -= 1
		leftSmallIdx += 1
		rightSmallIdx -= 1
	}
	return true
}

func main() {
	bigString := "this is a big string"
	smallStrings := []string{"this", "yo", "is", "a", "bigger", "string", "kappa"}
	fmt.Println(MultiStringSearch(bigString, smallStrings))
}