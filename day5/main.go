package main

import (
	"adventofcode/day5/alamac"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./day5/data.txt")
	if err != nil {
		panic(err)
	}
	fileContents := string(file)
	inputs := strings.Split(fileContents, "\r\n\r\n")
	seeds := alamac.GetSeeds(inputs[0])
	fmt.Println(seeds)
	seedToSoil := alamac.CreateMap(inputs[1])
	soilToFertilizer := alamac.CreateMap(inputs[2])
	fertilizerToWater := alamac.CreateMap(inputs[3])
	waterToLight := alamac.CreateMap(inputs[4])
	lightToTemperature := alamac.CreateMap(inputs[5])
	temperatureToHumidity := alamac.CreateMap(inputs[6])
	humiditiyToLocation := alamac.CreateMap(inputs[7])
	min := alamac.MinLocation(seeds, []alamac.ConversionMap{seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humiditiyToLocation})
	fmt.Println(min)
}
