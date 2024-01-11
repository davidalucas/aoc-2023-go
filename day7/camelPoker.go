package day7

import (
	"math"
	"strconv"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
	Score int
}

// MakeHand parses a raw data string, constructing the resulting
// Hand object.
func MakeHand(rawData string) (*Hand, error) {
	splitData := strings.Split(rawData, " ")
	cards := splitData[0]
	bid, err := strconv.ParseInt(splitData[1], 10, 16)
	if err != nil {
		return nil, err
	}
	score := calcCardsScore(cards)

	return &Hand{Cards: cards, Bid: int(bid), Score: score}, nil
}

// CalcCardsScore calculates the base score for a set of cards.
func calcCardsScore(cards string) int {
	cardMap := make(map[byte]float64)
	for i := 0; i < len(cards); i++ {
		_, ok := cardMap[cards[i]]
		if !ok {
			cardMap[cards[i]] = 1
			continue
		}
		cardMap[cards[i]]++
	}
	score := 0
	for _, v := range cardMap {
		score += int(math.Pow(v, 2))
	}
	return score
}

// CompareHands compares two hands and returns <0 if 'a' is less than 'b',
// >0 if 'a' is greater than 'b', and 0 if 'a' and 'b' are equal.
func CompareHands(a *Hand, b *Hand) int {
	diff := a.Score - b.Score

	if diff != 0 {
		return diff
	}

	for i := 0; i < len(a.Cards); i++ {
		if CardValues[a.Cards[i]] == CardValues[b.Cards[i]] {
			continue
		}
		return CardValues[a.Cards[i]] - CardValues[b.Cards[i]]
	}

	return 0
}
