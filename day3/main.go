package main

import (
	"adventofcode/day3/schematic"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./day3/data.txt")
	if err != nil {
		panic(err)
	}
	fileContents := string(file)
	inputs := strings.Split(fileContents, "\r\n")
	s := schematic.Create2dArray(inputs)
	nums := schematic.Scan(s)
	fmt.Println(schematic.Reduce(nums))
}
