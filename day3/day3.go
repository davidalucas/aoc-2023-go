package day3

type PartNumber struct {
	Number   int
	IsValid  bool
	StartIdx int
	EndIdx   int
}

type Gear struct {
	PartNumbers []int
}

func (gear *Gear) Ratio() int {
	return gear.PartNumbers[0] * gear.PartNumbers[1]
}
