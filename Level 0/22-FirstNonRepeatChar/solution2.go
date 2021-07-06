package main

import "fmt"

// O(n) time | O(1) space
// constant space bcoz string only has lowercase letters
func FirstNonRepeatingCharacter(str string) int {
	charFreq := map[rune]int{}

	for _, char := range str {
		charFreq[char] += 1
	}

	for idx, char := range str {
		if charFreq[char] == 1 {
			return idx
		}
	}
	return -1
}

func main() {
	fmt.Println(FirstNonRepeatingCharacter("acdcaf"))
}