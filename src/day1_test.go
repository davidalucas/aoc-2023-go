package aoc2023go

import (
	"os"
	"strings"
	"testing"
)

// Test_Calibrate_With_Sample_Data tests the Calibrate function with sample data.
// It verifies that the expected result matches the actual result.
func Test_Calibrate_With_Sample_Data(t *testing.T) {
	data := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	expected := 142
	actual := Calibrate(&data)
	if expected != actual {
		t.Fatalf(`Expected %v, received %v`, expected, actual)
	}
}

// Test_Calibrate_With_Real_Data tests the Calibrate function with real data.
// It reads the content from "../data/day1.txt", splits it by newline,
// and compares the expected result with the actual result.
// If they don't match, it fails the test.
func Test_Calibrate_With_Real_Data(t *testing.T) {
	content, err := os.ReadFile("../data/day1.txt")
	if err != nil {
		t.Fatal(err)
	}

	data := strings.Split(string(content), "\n")
	expected := 55477
	actual := Calibrate(&data)
	if expected != actual {
		t.Fatalf(`Expected %v, received %v`, expected, actual)
	}
}
