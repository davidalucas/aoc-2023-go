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
