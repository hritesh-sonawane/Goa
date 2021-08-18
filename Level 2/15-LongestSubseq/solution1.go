package main

import "fmt"

// Naive approach
// O(mn*min(m, n)) time | O(mn*min(m, n)) space
func LongestCommonSubsequence(s1, s2 string) []byte {
	lcs := make([][][]byte, len(s2)+1)
	for i := range lcs {
		lcs[i] = make([][]byte, len(s1)+1)
	}
	for i := 1; i < len(lcs); i++ {
		for j := 1; j < len(lcs[i]); j++ {
			if s2[i-1] == s1[j-1] {
				tmp := make([]byte, len(lcs[i-1][j-1]))
				copy(tmp, lcs[i-1][j-1])
				lcs[i][j] = append(tmp, s2[i-1])
			} else {
				if len(lcs[i-1][j]) < len(lcs[i][j-1]) {
					lcs[i][j] = lcs[i][j-1]
				} else {
					lcs[i][j] = lcs[i-1][j]
				}
			}
		}
	}
	return lcs[len(s2)][len(s1)]
}

func main() {
	s1 := "ZXWYZW"
	s2 := "XKYKZPW"
	fmt.Println(LongestCommonSubsequence(s1, s2))
}