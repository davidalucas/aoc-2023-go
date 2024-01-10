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
