package main

import (
	"adventofcode/day6/race"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.ReadFile("./day6/data.txt")
	if err != nil {
		panic(err)
	}
	fileContents := string(file)
	inputs := strings.Split(fileContents, "\r\n")
	regex := regexp.MustCompile("(\\d+)")
	times := regex.FindAllString(inputs[0], -1)
	distances := regex.FindAllString(inputs[1], -1)
	races := race.CreateRaces(times, distances)
	total := races.TotalPossibleNoOfCases()
	fmt.Println(total)
}
