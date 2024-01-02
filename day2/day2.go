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

	reveal := Reveal{
		Reds:   revMap["red"],
		Greens: revMap["green"],
		Blues:  revMap["blue"],
	}

	return &reveal, nil
}

func MakeCubeGame(strData string) (CubeGame, error) {
	cg := CubeGame{}
	var err error

	splitData := strings.Split(strData, ": ")
	cg.GameNumber, err = strconv.Atoi(strings.Split(splitData[0], " ")[1])
	if err != nil {
		return cg, fmt.Errorf(`failed to parse out game number for cube game: %v`, splitData[0])
	}

	rawRevealData := strings.Split(splitData[1], "; ")
	for _, r := range rawRevealData {
		reveal, err := MakeReveal(r)
		if err != nil {
			return cg, err
		}
		cg.Reveals = append(cg.Reveals, *reveal) // not the most performant, but it's fine since it's only going to re-assign a max of 3 times
	}

	return cg, err
}

func (game *CubeGame) IsGamePossible(maxRed int, maxGreen int, maxBlue int) bool {
	for _, r := range game.Reveals {
		if r.Reds > maxRed || r.Greens > maxGreen || r.Blues > maxBlue {
			return false
		}
	}
	return true
}
