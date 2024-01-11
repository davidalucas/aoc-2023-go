package day7

import (
	"reflect"
	"testing"
)

func Test_MakeHand_Returns_Expected_Hand(t *testing.T) {
	expected := &Hand{Cards: "32T3K", Bid: 765, Score: 7}
	actual, err := MakeHand("32T3K 765")

	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_ParsePokerFile_Returns_Expected_Data_For_Example_File(t *testing.T) {
	expected := []Hand{
		{Cards: "32T3K", Bid: 765, Score: 7},
		{Cards: "T55J5", Bid: 684, Score: 11},
		{Cards: "KK677", Bid: 28, Score: 9},
		{Cards: "KTJJT", Bid: 220, Score: 9},
		{Cards: "QQQJA", Bid: 483, Score: 11},
	}
	actual, err := ParsePokerFile("example.txt")

	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_CompareHands_Returns_Negative_When_Expected(t *testing.T) {
	hand1, err := MakeHand("T55J5 684")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	hand2, err := MakeHand("QQQJA 483")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	compareResult := CompareHands(hand1, hand2)

	if compareResult >= 0 {
		t.Errorf("Expected comparing hand %v against %v would be less than zero, but received %v", hand1.Cards, hand2.Cards, compareResult)
	}
}

func Test_CompareHands_Returns_Positive_When_Expected(t *testing.T) {
	hand1, err := MakeHand("T55J5 684")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	hand2, err := MakeHand("QQQJA 483")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	compareResult := CompareHands(hand2, hand1)

	if compareResult <= 0 {
		t.Errorf("Expected comparing hand %v against %v would be greater than zero, but received %v", hand2.Cards, hand1.Cards, compareResult)
	}
}

func Test_CompareHands_Returns_Zero_When_Expected(t *testing.T) {
	hand, err := MakeHand("T55J5 684")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	compareResult := CompareHands(hand, hand)

	if compareResult != 0 {
		t.Errorf("Expected comparing hand %v against itself would return zero, but received %v", hand.Cards, compareResult)
	}
}
