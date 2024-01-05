package day3

import (
	"reflect"
	"slices"
	"testing"
)

func Test_MakeSchematic_With_Example_Data(t *testing.T) {
	expected := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	schematic, err := MakeSchematic("example.txt")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	actual := schematic.Data
	if !slices.Equal(actual, expected) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func Test_FindAllParts_With_Sample_Data(t *testing.T) {
	expected := [][]PartNumber{
		{
			PartNumber{Number: 467, IsValid: true, StartIdx: 0, EndIdx: 2},
			PartNumber{Number: 114, IsValid: false, StartIdx: 5, EndIdx: 7},
		},
		nil,
		{
			PartNumber{Number: 35, IsValid: true, StartIdx: 2, EndIdx: 3},
			PartNumber{Number: 633, IsValid: true, StartIdx: 6, EndIdx: 8},
		},
		nil,
		{
			PartNumber{Number: 617, IsValid: true, StartIdx: 0, EndIdx: 2},
		},
		{
			PartNumber{Number: 58, IsValid: false, StartIdx: 7, EndIdx: 8},
		},
		{
			PartNumber{Number: 592, IsValid: true, StartIdx: 2, EndIdx: 4},
		},
		{
			PartNumber{Number: 755, IsValid: true, StartIdx: 6, EndIdx: 8},
		},
		nil,
		{
			PartNumber{Number: 664, IsValid: true, StartIdx: 1, EndIdx: 3},
			PartNumber{Number: 598, IsValid: true, StartIdx: 5, EndIdx: 7},
		},
	}
	schematic, err := MakeSchematic("example.txt")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	actual, err := schematic.FindAllParts()
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	for i, a := range actual {
		if !reflect.DeepEqual(a, expected[i]) {
			t.Errorf("Expected %+v, got %+v", expected[i], a)
		}
	}
}

func Test_FindAllParts_With_Real_Data(t *testing.T) {
	expected := 544664

	schematic, err := MakeSchematic("day3.txt")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	partNumsArr, err := schematic.FindAllParts()
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	actual := 0
	for _, partNums := range partNumsArr {
		for _, pn := range partNums {
			if pn.IsValid {
				actual += pn.Number
			}
		}
	}

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}

func Test_MakeAsteriskMap_With_Example_Data(t *testing.T) {
	schematic, err := MakeSchematic("day3.txt")
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}
	expected := map[int]map[int]bool{
		0: {2: true, 3: true, 4: true},
		1: {2: true, 3: false, 4: true},
		2: {2: true, 3: true, 4: true},
	}
	actual := schematic.MakeAsteriskMap(1, 3)
	if !mapsEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func mapsEqual(a, b map[int]map[int]bool) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if bv, ok := b[k]; !ok || !innerMapsEqual(v, bv) {
			return false
		}
	}

	return true
}

func innerMapsEqual(a, b map[int]bool) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}

	return true
}
