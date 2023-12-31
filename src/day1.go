package aoc2023go

import "strconv"

// Calibrate calculates the calibration sum based on the provided lines.
// It iterates through each line and finds the first and last digit in each line.
// The first and last digits are converted to integers and added to the sum.
// The final sum is returned.
func Calibrate(lines *[]string) int {
	sum := 0;

	first := 0;
	second := 0;
	line := "";

	for i := 0; i < len(*lines); i++ {
		line = (*lines)[i]
		// loop forward
		for j := 0; j < len(line); j++ {
			val, err := strconv.Atoi(string(line[j]))
			if err != nil {
				continue;
			}
			first = val
			break;
		}

		// loop backward
		for j := len(line) - 1; j >= 0; j-- {
			val, err := strconv.Atoi(string(line[j]))
			if err != nil {
				continue;
			}
			second = val
			break;
		}

		sum += first * 10 + second
	}
	
	return sum;
}