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
