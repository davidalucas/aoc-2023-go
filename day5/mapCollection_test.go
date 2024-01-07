package day5

import (
	"container/list"
	"reflect"
	"testing"
)

func Test_MakeMapCollection_Constructs_Expected_Slice(t *testing.T) {
	dataQueue := list.New()
	dataQueue.PushBack("soil-to-fertilizer map:")
	dataQueue.PushBack("0 15 37")
	dataQueue.PushBack("37 52 2")
	dataQueue.PushBack("39 0 15")

	expected := []AlmanacMap{
		{Source: 0, Destination: 39, Range: 15},
		{Source: 15, Destination: 0, Range: 37},
		{Source: 52, Destination: 37, Range: 2},
	}

	actual, err := MakeMapCollection(dataQueue)
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
