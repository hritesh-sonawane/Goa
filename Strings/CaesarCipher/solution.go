package main

import (
	"fmt"
	"strings"
)

// O(n) time | O(n) space
func CaesarCipherEncryptor(str string, key int) string {
	runes := []rune(str)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for i, char := range runes {
		index := strings.Index(alphabet, string(char))
		if index == -1 {
			return "" // Bad Input
		}
		newIndex := (index + key) % 26
		runes[i] = rune(alphabet[newIndex])
	}
	return string(runes)
}

func main() {
	fmt.Println(CaesarCipherEncryptor("abc", 4))
}