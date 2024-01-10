package day6

import "math"

// CalcWaysToWin calculates the number of ways to win a boat
// race, as described in AoC Day 6, Part 1
func CalcWaysToWin(raceData RaceData) int {
	T := raceData.TotalTime
	d := raceData.RecordDistance
	t_max := (T + math.Sqrt(T*T-4*d)) / 2
	t_max_floor := math.Floor(t_max)
	t_min := (T - math.Sqrt(T*T-4*d)) / 2
	t_min_ceil := math.Ceil(t_min)

	if t_max == t_max_floor {
		t_max_floor--
	}

	if t_min == t_min_ceil {
		t_min_ceil++
	}

	return int(t_max_floor) - int(t_min_ceil) + 1
}

// CalcErrorMargin solves the Day 6 Part 1 problem
func CalcErrorMargin(races []RaceData) int {
	errorMargin := 1

	for _, rd := range races {
		errorMargin *= CalcWaysToWin(rd)
	}

	return errorMargin
}
