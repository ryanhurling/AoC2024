package main

import (
	"adventofcode/day8/haunted"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./day8/data.txt")
	if err != nil {
		panic(err)
	}
	fileContents := string(file)
	inputs := strings.Split(fileContents, "\r\n")
	instructions := haunted.CreateInstructions(inputs[0])
	graph := haunted.CreateGraph(inputs[2:])
	steps := haunted.CountAllSteps(graph, instructions)
	fmt.Println(steps)
}
