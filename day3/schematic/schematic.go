package schematic

import (
	"strconv"
	"unicode"
)

type Schematic [][]rune

type Coordinate struct {
	x, y   int
	marked bool
}

type Gear struct {
	x, y int
}

func (g Gear) Ratio() int {
	return g.x * g.y
}

func Create2dArray(input []string) Schematic {
	container := make([][]rune, len(input))
	for i, s := range input {
		line := make([]rune, len(s))
		for i2, i3 := range s {
			line[i2] = i3
		}
		container[i] = line
	}
	return container
}

func isSpecialCharacter(char rune) (special bool, gear bool) {
	return char != '.' && (unicode.IsPunct(char) || unicode.IsSymbol(char)), isPossibleGear(char)
}

func isPossibleGear(char rune) bool {
	return char == '*'
}

func Scan(schematic Schematic) []int {
	results := make([]int, 0)
	for x, line := range schematic {
		for y, char := range line {
			if isSpecial, isGear := isSpecialCharacter(char); isSpecial {
				coorindates := schematic.GetCoordinatesAroundPoint(x, y)
				//foundNumbers := scanCoordinates(coorindates, schematic)
				if isGear {
					gears := scanCoordinatesForGears(coorindates, schematic)
					for _, gear := range gears {
						results = append(results, gear.Ratio())
					}
				}
				//results = append(results, foundNumbers...)
			}

		}
	}
	return results
}

func Reduce(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func scanCoordinatesForGears(coordinates []Coordinate, schematic Schematic) []Gear {
	var gears []Gear
	nums := scanCoordinates(coordinates, schematic)
	if len(nums) == 2 {
		gears = append(gears, Gear{
			x: nums[0],
			y: nums[1],
		})
	}
	return gears
}

func scanCoordinates(coordinates []Coordinate, schematic Schematic) []int {
	yLength := len(schematic[0])
	result := make([]int, 0)
	for _, coord := range coordinates {
		v := schematic[coord.x][coord.y]

		prefix := ""
		tail := ""
		if coord.marked {
			continue
		}
		if is, _, _ := isNumeric(v); is {
			coord.marked = true
			if coord.y > 0 {
				y := coord.y - 1
				value := schematic[coord.x][y]
				isNum, num, _ := isNumeric(value)
				for y >= 0 && isNum {
					prefix = strconv.Itoa(num) + prefix
					y--

					if y >= 0 {
						value = schematic[coord.x][y]
						schematic[coord.x][y] = 'd'
						isNum, num, _ = isNumeric(value)
					}
				}

			}
			if coord.y < yLength-1 {
				y := coord.y + 1
				value := schematic[coord.x][y]
				schematic[coord.x][y] = 'd'
				isNum, num, _ := isNumeric(value)
				for y < yLength && isNum {
					tail += strconv.Itoa(num)
					y++
					if y < yLength {
						value = schematic[coord.x][y]
						schematic[coord.x][y] = 'd'
						isNum, num, _ = isNumeric(value)
					}
				}
			}
			temp := prefix + string(schematic[coord.x][coord.y]) + tail
			atoi, _ := strconv.Atoi(temp)
			schematic[coord.x][coord.y] = 'd'
			result = append(result, atoi)
		}
	}

	return result
}

func isNumeric(char rune) (bool, int, error) {
	n, err := strconv.Atoi(string(char))
	if err != nil {
		return false, 0, err
	}
	return true, n, nil
}

func (s Schematic) GetCoordinatesAroundPoint(x, y int) []Coordinate {
	var coordinates []Coordinate

	// Adding conditionals to take the edges into account
	if x > 0 && y > 0 {
		coordinates = append(coordinates, Coordinate{x - 1, y - 1, false})
	}
	if x > 0 {
		coordinates = append(coordinates, Coordinate{x - 1, y, false})
	}
	if x > 0 && y < len(s[0])-1 {
		coordinates = append(coordinates, Coordinate{x - 1, y + 1, false})
	}
	if y > 0 {
		coordinates = append(coordinates, Coordinate{x, y - 1, false})
	}
	if y < len(s[0])-1 {
		coordinates = append(coordinates, Coordinate{x, y + 1, false})
	}
	if x < len(s)-1 && y > 0 {
		coordinates = append(coordinates, Coordinate{x + 1, y - 1, false})
	}
	if x < len(s)-1 {
		coordinates = append(coordinates, Coordinate{x + 1, y, false})
	}
	if x < len(s)-1 && y < len(s[0])-1 {
		coordinates = append(coordinates, Coordinate{x + 1, y + 1, false})
	}

	return coordinates
}
