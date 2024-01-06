package day4

import (
	"bufio"
	"math"
	"os"
	"strings"
)

// CountMatches counts the total number of matches for a given "Card".
func CountMatches(cardData string) int {
	winnersAndReveals := strings.Split(strings.Split(cardData, ": ")[1], " | ")

	revealed := make(map[string]bool)

	for _, v := range strings.Split(winnersAndReveals[1], " ") {
		if v != "" {
			revealed[v] = true
		}
	}

	matches := 0

	for _, v := range strings.Split(winnersAndReveals[0], " ") {
		if v == "" {
			continue
		}
		if revealed[v] {
			matches++
		}
	}
	return matches
}

// SumAllPoints returns the total score using all of the matches found for every "Card"
// in the text file at the provided path location.
func SumAllPoints(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := CountMatches(scanner.Text())
		score += int(math.Pow(2, float64(matches-1)))
	}
	return score, nil
}
