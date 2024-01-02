package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type CubeGame struct {
	GameNumber int
	Reveals    []Reveal
}

type Reveal struct {
	Reds   int
	Greens int
	Blues  int
}

func MakeReveal(strData string) (*Reveal, error) {
	splitData := strings.Split(strData, ", ")
	revMap := make(map[string]int, len(splitData))
	for _, s := range splitData {
		couple := strings.Split(s, " ")
		val, err := strconv.Atoi(couple[0])
		if err != nil {
			return nil, err
		}
		revMap[couple[1]] = val
	}

	return &Reveal{
		Reds:   revMap["red"],
		Greens: revMap["green"],
		Blues:  revMap["blue"],
	}, nil
}

func MakeCubeGame(strData string) (*CubeGame, error) {
	splitData := strings.Split(strData, ": ")
	gameNumber, err := strconv.Atoi(strings.Split(splitData[0], " ")[1])
	if err != nil {
		return nil, fmt.Errorf(`failed to parse out game number for cube game: %v`, splitData[0])
	}

	rawRevealData := strings.Split(splitData[1], "; ")
	reveals := make([]Reveal, len(rawRevealData))
	for i, r := range rawRevealData {
		reveal, err := MakeReveal(r)
		if err != nil {
			return nil, err
		}
		reveals[i] = *reveal
	}

	return &CubeGame{GameNumber: gameNumber, Reveals: reveals}, nil
}

func (game *CubeGame) IsPossible(maxRed int, maxGreen int, maxBlue int) bool {
	for _, r := range game.Reveals {
		if r.Reds > maxRed || r.Greens > maxGreen || r.Blues > maxBlue {
			return false
		}
	}
	return true
}

func SumPossibleGames(gameData []string, maxRed int, maxGreen int, maxBlue int) (int, error) {
	game := &CubeGame{}
	var err error
	sum := 0

	for _, data := range gameData {
		game, err = MakeCubeGame(data)
		if err != nil {
			return 0, err
		}
		if game.IsPossible(maxRed, maxGreen, maxBlue) {
			sum += game.GameNumber
		}
	}

	return sum, nil
}

// GetPower calculates the power factor described in the Day 2 Part 2 problem
// for this game.
func (game *CubeGame) GetPower() int {
	redsRequired := 0
	greensRequired := 0
	bluesRequired := 0

	for _, r := range game.Reveals {
		if redsRequired < r.Reds {
			redsRequired = r.Reds
		}
		if greensRequired < r.Greens {
			greensRequired = r.Greens
		}
		if bluesRequired < r.Blues {
			bluesRequired = r.Blues
		}
	}

	return redsRequired * greensRequired * bluesRequired
}

// GetPower calculates the sum of all of the powers for each of the games
// in the provided data set. Solves Day 2 Part 2.
func GetPower(data []string) (int, error) {
	return 0, nil
}
