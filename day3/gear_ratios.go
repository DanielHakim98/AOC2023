package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	SHOW_PREV_NUMS       = "prevNums: "
	SHOW_CURRENT_NUMS    = "currentNums: "
	SHOW_NEXT_NUMS       = "nextNums: "
	SHOW_PREV_SYMBOLS    = "prevSymbols: "
	SHOW_CURRENT_SYMBOLS = "currentSymbols: "
	SHOW_NEXT_SYMBOLS    = "nextSymbols: "
)

func (d *Day3) PartOne(filename string, reader func(string) ([]string, error)) int {
	lines, err := reader(filename)
	if err != nil {
		log.Fatal(err)
	}

	lastIdx := len(lines) - 1
	for i, line := range lines {
		currentNums := d.FindNumberInRow(line, i)
		currentSymbols := d.FindSymbol(line, i)
		if i == 0 {
			fmt.Println(SHOW_CURRENT_NUMS, currentNums)
			fmt.Println(SHOW_CURRENT_SYMBOLS, currentSymbols)

			nextNums := d.FindNumberInRow(lines[i+1], i+1)
			fmt.Println(SHOW_NEXT_NUMS, nextNums)
			nextSymbols := d.FindSymbol(lines[i+1], i+1)
			fmt.Println(SHOW_NEXT_SYMBOLS, nextSymbols)

		} else if i == lastIdx {
			prevNums := d.FindNumberInRow(lines[i-1], i-1)
			fmt.Println(SHOW_PREV_NUMS, prevNums)
			prevSymbols := d.FindSymbol(lines[i-1], i-1)
			fmt.Println(SHOW_PREV_SYMBOLS, prevSymbols)

			fmt.Println(SHOW_CURRENT_NUMS, currentNums)
			fmt.Println(SHOW_CURRENT_SYMBOLS, currentSymbols)

		} else {
			prevNums := d.FindNumberInRow(lines[i-1], i-1)
			fmt.Println(SHOW_PREV_NUMS, prevNums)
			prevSymbols := d.FindSymbol(lines[i-1], i-1)
			fmt.Println(SHOW_PREV_SYMBOLS, prevSymbols)

			fmt.Println(SHOW_CURRENT_NUMS, currentNums)
			fmt.Println(SHOW_CURRENT_SYMBOLS, currentSymbols)

			nextNums := d.FindNumberInRow(lines[i+1], i+1)
			fmt.Println(SHOW_NEXT_NUMS, nextNums)
			nextSymbols := d.FindSymbol(lines[i+1], i+1)
			fmt.Println(SHOW_NEXT_SYMBOLS, nextSymbols)

		}
		fmt.Println()
	}
	return 0
}

type BoardSize struct {
	Row int
	Col int
}

func (bd BoardSize) String() string {
	return fmt.Sprintf("BoardSize(%v, %v)", bd.Row, bd.Col)
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

type NumGroup struct {
	Value string
	List  map[NumSingle]Coordinates
}

type Coordinates []int
type NumSingle struct {
	Value    string
	Row, Col int
}

func (nr NumSingle) String() string {
	return fmt.Sprintf("{v: \"%v\"}", nr.Value)
}

func (d *Day3) FindNumberInRow(line string, row int) []NumGroup {
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllStringIndex(line, -1)
	var numInRows []NumGroup

	for _, match := range matches {
		i := match[0]
		limit := match[1]
		fullNum := line[i:limit]
		nums := make(map[NumSingle]Coordinates)
		for ; i < limit; i++ {
			val := string(line[i])
			num := NumSingle{
				Value: val,
				Row:   row,
				Col:   i,
			}
			nums[num] = []int{num.Row, num.Col}
		}
		numInRows = append(numInRows, NumGroup{
			Value: fullNum,
			List:  nums,
		})
	}
	return numInRows
}

type Symbol struct {
	Value       string
	Coordinates []int
}

func (d *Day3) FindSymbol(line string, row int) []Symbol {
	re := regexp.MustCompile("[^.0-9]+")
	matches := re.FindAllStringIndex(line, -1)
	var symbols []Symbol
	for _, match := range matches {
		first := match[0]
		last := match[1]
		symbol := line[first:last]
		symbols = append(symbols, Symbol{
			Value:       symbol,
			Coordinates: []int{row, first},
		})
	}
	return symbols
}
