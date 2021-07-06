package main

import "fmt"

// O(n^2) time | O(1) space
func FirstNonRepeatingCharacter(str string) int {
	for idx := range str {
		var foundDuplicate = false
		for idx2 := range str {
			if str[idx] == str[idx2] && idx != idx2 {
				foundDuplicate = true
			}
		}
		if !foundDuplicate {
			return idx
		}
	}
	return -1
}

func main() {
	fmt.Println(FirstNonRepeatingCharacter("abcdcaf"))
}