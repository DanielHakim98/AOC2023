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

	// size := d.GetSize(&lines)
	// fmt.Println(size)

	lastIdx := len(lines) - 1
	for i, line := range lines {
		currentNums := d.FindNumberInRow(line, i)
		if i == 0 {
			nextNums := d.FindNumberInRow(lines[i+1], i+1)
			fmt.Println("currentNums: ", currentNums)
			fmt.Println("nextNums: ", nextNums)
			fmt.Println()
		} else if i == lastIdx {
			prevNums := d.FindNumberInRow(lines[i-1], i-1)
			fmt.Println("prevNums: ", prevNums)
			fmt.Println("currentNums: ", currentNums)
			fmt.Println()
		} else {
			prevNums := d.FindNumberInRow(lines[i-1], i-1)
			nextNums := d.FindNumberInRow(lines[i+1], i+1)
			fmt.Println("prevNums: ", prevNums)
			fmt.Println("currentNums: ", currentNums)
			fmt.Println("nextNums: ", nextNums)
			fmt.Println()
		}

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

type NumInRow struct {
	Values      string
	Coordinates []int
}

func (nr NumInRow) String() string {
	return fmt.Sprintf("NumInRow( Values: %v, Coordinates: %v)", nr.Values, nr.Coordinates)
}

func (d *Day3) FindNumberInRow(line string, row int) []NumInRow {
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllStringIndex(line, -1)
	var nums []NumInRow
	for _, match := range matches {
		first := match[0]
		last := match[1]
		num := line[first:last]
		nums = append(nums, NumInRow{
			Values: num,
			Coordinates: []int{
				row,
				first,
			},
		})
	}
	return nums
}
