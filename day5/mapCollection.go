package day5

import (
	"container/list"
	"sort"
)

// MakeMapCollection constructs a slice of AlmanacMap objects from a provided
// queue. The AlmanacMaps are also sorted in ascending order by their Source values.
func MakeMapCollection(dataQueue *list.List) ([]AlmanacMap, error) {
	dataQueue.Remove(dataQueue.Front()) // dequeue the first item (the title)
	mapCollection := make([]AlmanacMap, dataQueue.Len())
	for i := range mapCollection {
		am, err := MakeAlmanacMap(dataQueue.Remove(dataQueue.Front()).(string))
		if err != nil {
			return mapCollection, err
		}
		mapCollection[i] = *am
	}
	sort.Slice(mapCollection,
		func(i, j int) bool {
			return mapCollection[i].Source < mapCollection[j].Source
		},
	)
	return mapCollection, nil
}

// SelectCorrespondingMap finds the map which corresponds to the provided src value.
// If the src value comes before any of the maps, the first map will be returned. If
// the src value falls between two maps, the map with the larger Source value will be
// returned. If the src value comes after all of the maps, then the last map will be
// returned.
//
// Note that the maps must be sorted in ascending order by their Source values in order
// for this method to work correctly.
func SelectCorrespondingMap(mapCol []AlmanacMap, src int64) AlmanacMap {
	for _, alMap := range mapCol {
		if src <= alMap.Source+alMap.Range-1 {
			return alMap
		}
	}
	return mapCol[len(mapCol)-1]
}
