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

func (d *Day3) PartOne(filename string, reader func(string) ([]string, error)) int {
	lines, err := reader(filename)
	if err != nil {
		log.Fatal(err)
	}

	for i, line := range lines {
		fmt.Println(d.FindSymbol(line, i))
		fmt.Println(d.FindNumberInRow(line, i))
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
