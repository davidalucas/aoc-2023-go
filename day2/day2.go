package day2

type CubeGame struct {
	GameNumber int
	Reveals    []Reveal
}

type Reveal struct {
	Reds   int
	Greens int
	Blues  int
}

func MakeReveal(strData string) Reveal {
	var reveal = Reveal{}

	// TODO: implement string parsing logic

	return reveal
}

func MakeCubeGame(strData string) CubeGame {
	var cg = CubeGame{}

	// TODO: implement string parsing logic

	return cg
}
