package day6

type RaceData struct {
	TotalTime      float64
	RecordDistance float64
}

var exampleData []RaceData = []RaceData{
	{TotalTime: 7, RecordDistance: 9},
	{TotalTime: 15, RecordDistance: 40},
	{TotalTime: 30, RecordDistance: 200},
}

var exampleData2 []RaceData = []RaceData{
	{TotalTime: 71530, RecordDistance: 940200},
}

var realData []RaceData = []RaceData{
	{TotalTime: 46, RecordDistance: 358},
	{TotalTime: 68, RecordDistance: 1054},
	{TotalTime: 98, RecordDistance: 1807},
	{TotalTime: 66, RecordDistance: 1080},
}

var realData2 []RaceData = []RaceData{
	{TotalTime: 46689866, RecordDistance: 358105418071080},
}
