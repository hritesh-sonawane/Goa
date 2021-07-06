package main

import "fmt"

// O(m*(n+m)) time | O(1) space
// n => num of characters | m => length of doc
func GenerateDocument(characters string, document string) bool {
	for _, character := range document {
		docFreq := countCharFreq(character, document)
		charFreq := countCharFreq(character, characters)
		if docFreq > charFreq {
			return false
		}
	}
	return true
}

func countCharFreq(character rune, target string) int {
	var frequency = 0
	for _, char := range target {
		if char == character {
			frequency += 1
		}
	}
	return frequency
}

func main() {
	fmt.Println(GenerateDocument("holelhol", "hello"))
}