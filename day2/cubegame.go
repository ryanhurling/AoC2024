package main

import (
	"adventofcode/day2/cube"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./day2/data.txt")
	if err != nil {
		panic(err)
	}
	fileContents := string(file)
	strs := strings.Split(fileContents, "\r\n")
	worker := cube.Worker{}
	games := worker.Extract(strs)
	var validGames []cube.Game
	for _, game := range games {
		valid := worker.Test(game)
		if valid {
			validGames = append(validGames, game)
		}
	}

	sum := 0
	for _, game := range validGames {
		sum += game.Id
	}
	var powerSum int64
	for _, game := range games {
		power := game.MinimumPossibleCubes().Power()
		powerSum += power
	}

	fmt.Println(sum)
	fmt.Println(powerSum)
}
