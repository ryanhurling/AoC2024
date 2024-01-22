package cube

import (
	"strconv"
	"strings"
)

type Colour string
type RoundString string
type TurnString string

const (
	Red   Colour = "red"
	Green Colour = "green"
	Blue  Colour = "blue"
)

const (
	red_Max   = 12
	green_Max = 13
	blue_Max  = 14
)

type Turns []Turn
type Rounds []Round

type Game struct {
	Id     int
	rounds Rounds
}

type MinimumNeeded struct {
	Red   Turn
	Green Turn
	Blue  Turn
}

func (m MinimumNeeded) Power() int64 {
	return int64(m.Green.number * m.Red.number * m.Blue.number)
}

func (g Game) MinimumPossibleCubes() MinimumNeeded {
	red := Turn{colour: Red}
	blue := Turn{colour: Blue}
	green := Turn{colour: Green}

	for _, rounds := range g.rounds {
		for _, turn := range rounds.turns {
			switch turn.colour {
			case Red:
				if turn.number > red.number {
					red = turn
				}
				break
			case Blue:
				if turn.number > blue.number {
					blue = turn
				}
				break
			case Green:
				if turn.number > green.number {
					green = turn
				}
			}
		}
	}
	return MinimumNeeded{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func NewGame(gameString string) *Game {
	parts := strings.Split(gameString, ": ")
	id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
	roundStrings := strings.Split(parts[1], "; ")
	rounds := make(Rounds, len(roundStrings))
	for i, roundString := range roundStrings {
		round := *NewRound(roundString)
		rounds[i] = round
	}
	return &Game{id, rounds}
}

type Round struct {
	turns Turns
}

func NewRound(roundString string) *Round {
	convertedString := RoundString(roundString)
	turns := convertedString.ExtractTurns()
	return &Round{turns}
}

func (s RoundString) ExtractTurns() Turns {
	trimmedColours := strings.TrimRight(string(s), ";")
	colourTurnStrings := strings.Split(trimmedColours, ", ")
	turns := make(Turns, len(colourTurnStrings))
	for i, turnString := range colourTurnStrings {
		turns[i] = *NewTurn(turnString)
	}
	return turns
}

type Turn struct {
	colour Colour
	number int
}

func NewTurn(turnString string) *Turn {
	parts := strings.Split(turnString, " ")
	atoi, err := strconv.Atoi(parts[0])
	if err != nil {
		panic("something wrong")
	}
	return &Turn{
		colour: Colour(parts[1]),
		number: atoi,
	}
}

type Extractor interface {
	Extract(gameString []string) []Game
}

type GameTester interface {
	Test(game Game) bool
}

type Worker struct {
}

func (w Worker) Test(game Game) bool {
	for _, round := range game.rounds {
		for _, turn := range round.turns {
			switch turn.colour {
			case Red:
				if turn.number > red_Max {
					return false
				}
				break
			case Blue:
				if turn.number > blue_Max {
					return false
				}
				break
			case Green:
				if turn.number > green_Max {
					return false
				}
				break
			}

		}
	}
	return true
}

func (w Worker) Extract(gameString []string) []Game {
	games := make([]Game, len(gameString))
	for i, s := range gameString {
		games[i] = *NewGame(s)
	}
	return games
}
