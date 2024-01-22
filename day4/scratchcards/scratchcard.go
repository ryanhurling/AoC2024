package scratchcards

import (
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type ScratchCard struct {
	id                           string
	winningNumbers, givenNumbers []int
	count                        int
}

type ScratchCardPile []*ScratchCard

func (c *ScratchCard) GetScore() uint {
	count := c.MatchedNumberCount()
	if count > 0 {
		return uint(math.Pow(2, float64(count-1)))
	}
	return 0
}

func (c *ScratchCard) MatchedNumberCount() uint {
	var count uint
	for _, number := range c.givenNumbers {
		if slices.Contains(c.winningNumbers, number) {
			count++
		}
	}
	return count
}

func (c *ScratchCard) IncrementCount() {
	c.count++
}

func (p ScratchCardPile) AddCounts(startingIndex, count int) {
	for i := 0; i < count; i++ {
		p[startingIndex+i].IncrementCount()
	}
}

func (p ScratchCardPile) GetScore() uint {
	var score uint
	for _, card := range p {
		score += card.GetScore()
	}
	return score
}

func (p ScratchCardPile) GetCardCount() uint {
	var count uint
	for i, card := range p {
		winningCount := card.MatchedNumberCount()
		if winningCount > 0 {
			for x := 0; x < card.count; x++ {
				p.AddCounts(i+1, int(winningCount))

			}
		}
		count += uint(card.count)
	}
	return count
}

func CreateScratchCardPile(inputs []string) ScratchCardPile {
	pile := make(ScratchCardPile, len(inputs))
	for i, input := range inputs {
		pile[i] = createScratchCard(input)
	}
	return pile
}

func createScratchCard(input string) *ScratchCard {
	parts := strings.Split(input, ": ")
	id := parts[0]
	possibleNumbers := parts[1]
	numberParts := strings.Split(possibleNumbers, " | ")

	//winningStrings := strings.Split(numberParts[0], " ")
	regex := regexp.MustCompile("(\\d+)")
	winningStrings := regex.FindAllString(numberParts[0], -1)
	winningNumbers := make([]int, len(winningStrings))
	for i, winningString := range winningStrings {
		n, _ := strconv.Atoi(winningString)
		winningNumbers[i] = n
	}

	givenStrings := regex.FindAllString(numberParts[1], -1)
	givenNumbers := make([]int, len(givenStrings))
	for i, givenString := range givenStrings {
		if n, err := strconv.Atoi(givenString); err == nil {
			givenNumbers[i] = n
		}
	}

	return &ScratchCard{id, winningNumbers, givenNumbers, 1}
}
