package day5

import (
	"reflect"
	"testing"
)

func Test_MakeAlmanacMap_Constructs_Valid_AlmanacMap(t *testing.T) {
	expected := &AlmanacMap{
		Source:      56,
		Destination: 60,
		Range:       37,
	}
	actual, err := MakeAlmanacMap("60 56 37")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
