package main

import (
	"fmt"
	"strconv"
)

// O(n) time | O(n) space - where n is the length of string
func RunLengthEncoding(str string) string {
	encodedStringChar := []byte{}
	currentRunLength := 1

	for i := 1; i < len(str); i++ {
		currentChar := str[i]
		prevChar := str[i-1]

		if currentChar != prevChar || currentRunLength == 9 {
			encodedStringChar = append(encodedStringChar, strconv.Itoa(currentRunLength)[0])
			encodedStringChar = append(encodedStringChar, prevChar)
			currentRunLength = 0
		}
		currentRunLength++
	}
	// Handle the last run
	encodedStringChar = append(encodedStringChar, strconv.Itoa(currentRunLength)[0])
	encodedStringChar = append(encodedStringChar, str[len(str)-1])
	return string(encodedStringChar)
}

func main() {
	fmt.Println(RunLengthEncoding("AAAAAAAAAAAAABBCCCCDD"))
}