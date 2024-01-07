package day5

import (
	"reflect"
	"testing"
)

func Test_MakeAlmanac_Constructs_Expected_Almanac(t *testing.T) {
	expected := &Almanac{
		Seeds: []int64{79, 14, 55, 13},
		MapCollections: [][]AlmanacMap{
			{
				{Source: 50, Destination: 52, Range: 48},
				{Source: 98, Destination: 50, Range: 2},
			},
			{
				{Source: 0, Destination: 39, Range: 15},
				{Source: 15, Destination: 0, Range: 37},
				{Source: 52, Destination: 37, Range: 2},
			},
			{
				{Source: 0, Destination: 42, Range: 7},
				{Source: 7, Destination: 57, Range: 4},
				{Source: 11, Destination: 0, Range: 42},
				{Source: 53, Destination: 49, Range: 8},
			},
			{
				{Source: 18, Destination: 88, Range: 7},
				{Source: 25, Destination: 18, Range: 70},
			},
			{
				{Source: 45, Destination: 81, Range: 19},
				{Source: 64, Destination: 68, Range: 13},
				{Source: 77, Destination: 45, Range: 23},
			},
			{
				{Source: 0, Destination: 1, Range: 69},
				{Source: 69, Destination: 0, Range: 1},
			},
			{
				{Source: 56, Destination: 60, Range: 37},
				{Source: 93, Destination: 56, Range: 4},
			},
		},
	}
	actual, err := MakeAlmanac("example.txt")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Test_FindMinimumLocation_Finds_Correct_Location_For_Example_Data(t *testing.T) {
	var expected int64 = 35
	almanac, err := MakeAlmanac("example.txt")
	if err != nil {
		t.Errorf("test failed with the following error: %v", err)
	}
	actual := almanac.FindMinimumLocation()
	if actual != expected {
		t.Errorf("expected %v, received %v", expected, actual)
	}
}

func Test_FindMinimumLocation_Finds_Correct_Location_For_Real_Data(t *testing.T) {
	var expected int64 = 227653707
	almanac, err := MakeAlmanac("data.txt")
	if err != nil {
		t.Errorf("test failed with the following error: %v", err)
	}
	actual := almanac.FindMinimumLocation()
	if actual != expected {
		t.Errorf("expected %v, received %v", expected, actual)
	}
}
