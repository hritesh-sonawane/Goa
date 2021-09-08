package main

import "fmt"

// O(mn) time | O(mn) space
func LongestCommonSubsequence(str1, str2 string) []byte {
	lengths := make([][]int, len(str2)+1)
	for i := range lengths {
		lengths[i] = make([]int, len(str1)+1)
	}
	for i := 1; i < len(str2)+1; i++ {
		for j := 1; j < len(str1)+1; j++ {
			if str2[i-1] == str1[j-1] {
				lengths[i][j] = lengths[i-1][j-1] + 1
			} else {
				lengths[i][j] = max(lengths[i-1][j], lengths[i][j-1])
			}
		}
	}
	return buildSequence(lengths, str1)
}

func buildSequence(lengths [][]int, str1 string) []byte {
	sequence := make([]byte, 0)
	i := len(lengths) - 1
	j := len(lengths[0]) - 1
	for i != 0 && j != 0 {
		if lengths[i][j] == lengths[i-1][j] {
			i -= 1
		} else if lengths[i][j] == lengths[i][j-1] {
			j -= 1
		} else {
			sequence = append(sequence, str1[j-1])
			i -= 1
			j -= 1
		}
	}
	return reverse(sequence)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func reverse(array []byte) []byte {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func main() {
	s1 := "ZXWYZW"
	s2 := "XKYKZPW"
	fmt.Println(LongestCommonSubsequence(s1, s2))
}