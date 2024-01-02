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
