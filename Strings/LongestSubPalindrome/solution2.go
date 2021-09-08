package main

import "fmt"

type substring struct {
	left int
	right int
}

func (ss substring) length() int {  //length method
	return ss.right - ss.left
}

// O(n^2) time | O(n) space
func LongestPalindromicSubstring(str string) string {
	result := substring{0, 1}
	for i := 1; i < len(str); i++ {
		odd := getLongestPalindromeFrom(str, i-1, i+1)
		even := getLongestPalindromeFrom(str, i-1, i)
		longest := even
		if odd.length() > even.length() {
			longest = odd
		}
		if longest.length() > result.length() {
			result = longest
		}
	}
	return str[result.left:result.right]
}

func getLongestPalindromeFrom(str string, leftIdx, rightIdx int) substring {
	for leftIdx >= 0 && rightIdx < len(str) {
		if str[leftIdx] != str[rightIdx] {
			break
		}
		leftIdx -= 1
		rightIdx += 1
	}
	return substring{leftIdx+1, rightIdx}
}

func main() {
	string := "abaxyzzyxf"
	fmt.Println(LongestPalindromicSubstring(string))
}