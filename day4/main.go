package main

import (
	"adventofcode/day4/scratchcards"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./day4/data.txt")
	if err != nil {
		panic(err)
	}
	fileContents := string(file)
	inputs := strings.Split(fileContents, "\r\n")
	pile := scratchcards.CreateScratchCardPile(inputs)
	fmt.Println(pile.GetCardCount())
}
