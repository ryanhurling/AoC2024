package alamac

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type ConversionMap struct {
	entries []ConversionMapEntry
}

type ConversionMapEntry struct {
	destinationRangeStart, sourceRangeStart, rangeLength int
}

type SeedRange struct {
	startIndex, amount int
}

func (m ConversionMap) IsMapped(sourceNo int) bool {
	for _, entry := range m.entries {
		if entry.sourceRangeStart <= sourceNo && sourceNo <= entry.sourceRangeStart+(entry.rangeLength-1) {
			return true
		}
	}
	return false
}

func (m ConversionMap) IsMappedWithEntry(sourceNo int) (bool, ConversionMapEntry) {
	for _, entry := range m.entries {
		if entry.sourceRangeStart <= sourceNo && sourceNo <= entry.sourceRangeStart+(entry.rangeLength-1) {
			return true, entry
		}
	}
	return false, ConversionMapEntry{}
}

func (m ConversionMap) GetConvertedIndex(sourceNo int) int {
	if mapped, entry := m.IsMappedWithEntry(sourceNo); mapped {
		difference := sourceNo - entry.sourceRangeStart
		return difference + entry.destinationRangeStart
	}
	return sourceNo
}

func CreateConversionMap(input []string) ConversionMap {
	entries := make([]ConversionMapEntry, len(input))
	wg := &sync.WaitGroup{}
	wg.Add(len(input))
	for i, s := range input {
		go func(in string, index int, group *sync.WaitGroup) {
			entries[index] = createConversionEntry(in)
			group.Done()
		}(s, i, wg)
	}
	wg.Wait()
	return ConversionMap{entries}
}

func createConversionEntry(input string) ConversionMapEntry {
	parts := strings.Split(input, " ")
	var destRangeStart, sourceRangeStart, rangeLength int
	if i, err := strconv.Atoi(parts[0]); err == nil {
		destRangeStart = i
	}

	if i, err := strconv.Atoi(parts[1]); err == nil {
		sourceRangeStart = i
	}

	if i, err := strconv.Atoi(parts[2]); err == nil {
		rangeLength = i
	}

	return ConversionMapEntry{
		destinationRangeStart: destRangeStart,
		sourceRangeStart:      sourceRangeStart,
		rangeLength:           rangeLength,
	}
}

func GetSeeds(input string) []SeedRange {
	regex := regexp.MustCompile("(\\d+\\s\\d+)")
	numberStrings := regex.FindAllString(input, -1)
	seeds := make([]SeedRange, len(numberStrings))
	for i, numberString := range numberStrings {
		parts := strings.Split(numberString, " ")
		var start, amount int
		if num, err := strconv.Atoi(parts[0]); err == nil {
			start = num
		}
		if num, err := strconv.Atoi(parts[1]); err == nil {
			amount = num
		}
		seeds[i] = SeedRange{
			startIndex: start,
			amount:     amount,
		}
	}
	return seeds
}

func CreateMap(input string) ConversionMap {
	strs := strings.Split(input, "\r\n")
	return CreateConversionMap(strs[1:])
}

func MinLocation(seeds []SeedRange, maps []ConversionMap) int {
	minLocation := make([]int, len(seeds))
	wg := sync.WaitGroup{}
	wg.Add(len(seeds))
	for seedIndex, seed := range seeds {
		go func(s SeedRange, sI int) {
			upperLimit := s.startIndex + s.amount
			for i := s.startIndex; i < upperLimit; i++ {
				input := i
				for _, conversionMap := range maps {
					input = conversionMap.GetConvertedIndex(input)
				}
				if input < minLocation[sI] || minLocation[sI] == 0 {
					minLocation[sI] = input
				}
			}
			wg.Done()
		}(seed, seedIndex)
	}
	wg.Wait()

	var finalMin int
	for _, value := range minLocation {
		if value < finalMin || finalMin == 0 {
			finalMin = value
		}
	}

	return finalMin
}
