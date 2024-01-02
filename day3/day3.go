package day3

import (
	"regexp"
	"strconv"
	"strings"
)

type PartNumber struct {
	Number   int
	IsValid  bool
	StartIdx int
	EndIdx   int
}

// ParseSchematic parses the provided string data, producing a
// 2D array of PartNumbers.
func ParseSchematic(data []string) ([][]PartNumber, error) {
	partNumbers := make([][]PartNumber, len(data))

	// find all part numbers
	for i, line := range data {
		chars := strings.Split(line, "")
		for j := 0; j < len(chars); j++ {
			_, err := strconv.Atoi(chars[j])
			if err != nil {
				continue
			}
			start := j

			for err == nil {
				j++
				_, err = strconv.Atoi(chars[j])
			}
			end := j - 1
			number, _ := strconv.Atoi(line[start : end+1])
			partNumbers[i] = append(partNumbers[i], PartNumber{
				Number:   number,
				StartIdx: start,
				EndIdx:   end,
			})
		}
	}

	// validate all part numbers against schematic
	for i, line := range data {
		chars := strings.Split(line, "")
		for j, char := range chars {
			match, _ := regexp.MatchString(`[^0-9\.]`, char)
			if !match {
				continue
			}
			// check PNs above
			for k := 0; k < len(partNumbers[i-1]); k++ {
				if partNumbers[i-1][k].IsValid {
					continue
				}
				if partNumbers[i-1][k].StartIdx <= j-1 && j-1 <= partNumbers[i-1][k].EndIdx {
					partNumbers[i-1][k].IsValid = true
				} else if partNumbers[i-1][k].StartIdx <= j && j <= partNumbers[i-1][k].EndIdx {
					partNumbers[i-1][k].IsValid = true
				} else if partNumbers[i-1][k].StartIdx <= j+1 && j+1 <= partNumbers[i-1][k].EndIdx {
					partNumbers[i-1][k].IsValid = true
				}
			}
			// check left and right
			for k := 0; k < len(partNumbers[i]); k++ {
				if partNumbers[i][k].IsValid {
					continue
				}
				if partNumbers[i][k].EndIdx == j-1 || partNumbers[i][k].StartIdx == j+1 {
					partNumbers[i][k].IsValid = true
				}
			}
			// check PNs below
			for k := 0; k < len(partNumbers[i+1]); k++ {
				if partNumbers[i+1][k].IsValid {
					continue
				}
				if partNumbers[i+1][k].StartIdx <= j-1 && j-1 <= partNumbers[i+1][k].EndIdx {
					partNumbers[i+1][k].IsValid = true
				} else if partNumbers[i+1][k].StartIdx <= j && j <= partNumbers[i+1][k].EndIdx {
					partNumbers[i+1][k].IsValid = true
				} else if partNumbers[i+1][k].StartIdx <= j+1 && j+1 <= partNumbers[i+1][k].EndIdx {
					partNumbers[i+1][k].IsValid = true
				}
			}
		}
	}

	return partNumbers, nil
}
