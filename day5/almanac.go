package day5

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Almanac struct {
	Seeds          []int64
	MapCollections [][]AlmanacMap
}

// MakeAlmanac makes an Almanac object from the data contained
// within a text file located at the specified path location.
func MakeAlmanac(path string) (*Almanac, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seeds, err := ParseSeeds(scanner)
	if err != nil {
		return nil, err
	}

	mapCollections, err := ParseMapCollections(scanner)
	if err != nil {
		return nil, err
	}

	return &Almanac{Seeds: seeds, MapCollections: mapCollections}, nil
}

// ParseSeeds parses the array of seed values from the first line of the data file.
// The scanner is assumed to be at the beginning of the file (not on the first line yet)
// when this function is called, and it is assumed that the seed values are found on
// the first line of the file.
func ParseSeeds(scanner *bufio.Scanner) ([]int64, error) {
	scanner.Scan() // move to first line
	seedData := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")
	seeds := make([]int64, len(seedData))
	for i := 0; i < len(seedData); i++ {
		val, err := strconv.ParseInt(seedData[i], 10, 64)
		if err != nil {
			return seeds, fmt.Errorf("failed to parse seed value '%s': %w", seedData[i], err)
		}
		seeds[i] = val
	}
	scanner.Scan() // progress the scanner once more to skip blank line
	return seeds, nil
}

// ParseMapCollections walks the file with the provided scanner, and parses out
// all of the AlmanacMaps, returning the slice of map collections. Just for clarity,
// the "soil-to-fertilizer" map (and all other "maps") is represented as a []AlmanacMap
// slice.
func ParseMapCollections(scanner *bufio.Scanner) ([][]AlmanacMap, error) {
	dataQueue := list.New()
	var mapCollections [][]AlmanacMap
	for scanner.Scan() {
		if scanner.Text() != "" {
			dataQueue.PushBack(scanner.Text())
			continue
		}
		mapCollection, err := MakeMapCollection(dataQueue)
		if err != nil {
			return nil, err
		}
		mapCollections = append(mapCollections, mapCollection)
	}
	// empty out the queue at the end of the file
	mapCollection, err := MakeMapCollection(dataQueue)
	if err != nil {
		return nil, err
	}
	mapCollections = append(mapCollections, mapCollection)
	return mapCollections, nil
}
