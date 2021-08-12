package main

import "fmt"

// O(n*l) time | O(c) space | c: no. of unique char
// n: no. of words | l: len of longest word
func MinCharsForWords(words []string) []string {
	maxCharFreq := map[rune]int{}

	for _, word := range words {
		charFreq := countCharFreq(word)
		updateMaxFreq(charFreq, maxCharFreq)
	}

	return makeArrayFromCharFreq(maxCharFreq)
}

func countCharFreq(str string) map[rune]int {
	charFreq := map[rune]int{}
	for _, character := range str {
		charFreq[character] += 1
	}
	return charFreq
}

func updateMaxFreq(frequencies map[rune]int, maxFrequencies map[rune]int) {
	for character, frequency := range frequencies {
		if maxFrequency, found := maxFrequencies[character]; found {
			maxFrequencies[character] = max(frequency, maxFrequency)
		} else {
			maxFrequencies[character] = frequency
		}
	}
}

func makeArrayFromCharFreq(charFreq map[rune]int) []string {
	characters := make([]string, 0)
	for character, frequency := range charFreq {
		for i := 0; i < frequency; i++ {
			characters = append(characters, string(character))
		}
	}
	return characters
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	words := []string{"this", "that", "did", "deed", "them!", "a"}
	fmt.Println(MinCharsForWords(words))
}