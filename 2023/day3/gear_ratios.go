package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func GetInput(filename string) ([]string, error) {
	path := fmt.Sprintf("./day3/%v.txt", filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

type Day3 struct{}

const (
	COORDINATES_TEMPLATE = "%v_%v"
)

func (d *Day3) PartOne(filename string, reader func(string) ([]string, error)) int {
	lines, err := reader(filename)
	if err != nil {
		log.Fatal(err)
	}

	symbolsRef := make(map[string][]int)
	for i, line := range lines {
		symbols := d.FindSymbol(line, i)
		for k, v := range symbols {
			symbolsRef[k] = v
		}
	}

	var sum int
	for i, line := range lines {
		numbers := d.FindNumberInRow(line, i)
		for _, num := range numbers {
		loopEveryDigit:
			for _, digit := range num.List {
				if isValid := d.CheckCoordinate(digit, symbolsRef); isValid {
					v, _ := strconv.Atoi(num.Value)
					sum += v
					break loopEveryDigit
				}
			}
		}
	}

	return sum
}

func (d *Day3) FindSymbol(line string, row int) map[string][]int {
	re := regexp.MustCompile("[^.0-9]+")
	matches := re.FindAllStringIndex(line, -1)

	symbols := make(map[string][]int)
	for _, match := range matches {
		first := match[0]
		symbols[fmt.Sprintf(COORDINATES_TEMPLATE, row, first)] = []int{row, first}
	}
	return symbols
}

type NumGroup struct {
	Value string
	List  map[string]Coordinates
}

type Coordinates []int

func (d *Day3) FindNumberInRow(line string, row int) []NumGroup {
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllStringIndex(line, -1)
	var numInRows []NumGroup

	for _, match := range matches {
		i := match[0]
		limit := match[1]
		fullNum := line[i:limit]
		nums := make(map[string]Coordinates)
		for ; i < limit; i++ {
			key := fmt.Sprintf("%v_%v", row, i)
			nums[key] = []int{row, i}
		}
		numInRows = append(numInRows, NumGroup{
			Value: fullNum,
			List:  nums,
		})
	}
	return numInRows
}

func (d *Day3) CheckCoordinate(coord Coordinates, symbolsRef map[string][]int) bool {
	row := coord[0]
	col := coord[1]

	up := fmt.Sprintf(COORDINATES_TEMPLATE, row-1, col)
	down := fmt.Sprintf(COORDINATES_TEMPLATE, row+1, col)
	left := fmt.Sprintf(COORDINATES_TEMPLATE, row, col-1)
	right := fmt.Sprintf(COORDINATES_TEMPLATE, row, col+1)
	diagRUp := fmt.Sprintf(COORDINATES_TEMPLATE, row-1, col+1)
	diagRDown := fmt.Sprintf(COORDINATES_TEMPLATE, row+1, col+1)
	diagLUp := fmt.Sprintf(COORDINATES_TEMPLATE, row-1, col-1)
	diagLDown := fmt.Sprintf(COORDINATES_TEMPLATE, row+1, col-1)

	updatedCoords := []string{up, down, left, right, diagRUp, diagRDown, diagLUp, diagLDown}
	for _, coords := range updatedCoords {
		_, ok := symbolsRef[coords]
		if ok {
			return true
		}
	}
	return false
}

/* Might be useful later (I don't know)
type BoardSize struct {
	Row int
	Col int
}

func (d *Day3) GetSize(lines *[]string) BoardSize {
	row := len(*lines)
	if row == 0 {
		return BoardSize{}
	}

	col := len((*lines)[0])

	return BoardSize{
		Row: row,
		Col: col,
	}
}
*/

func (d *Day3) PartTwo(filename string, reader func(string) ([]string, error)) int {
	lines, err := reader(filename)
	if err != nil {
		log.Fatal(err)
	}

	numberRefs := make(map[string]PartNumber)
	symbolsRef := make(map[string][]int)
	for i, line := range lines {
		symbols := d.FindAsteriskSymbol(line, i)
		for k, v := range symbols {
			symbolsRef[k] = v
		}

		numInRows := d.FindNumberInRow(line, i)
		for _, num := range numInRows {
			for k, v := range num.List {
				numberRefs[k] = PartNumber{
					Value:       num.Value,
					Coordinates: v,
				}
			}
		}
	}

	var sum int
	for _, v := range symbolsRef {
		nums := d.CheckGearNearby(v, numberRefs)
		if len(nums) == 2 {
			v1, _ := strconv.Atoi(nums[0])
			v2, _ := strconv.Atoi(nums[1])

			sum += v1 * v2
		}
	}
	return sum
}

func (d *Day3) FindAsteriskSymbol(line string, row int) map[string][]int {
	re := regexp.MustCompile("[*]")
	matches := re.FindAllStringIndex(line, -1)

	symbols := make(map[string][]int)
	for _, match := range matches {
		first := match[0]
		symbols[fmt.Sprintf(COORDINATES_TEMPLATE, row, first)] = []int{row, first}
	}
	return symbols
}

type PartNumber struct {
	Value string
	Coordinates
}

func (d *Day3) CheckGearNearby(coord Coordinates, numbersRef map[string]PartNumber) []string {
	row := coord[0]
	col := coord[1]

	up := fmt.Sprintf(COORDINATES_TEMPLATE, row-1, col)
	down := fmt.Sprintf(COORDINATES_TEMPLATE, row+1, col)
	left := fmt.Sprintf(COORDINATES_TEMPLATE, row, col-1)
	right := fmt.Sprintf(COORDINATES_TEMPLATE, row, col+1)
	diagRUp := fmt.Sprintf(COORDINATES_TEMPLATE, row-1, col+1)
	diagRDown := fmt.Sprintf(COORDINATES_TEMPLATE, row+1, col+1)
	diagLUp := fmt.Sprintf(COORDINATES_TEMPLATE, row-1, col-1)
	diagLDown := fmt.Sprintf(COORDINATES_TEMPLATE, row+1, col-1)

	var digits []string
	var count int
	groups := make(map[string]struct{})
	updatedCoords := []string{up, down, left, right, diagRUp, diagRDown, diagLUp, diagLDown}
	for _, coords := range updatedCoords {
		num, ok := numbersRef[coords]
		_, exists := groups[num.Value]
		if ok && !exists {
			digits = append(digits, num.Value)
			groups[num.Value] = struct{}{}
			count++
		}
	}

	if len(digits) != 2 {
		return []string{}
	}

	return digits
}
