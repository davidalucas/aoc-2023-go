package aoc2023go

import "strconv"

// Calibrate calculates the calibration sum based on the provided lines.
// It iterates through each line and finds the first and last digit in each line.
// The first and last digits are converted to integers and added to the sum.
// The final sum is returned.
func Calibrate(lines *[]string) int {
	sum := 0

	first := 0
	second := 0
	line := ""

	for i := 0; i < len(*lines); i++ {
		line = (*lines)[i]
		// loop forward
		for j := 0; j < len(line); j++ {
			val, err := strconv.Atoi(string(line[j]))
			if err != nil {
				continue
			}
			first = val
			break
		}

		// loop backward
		for j := len(line) - 1; j >= 0; j-- {
			val, err := strconv.Atoi(string(line[j]))
			if err != nil {
				continue
			}
			second = val
			break
		}

		sum += first*10 + second
		first, second = 0, 0
	}

	return sum
}

func CalibrateEnhanced(lines *[]string, numMap map[string]int) int {
	sum := 0

	first := 0
	second := 0
	line := ""

	for i := 0; i < len(*lines); i++ {
		line = (*lines)[i]
		for j := 0; j < len(line); j++ {
			val, found := PerformEnhancedForwardSearch(line, j, numMap)
			if found {
				first = val
				break
			}
		}
		for j := len(line) - 1; j >= 0; j-- {
			val, found := PerformEnhancedReverseSearch(line, j, numMap)
			if found {
				second = val
				break
			}
		}
		sum += first*10 + second
		first, second = 0, 0
	}

	return sum
}

func PerformEnhancedForwardSearch(line string, idx int, numMap map[string]int) (int, bool) {
	num := 0
	found := false

	return num, found
}

func PerformEnhancedReverseSearch(line string, idx int, numMap map[string]int) (int, bool) {
	num := 0
	found := false

	return num, found
}
