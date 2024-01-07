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

// FindMinimumLocation finds the minimum seed location, as described in the
// AoC Day 5 Part 1 problem.
func (almanac *Almanac) FindMinimumLocation() int64 {
	var minLocation int64

	for i, seed := range almanac.Seeds {
		source := seed

		for _, mapCol := range almanac.MapCollections {
			for _, alMap := range mapCol {
				dest, valid := alMap.GetDestination(source)
				if !valid {
					continue
				}
				source = dest
				break
			}
		}

		if i == 0 {
			minLocation = source
		} else if source < minLocation {
			minLocation = source
		}
	}

	return minLocation
}

// FindMinimumLocationImproved finds the minimum seed location, as described in the
// AoC Day 5 Part 2 problem.
func (almanac *Almanac) FindMinimumLocationImproved() int64 {
	var minLocation int64 = -1

	// loop over each seed-range pair
	for i := 0; i < len(almanac.Seeds); i += 2 {
		seed := almanac.Seeds[i]
		seedRange := almanac.Seeds[i+1]
		// loop over each seed (but jump using ranges)
		for seed < almanac.Seeds[i]+almanac.Seeds[i+1] {
			src := seed
			rng := seedRange - seed

			// traverse the map collections
			for _, mapColl := range almanac.MapCollections {
				corrMap := SelectCorrespondingMap(mapColl, src)
				src, rng = corrMap.GetDestinationRange(src, rng)
			}

			if minLocation == -1 {
				minLocation = src
			} else if src < minLocation {
				minLocation = src
			}

			seed += rng
		}
	}
	return minLocation
}
