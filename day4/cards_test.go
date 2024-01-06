package day4

import "testing"

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
