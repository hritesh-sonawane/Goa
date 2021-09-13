package main

import "fmt"

type Trie struct {
	children map[rune]Trie
	word     string
}

func (t Trie) Add(word string) {
	current := t

	for _, letter := range word {
		if _, found := current.children[letter]; !found {
			current.children[letter] = Trie {
				children: map[rune]Trie{},
			}
		}
		current = current.children[letter]
	}

	current.children['*'] = Trie {
		children: map[rune]Trie{},
		word:     word,
	}
}


// O(nm*8^s + ws) time | O(nm + ws) space
// n: width | m: height of board
// w: no. of words | s: len of longest word
func BoggleBoard(board [][]rune, words []string) []string {
	trie := Trie{children: map[rune]Trie{}}

	for _, word := range words {
		trie.Add(word)
	}

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}

	finalWords := map[string]bool{}
	for i := range board {
		for j := range board[i] {
			explore(i, j, board, trie, visited, finalWords)
		}
	}

	result := []string{}
	for word := range finalWords {
		result = append(result, word)
	}

	return result
}

func explore(i, j int, board [][]rune, trie Trie, visited [][]bool, finalWords map[string]bool) {
	if visited[i][j] {
		return
	}

	letter := board[i][j]
	if _, found := trie.children[letter]; !found {
		return
	}

	visited[i][j] = true
	trie = trie.children[letter]
	if end, found := trie.children['*']; found {
		finalWords[end.word] = true
	}

	neighbors := getNeighbors(i, j, board)
	for _, neighbor := range neighbors {
		explore(neighbor[0], neighbor[1], board, trie, visited, finalWords)
	}

	visited[i][j] = false
}

func getNeighbors(i, j int, board [][]rune) [][]int {
	neighbors := [][]int{}

	if i > 0 && j > 0 {
		neighbors = append(neighbors, []int{i-1, j-1})
	}

	if i > 0 && j < len(board[0])-1 {
		neighbors = append(neighbors, []int{i-1, j+1})
	}

	if i < len(board)-1 && j < len(board[0])-1 {
		neighbors = append(neighbors, []int{i+1, j+1})
	}

	if i < len(board)-1 && j > 0 {
		neighbors = append(neighbors, []int{i+1, j-1})
	}

	if i > 0 {
		neighbors = append(neighbors, []int{i-1, j})
	}

	if i < len(board)-1 {
		neighbors = append(neighbors, []int{i+1, j})
	}

	if j > 0 {
		neighbors = append(neighbors, []int{i, j-1})
	}

	if j < len(board[0])-1 {
		neighbors = append(neighbors, []int{i, j+1})
	}

	return neighbors
}

func main() {
	board := [][]rune {
		{'t', 'h', 'i', 's', 'i', 's', 'a'},
    	{'s', 'i', 'm', 'p', 'l', 'e', 'x'},
    	{'b', 'x', 'x', 'x', 'x', 'e', 'b'},
    	{'x', 'o', 'g', 'g', 'l', 'x', 'o'},
    	{'x', 'x', 'x', 'D', 'T', 'r', 'a'},
    	{'R', 'E', 'P', 'E', 'A', 'd', 'x'},
    	{'x', 'x', 'x', 'x', 'x', 'x', 'x'},
    	{'N', 'O', 'T', 'R', 'E', '-', 'P'},
    	{'x', 'x', 'D', 'E', 'T', 'A', 'E'},
	}
	words := []string{"this", "is", "not", "a", "simple", "boggle", "board", "test", "REPEATED", "NOTRE-PEATED"}

	fmt.Println(BoggleBoard(board, words))
}