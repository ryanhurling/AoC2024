package main

import (
	"adventofcode/day1/trebuchet"
	"fmt"
	"os"
	"strings"
)

// correct answer is 53515
func main() {
	file, err := os.ReadFile("./day1/data.txt")
	if err != nil {
		panic(err)
	}
	fileContents := string(file)
	inputs := strings.Split(fileContents, "\r\n")
	calibration, err := trebuchet.CalculateCalibration(inputs)
	if err != nil {
		panic(err)
	}
	fmt.Println(calibration)
}
