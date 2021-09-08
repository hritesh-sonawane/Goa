package main

import (
	"fmt"
	"sort"
)

// O(w*n*log(n)) time | O(wn) space
// w: no. of words | n: len of longest word
func GroupAnagrams(words []string) [][]string {
	anagrams := map[string][]string{}
	
	for _, word := range words {
		sortedWord := sortWord(word)
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}
	
	result := [][]string{}
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}

func sortWord(word string) string {
	wordBytes := []byte(word)
	sort.Slice(wordBytes, func(i, j int) bool {
		return wordBytes[i] < wordBytes[j]
	})
	return string(wordBytes)
}

func main() {
	words := []string{"yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"}
	fmt.Println(GroupAnagrams(words))
}