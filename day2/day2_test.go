package day2

import (
	"reflect"
	"testing"
)

func Test_MakeReveal(t *testing.T) {
	input := "1 red, 2 green, 6 blue"
	expected := Reveal{Reds: 1, Greens: 2, Blues: 6}

	result, err := MakeReveal(input)

	if err != nil {
		t.Errorf("Method returned the following error: %v", err.Error())
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func Test_MakeCubeGame(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	expected := CubeGame{GameNumber: 1,
		Reveals: []Reveal{
			{Reds: 4, Greens: 0, Blues: 3},
			{Reds: 1, Greens: 2, Blues: 6},
			{Reds: 0, Greens: 2, Blues: 0},
		},
	}

	result, err := MakeCubeGame(input)

	if err != nil {
		t.Errorf("Method returned the following error: %v", err.Error())
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func Test_IsGamePossible_With_Sample_Data(t *testing.T) {
	// setup
	data := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}

	expected := []bool{
		true,
		true,
		false,
		false,
		true,
	}

	// test
	results := make([]bool, len(data))
	for i, d := range data {
		game, err := MakeCubeGame(d)
		if err != nil {
			t.Errorf("Failed to make CubeGame: %v", err)
			continue
		}
		results[i] = game.IsGamePossible(12, 13, 14)
	}

	// assert
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %+v, got %+v", expected, results)
	}
}
