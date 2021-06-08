package main

import "fmt"

const HOME_TEAM_WON = 1

// O(n) time | O(k) space
// n: no. of competitions | k: no. of teams
func TournamentWinner(competitions [][]string, results []int) string {
	currentBestTeam := "" 
	// Hash table to store total pts?
	scores := map[string]int{currentBestTeam: 0}

	for idx, competition := range competitions {
		result := results[idx]
		homeTeam, awayTeam := competition[0], competition[1]

		winningTeam := awayTeam
		if result == HOME_TEAM_WON {
			winningTeam = homeTeam
		}

		updateScores(winningTeam, 3, scores)

		if scores[winningTeam] > scores[currentBestTeam] {
			currentBestTeam = winningTeam
		}
	}
	return currentBestTeam
}

func updateScores(team string, points int, scores map[string]int) {
	scores[team] += points
}

func main() {
	comp := [][]string{
		{"HTML", "C#"},
    	{"C#", "Python"},
    	{"Python", "HTML"},
	}
	res := []int{0, 0, 1}
	fmt.Println("Tournament Winner:", TournamentWinner(comp, res))
}