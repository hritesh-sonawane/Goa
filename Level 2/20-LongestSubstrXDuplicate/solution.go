package main

import "fmt"

type substring struct {
	left  int
	right int
}

func (ss substring) length() int { return ss.right - ss.left}

// O(n) time | O(min(n, a)) space
func LongestSubstrXDuplicate(str string) string {
	lastSeen := map[rune]int{}
	longest := substring{0, 1}
	startIndex := 0

	for i, char := range str {
		if seenIndex, found := lastSeen[char]; found && startIndex < seenIndex+1 {
			startIndex = seenIndex + 1
		}
		if longest.length() < i+1-startIndex {
			longest = substring{startIndex, i+1}
		}
		lastSeen[char] = i
	}
	return str[longest.left:longest.right]
}

func main() {
	str := "abcdefghisacap"
	fmt.Println(LongestSubstrXDuplicate(str))
}