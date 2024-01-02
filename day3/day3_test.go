package day3

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_MakeSchematic_With_Sample_Data(t *testing.T) {
	data := []string{
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

	actual, err := ParseSchematic(data)

	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	for i, a := range actual {
		if !reflect.DeepEqual(a, expected[i]) {
			t.Errorf("Expected %+v, got %+v", expected[i], a)
		}
	}
}

func Test_MakeSchematic_With_Real_Data(t *testing.T) {
	content, err := os.ReadFile("day3.txt")
	if err != nil {
		t.Fatal(err)
	}

	data := strings.Split(string(content), "\n")
	expected := 544664

	schematic, err := ParseSchematic(data)
	if err != nil {
		t.Errorf("Test failed with the following error: %v", err)
	}

	actual := 0
	for _, partNums := range schematic {
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
