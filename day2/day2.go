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

func MakeReveal(strData string) (Reveal, error) {
	reveal := Reveal{}
	var err error

	splitData := strings.Split(strData, ", ")
	for _, s := range splitData {
		couple := strings.Split(s, " ")
		switch couple[1] {
		case "red":
			reveal.Reds, err = strconv.Atoi(couple[0])
		case "green":
			reveal.Greens, err = strconv.Atoi(couple[0])
		case "blue":
			reveal.Blues, err = strconv.Atoi(couple[0])
		}
		if err != nil {
			return reveal, err
		}
	}

	return reveal, err
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
		cg.Reveals = append(cg.Reveals, reveal) // not the most performant, but it's fine since it's only going to re-assign a max of 3 times
	}

	return cg, err
}
