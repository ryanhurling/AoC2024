package camel

import (
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards [5]Card
	bid   int
}

type Game []Hand

type HandType int
type CardStrength int
type Card rune
type Rank int

const (
	CARD_TWO   Card = '2'
	CARD_THREE Card = '3'
	CARD_FOUR  Card = '4'
	CARD_FIVE  Card = '5'
	CARD_SIX   Card = '6'
	CARD_SEVEN Card = '7'
	CARD_EIGHT Card = '8'
	CARD_NINE  Card = '9'
	CARD_T     Card = 'T'
	CARD_J     Card = 'J'
	CARD_Q     Card = 'Q'
	CARD_K     Card = 'K'
	CARD_A     Card = 'A'
)

func (c Card) Strength() CardStrength {
	switch c {
	case CARD_TWO:
		return STRENGTH_TWO
	case CARD_THREE:
		return STRENGTH_THREE
	case CARD_FOUR:
		return STRENGTH_FOUR
	case CARD_FIVE:
		return STRENGTH_FIVE
	case CARD_SIX:
		return STRENGTH_SIX
	case CARD_SEVEN:
		return STRENGTH_SEVEN
	case CARD_EIGHT:
		return STRENGTH_EIGHT
	case CARD_NINE:
		return STRENGTH_NINE
	case CARD_T:
		return STRENGTH_T
	case CARD_J:
		return STRENGTH_J
	case CARD_Q:
		return STRENGTH_Q
	case CARD_K:
		return STRENGTH_K
	case CARD_A:
		return STRENGTH_A
	default:
		return -1
	}
}

const (
	HIGH_CARD HandType = 1 + iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

const (
	STRENGTH_J CardStrength = 1 + iota
	STRENGTH_TWO
	STRENGTH_THREE
	STRENGTH_FOUR
	STRENGTH_FIVE
	STRENGTH_SIX
	STRENGTH_SEVEN
	STRENGTH_EIGHT
	STRENGTH_NINE
	STRENGTH_T
	STRENGTH_Q
	STRENGTH_K
	STRENGTH_A
)

func CreateHand(input string) Hand {
	parts := strings.Split(input, " ")

	finalCards := []Card(parts[0])
	var bid int
	if n, err := strconv.Atoi(parts[1]); err == nil {
		bid = n
	}
	return Hand{
		cards: [5]Card(finalCards),
		bid:   bid,
	}
}

func (g Game) TotalWinnings() int {
	sort.Sort(g)
	var winnings int
	for i := 0; i < g.Len(); i++ {
		winnings += g[i].bid * (i + 1)
	}
	return winnings
}

func (g Game) Len() int {
	return len(g)
}

func (g Game) Less(i, j int) bool {
	handI := g[i]
	handJ := g[j]

	if handI.GetHandType() > handJ.GetHandType() {
		return false
	} else if handI.GetHandType() == handJ.GetHandType() {
		for i := 0; i < 5; i++ {
			iCard := handI.cards[i]
			jCard := handJ.cards[i]
			if iCard.Strength() > jCard.Strength() {
				return false
			} else if iCard.Strength() < jCard.Strength() {
				return true
			}
		}
	}
	return true
}

func (g Game) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

//func (h Hand) GetHandType() HandType {
//	setMap := make(map[Card]int)
//	containsJ := false
//	var jCount int
//	for _, card := range h.cards {
//		if !containsJ {
//			containsJ = card == CARD_J
//		}
//		if val, ok := setMap[card]; ok {
//			setMap[card] = val + 1
//		} else {
//			setMap[card] = 1
//		}
//	}
//	if containsJ {
//		jCount = setMap[CARD_J]
//		for card, _ := range setMap {
//			setMap[card] += jCount
//		}
//	}
//	length := len(setMap)
//	if length == 1 {
//		return FIVE_OF_A_KIND
//	}
//	if length == 2 {
//		for card := range setMap {
//			if setMap[card] == 4 || setMap[card] == 1 {
//				return FOUR_OF_A_KIND
//			}
//		}
//		return FULL_HOUSE
//	}
//	if length == 3 {
//		var highest int
//		for card := range setMap {
//			if highest == 0 || setMap[card] > highest {
//				highest = setMap[card]
//			}
//		}
//		switch highest {
//		case 2:
//			return TWO_PAIR
//		case 3:
//			return THREE_OF_A_KIND
//		}
//	}
//	if length == 4 {
//		return ONE_PAIR
//	}
//	return HIGH_CARD
//}

func (h Hand) GetHandType() HandType {
	setMap := make(map[Card]int)
	for _, card := range h.cards {
		if val, ok := setMap[card]; ok {
			setMap[card] = val + 1
		} else {
			setMap[card] = 1
		}
	}
	part1 := func() HandType {
		length := len(setMap)
		if length == 1 {
			return FIVE_OF_A_KIND
		}
		if length == 2 {
			for card := range setMap {
				if setMap[card] == 4 || setMap[card] == 1 {
					return FOUR_OF_A_KIND
				}
				return FULL_HOUSE
			}
		}
		if length == 3 {
			for card := range setMap {
				if setMap[card] == 2 {
					return TWO_PAIR
				}
				if setMap[card] == 3 {
					return THREE_OF_A_KIND
				}
			}
			return THREE_OF_A_KIND
		}
		if length == 4 {
			return ONE_PAIR
		}
		return HIGH_CARD
	}
	switch result := part1(); result {
	case FOUR_OF_A_KIND:
		if jCardNo, ok := setMap[CARD_J]; ok && (jCardNo == 1 || jCardNo == 4) {
			return FIVE_OF_A_KIND
		} else {
			return FOUR_OF_A_KIND
		}
	case THREE_OF_A_KIND:
		if _, ok := setMap[CARD_J]; ok {
			return FOUR_OF_A_KIND
		} else {
			return THREE_OF_A_KIND
		}
	case FULL_HOUSE:
		if _, ok := setMap[CARD_J]; ok {
			return FIVE_OF_A_KIND
		} else {
			return FULL_HOUSE
		}
	case TWO_PAIR:
		if jCardNo, ok := setMap[CARD_J]; ok {
			if jCardNo == 1 {
				return FULL_HOUSE
			}
			if jCardNo == 2 {
				return FOUR_OF_A_KIND
			}
		} else {
			return TWO_PAIR
		}
	case ONE_PAIR:
		if _, ok := setMap[CARD_J]; ok {
			return THREE_OF_A_KIND
		} else {
			return ONE_PAIR
		}
	case HIGH_CARD:
		if _, ok := setMap[CARD_J]; ok {
			return ONE_PAIR
		} else {
			return HIGH_CARD
		}
	default:
		return result
	}
	panic("should get here.. something is wrong")
}
