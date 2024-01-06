package day4

import (
	"container/list"
	"testing"
)

func Test_CountMatches_Calculates_Correct_Number_Of_Matches(t *testing.T) {
	data := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expected := 4
	actual := CountMatches(data)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_CountAllMatches_Returns_Correct_Number_For_Example_Data(t *testing.T) {
	expected := 13
	actual, err := SumAllPoints("example.txt")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_CountAllMatches_Returns_Correct_Number_For_Real_Data(t *testing.T) {
	expected := 23235
	actual, err := SumAllPoints("data.txt")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_CountCards_Calculates_Correct_Number_Of_Cards_And_Memory(t *testing.T) {
	data := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	memory := list.New()
	expectedCount := 1
	actual := CountCards(data, memory)

	if actual != expectedCount {
		t.Errorf("Expected %v, got %v", expectedCount, actual)
	}

	if memory.Len() != 4 {
		t.Errorf("Expected memory length to be 4, got %v", memory.Len())
	}

	if memory.Front().Value.(int) != 1 {
		t.Errorf("Expected memory's first value to be 1, got %v", memory.Front().Value)
	}

	if memory.Back().Value.(int) != 1 {
		t.Errorf("Expected memory's last value to be 1, got %v", memory.Back().Value)
	}
}

func Test_SumAllCards_Calculates_Correct_Number_For_Example_Data(t *testing.T) {
	expected := 30
	actual, err := SumAllCards("example.txt")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_SumAllCards_Calculates_Correct_Number_For_Real_Data(t *testing.T) {
	expected := 5920640
	actual, err := SumAllCards("data.txt")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
