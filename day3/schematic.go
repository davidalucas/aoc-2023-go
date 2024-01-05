package day3

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Schematic struct {
	Data []string
}

// MakeSchematic builds a Schematic object using the data contained
// in the file at the specified path.
func MakeSchematic(path string) (*Schematic, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := strings.Split(string(content), "\n")
	return &Schematic{Data: data}, nil
}

func (schematic *Schematic) FindAllParts() ([][]PartNumber, error) {
	partNumbers := make([][]PartNumber, len(schematic.Data))

	// find all part numbers
	for i, line := range schematic.Data {
		chars := strings.Split(line, "")
		for j := 0; j < len(chars); j++ {
			_, err := strconv.Atoi(chars[j])
			if err != nil {
				continue
			}
			start := j

			for err == nil {
				j++
				if j >= len(chars) {
					break
				}
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

	regex, _ := regexp.Compile(`[^0-9\.]`)

	// validate all part numbers against schematic
	for i, line := range schematic.Data {
		chars := strings.Split(line, "")
		for j, char := range chars {
			match := regex.MatchString(char)
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

// MakeAsteriskMap returns a map representing the surrounding area of an asterisk in a grid.
// The asterisk is located at the specified line and linePosition.
// The map contains boolean values indicating whether each position needs to be assessed
// (i.e. 'false' indicates that this location can be skipped).
func (schematic *Schematic) MakeAsteriskMap(line int, linePosition int) map[int]map[int]bool {
	asteriskMap := map[int]map[int]bool{
		line - 1: {
			linePosition - 1: true,
			linePosition:     true,
			linePosition + 1: true,
		},
		line: {
			linePosition - 1: true,
			linePosition:     false, // this is the location of the asterisk
			linePosition + 1: true,
		},
		line + 1: {
			linePosition - 1: true,
			linePosition:     true,
			linePosition + 1: true,
		},
	}
	// if on top line
	if line == 0 {
		for k := range asteriskMap[line-1] {
			asteriskMap[line-1][k] = false
		}
	}
	// if on right edge
	if linePosition == len(schematic.Data[line])-1 {
		for k := range asteriskMap {
			asteriskMap[k][linePosition+1] = false
		}
	}
	// if on bottom line
	if line == len(schematic.Data)-1 {
		for k := range asteriskMap[line+1] {
			asteriskMap[line+1][k] = false
		}
	}
	// if on left edge
	if linePosition == 0 {
		for k := range asteriskMap {
			asteriskMap[k][linePosition-1] = false
		}
	}
	return asteriskMap
}
