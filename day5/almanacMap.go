package day5

import (
	"strconv"
	"strings"
)

type AlmanacMap struct {
	Source      int64
	Destination int64
	Range       int64
}

// MakeAlmanacMap constructs a new AlmanacMap object using the provided string
// data.
func MakeAlmanacMap(data string) (*AlmanacMap, error) {
	splitStr := strings.Split(data, " ")
	src, err := strconv.ParseInt(splitStr[1], 10, 64)
	if err != nil {
		return nil, err
	}
	dest, err := strconv.ParseInt(splitStr[0], 10, 64)
	if err != nil {
		return nil, err
	}
	rg, err := strconv.ParseInt(splitStr[2], 10, 64)
	if err != nil {
		return nil, err
	}
	return &AlmanacMap{
		Source:      src,
		Destination: dest,
		Range:       rg,
	}, nil
}

// GetDestination retrieves the destination value that corresponds to the provided
// source value. If source is invalid, it returns the provided source value, along with
// a bool value of 'false'.
func (almanacMap *AlmanacMap) GetDestination(src int64) (int64, bool) {
	diff := src - almanacMap.Source
	if diff < 0 || diff >= almanacMap.Range {
		return src, false
	}
	return almanacMap.Destination + (src - almanacMap.Source), true
}
