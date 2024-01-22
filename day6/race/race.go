package race

import (
	"strconv"
	"strings"
)

type Race struct {
	time, recordDistance int
}

type Races []Race

func CreateRaces(times, distances []string) Races {
	//races := make([]Race, len(distances))

	newTime := strings.Join(times, "")
	newDistance := strings.Join(distances, "")
	var raceTime, raceDistance int
	if convertedTime, err := strconv.Atoi(newTime); err == nil {
		raceTime = convertedTime
	}
	if convertedDistance, err := strconv.Atoi(newDistance); err == nil {
		raceDistance = convertedDistance
	}
	//case 1
	//for i, time := range times {
	//	var raceTime, raceDistance int
	//	if convertedTime, err := strconv.Atoi(time); err == nil {
	//		raceTime = convertedTime
	//	}
	//	if convertedDistance, err := strconv.Atoi(distances[i]); err == nil {
	//		raceDistance = convertedDistance
	//	}
	//
	//	races[i] = Race{
	//		time:           raceTime,
	//		recordDistance: raceDistance,
	//	}
	//}
	return Races{
		Race{
			time:           raceTime,
			recordDistance: raceDistance,
		},
	}
}

func (r Race) GetLowestPossibleHeldTime() int {
	target := r.recordDistance
	for i := 1; i < r.time; i++ {
		delta := r.time - i
		actual := i * delta
		if actual > target {
			return i
		}
	}
	return 0
}

func (r Race) GetHighestPossibleHeldTime() int {
	lowestTime := r.GetLowestPossibleHeldTime()
	return r.time - lowestTime
}

func (r Race) TotalPossibleWinCases() int {
	lowest := r.GetLowestPossibleHeldTime()
	highest := r.GetHighestPossibleHeldTime()

	return highest - lowest + 1
}

func (r Races) TotalPossibleNoOfCases() int {
	var total int
	for _, race := range r {
		if total == 0 {
			total = race.TotalPossibleWinCases()
		} else {
			total *= race.TotalPossibleWinCases()
		}
	}
	return total
}
