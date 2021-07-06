package main

import "fmt"

// O(n+m) time | O(c) space
// n => num of characters | m => length of doc
// c => unique no. of characters in characters string
func GenerateDocument(characters string, document string) bool {
	charCount := map[rune]int{}

	for _, character := range characters {
		charCount[character] = charCount[character] + 1
	}

	for _, character := range document {
		if charCount[character] == 0 {
			return false
		}
		charCount[character] = charCount[character] - 1
	}
	return true
}

func main() {
	fmt.Println(GenerateDocument("holelhol", "hello"))
}