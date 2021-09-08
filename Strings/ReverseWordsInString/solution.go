package main

import (
	"fmt"
	"strings"
)

// O(n) time | O(n) space
func ReverseWordsInString(str string) string {
	words := make([]string, 0)
	startOfWord := 0
	for idx, character := range str {
		if character == ' ' {
			words = append(words, str[startOfWord:idx])
			startOfWord = idx
		} else if str[startOfWord] == ' ' {
			words = append(words, " ")
			startOfWord = idx
		}
	}
	words = append(words, str[startOfWord:])
	reverseList(words)
	return strings.Join(words, "")
}

func reverseList(list []string) {
	start := 0
	end := len(list) - 1
	for start < end {
		temp := list[start]
		list[start] = list[end]
		list[end] = temp
		start += 1
		end -= 1
	}
}

func main() {
	str := "Daenerys Stormborn of House Targaryen"
	fmt.Println(ReverseWordsInString(str))
}