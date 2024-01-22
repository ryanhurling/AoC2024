package main

import (
	"adventofcode/day7/camel"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./day7/data.txt")
	if err != nil {
		panic(err)
	}
	fileContents := string(file)
	inputs := strings.Split(fileContents, "\r\n")
	hands := make([]camel.Hand, len(inputs))
	for i, str := range inputs {
		hands[i] = camel.CreateHand(str)
	}
	fmt.Println(camel.Game(hands).TotalWinnings())
}
