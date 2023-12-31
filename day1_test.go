package aoc2023go

import "testing"

// Test the Calibrate function with the provided sample data
func Test_Calibrate_With_Sample_Data(t *testing.T) {
	data := []string{
		"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	expected := 142
	actual := Calibrate(&data)
	if expected != actual {
		t.Fatalf(`Expected %v, received %v`, expected, actual)
	}
}
