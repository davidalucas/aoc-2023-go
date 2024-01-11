package day7

import (
	"bufio"
	"math"
	"os"
	"slices"
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

// MakeHandWithJokers parses a raw data string, constructing the resulting
// Hand object, with Jokers taken into account.
func MakeHandWithJokers(rawData string) (*Hand, error) {
	splitData := strings.Split(rawData, " ")
	cards := splitData[0]
	bid, err := strconv.ParseInt(splitData[1], 10, 16)
	if err != nil {
		return nil, err
	}
	score := calcCardsScoreWithJokers(cards)

	return &Hand{Cards: cards, Bid: int(bid), Score: score}, nil
}

// CalcCardsScore calculates the base score for a set of cards.
func calcCardsScoreWithJokers(cards string) int {
	// find matches
	cardMap := make(map[byte]float64)
	for i := 0; i < len(cards); i++ {
		_, ok := cardMap[cards[i]]
		if !ok {
			cardMap[cards[i]] = 1
			continue
		}
		cardMap[cards[i]]++
	}

	// correct for Jokers
	if cardMap['J'] != 0 {
		numJokers := cardMap['J']
		cardMap['J'] = 0
		var maxPairCard byte = 'J'
		var maxPairs float64 = 0
		for k, v := range cardMap {
			if v > maxPairs {
				maxPairCard = k
				maxPairs = v
			}
		}
		cardMap[maxPairCard] += numJokers
	}

	// score
	score := 0
	for _, v := range cardMap {
		score += int(math.Pow(v, 2))
	}
	return score
}

// ParsePokerFile parses a file at the specified path location, returning
// a slice of Hand objects. Returns errors if the file path was incorrect,
// or if the file's contents was in an unexpected format.
func ParsePokerFile(path string) ([]Hand, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hands := make([]Hand, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hand, err := MakeHand(scanner.Text())
		if err != nil {
			return nil, err
		}
		hands = append(hands, *hand)
	}
	return hands, nil
}

// ParsePokerFileWithJokers is the same as ParsePokerFile, accept 'J' cards
// are treated as Jokers instead of Jacks.
func ParsePokerFileWithJokers(path string) ([]Hand, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hands := make([]Hand, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hand, err := MakeHandWithJokers(scanner.Text())
		if err != nil {
			return nil, err
		}
		hands = append(hands, *hand)
	}
	return hands, nil
}

// CompareHands compares two hands and returns <0 if 'a' is less than 'b',
// >0 if 'a' is greater than 'b', and 0 if 'a' and 'b' are equal.
func CompareHands(a Hand, b Hand) int {
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

// CompareHandsWithJokers is the same as CompareHands, except 'J' cards are treated
// as Jokers instead of Jacks.
func CompareHandsWithJokers(a Hand, b Hand) int {
	diff := a.Score - b.Score

	if diff != 0 {
		return diff
	}

	for i := 0; i < len(a.Cards); i++ {
		if CardValuesWithJokers[a.Cards[i]] == CardValuesWithJokers[b.Cards[i]] {
			continue
		}
		return CardValuesWithJokers[a.Cards[i]] - CardValuesWithJokers[b.Cards[i]]
	}

	return 0
}

// CalcTotalWinnings calculates the total winnings, as described in the
// Day 7 Part 1 problem description.
func CalcTotalWinnings(hands []Hand) int {
	slices.SortFunc(hands, CompareHands)
	winnings := 0
	for i, h := range hands {
		winnings += h.Bid * (i + 1)
	}
	return winnings
}

// CalcTotalWinningsWithJokers calculates the total winnings, as described in the
// Day 7 Part 2 problem description.
func CalcTotalWinningsWithJokers(hands []Hand) int {
	slices.SortFunc(hands, CompareHandsWithJokers)
	winnings := 0
	for i, h := range hands {
		winnings += h.Bid * (i + 1)
	}
	return winnings
}
