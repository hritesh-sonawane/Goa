package main

import "fmt"

// O(n) time | O(1) space
func ValidStartingCity(distances []int, fuel []int, mpg int) int {
	numberOfCities := len(distances)
	milesRemaining := 0

	indexOfMaybeStartingCity := 0
	milesRemainingAtMaybeStartingCity := 0

	for cityIdx := 1; cityIdx < numberOfCities; cityIdx++ {
		distanceFromPrevCity := distances[cityIdx-1]
		fuelFromPrevCity := fuel[cityIdx-1]
		milesRemaining += fuelFromPrevCity*mpg - distanceFromPrevCity

		if milesRemaining < milesRemainingAtMaybeStartingCity {
			milesRemainingAtMaybeStartingCity = milesRemaining
			indexOfMaybeStartingCity = cityIdx
		}
	}
	return indexOfMaybeStartingCity
}

func main() {
	distances := []int{5, 25, 15, 10, 15}
	fuel := []int{1, 2, 1, 0, 3}
	mpg := 10
	fmt.Println(ValidStartingCity(distances, fuel, mpg))
}