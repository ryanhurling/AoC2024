package trebuchet

import (
	"strconv"
	"strings"
	"unicode"
)

type character rune
type numberSlice []int32

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

var startingLetters = getStartingLettersMap()

type NumberConversion struct {
	word   string
	number int32
}

func getStartingLettersMap() map[string][]NumberConversion {
	var temp = make(map[string][]NumberConversion)
	temp["o"] = []NumberConversion{{word: one, number: 1}}
	temp["t"] = []NumberConversion{{word: two, number: 2}, {word: three, number: 3}}
	temp["f"] = []NumberConversion{{word: four, number: 4}, {word: five, number: 5}}
	temp["s"] = []NumberConversion{{word: six, number: 6}, {word: seven, number: 7}}
	temp["e"] = []NumberConversion{{word: eight, number: 8}}
	temp["n"] = []NumberConversion{{word: nine, number: 9}}
	return temp
}

func ConvertWordToDigit(word string, startingIndex int32) (int32, int32) {
	chars := strings.Split(word, "")
	parts := chars[startingIndex:]
	possibleNumbers := startingLetters[parts[0]]
	if len(possibleNumbers) == 0 {
		return -1, 0
	}
	if len(possibleNumbers) == 1 {
		firstPossible := possibleNumbers[0].word
		var firstParts string
		if len(parts) >= len(firstPossible) {
			firstParts = strings.Join(parts[:len(firstPossible)], "")
			if firstParts == firstPossible {
				return possibleNumbers[0].number, int32(len(possibleNumbers[0].word) - 1)
			}
		}
		return -1, 0
	}
	firstPossible := possibleNumbers[0].word
	var firstParts string
	if len(parts) >= len(firstPossible) {
		firstParts = strings.Join(parts[:len(firstPossible)], "")
		if firstParts == firstPossible {
			return possibleNumbers[0].number, int32(len(possibleNumbers[0].word) - 1)
		}
	}
	firstPossible = possibleNumbers[1].word
	if len(parts) >= len(firstPossible) {
		firstParts = strings.Join(parts[:len(firstPossible)], "")
		if firstParts == firstPossible {
			return possibleNumbers[1].number, int32(len(possibleNumbers[1].word) - 1)
		}
	}
	return -1, 0
}

func CalculateCalibration(values []string) (int32, error) {
	var firstNumber, secondNumber int32
	numbersToSum := numberSlice(make([]int32, len(values)))
	for i, word := range values {
		for x, char := range word {

			if character(char).isNumeric() {
				if firstNumber == 0 {
					toi, err := rToi(char)
					if err != nil {
						return 0, err
					}
					firstNumber = toi
				} else {
					toi, err := rToi(char)
					if err != nil {
						return 0, err
					}
					secondNumber = toi
				}
			} else {
				number, skip := ConvertWordToDigit(word, int32(x))
				if skip > 0 {
					if firstNumber == 0 {
						firstNumber = number
					} else {
						secondNumber = number
					}
				}
			}
			if firstNumber != 0 && secondNumber != 0 {
				convFirst := strconv.Itoa(int(firstNumber))
				convSecond := strconv.Itoa(int(secondNumber))
				temp := convFirst + convSecond
				convertedNumber, err := strconv.Atoi(temp)
				if err != nil {
					return -1, err
				}
				numbersToSum[i] = int32(convertedNumber)
			}
		}
		if firstNumber != 0 && secondNumber == 0 {
			convFirst := strconv.Itoa(int(firstNumber))
			temp := convFirst + convFirst
			convertedNumber, err := strconv.Atoi(temp)
			if err != nil {
				return -1, err
			}
			numbersToSum[i] = int32(convertedNumber)
		}
		firstNumber = 0
		secondNumber = 0
	}
	return numbersToSum.reduce(), nil
}

func rToi[T ~rune](value T) (int32, error) {
	atoi, err := strconv.Atoi(string(value))
	if err != nil {
		return 0, err
	}
	return int32(atoi), nil
}

func (n numberSlice) reduce() int32 {
	var result int32
	for _, value := range n {
		result += value
	}
	return result
}

func (c character) isNumeric() bool {
	return unicode.IsNumber(rune(c))
}
