package day5

import (
	"fmt"
	"log"

	"github.com/DanielHakim98/aoc/utils"
)

type Day5 struct {
	Dnum int
}

func (d *Day5) PartOne(filename string, reader utils.AocReader) int {
	lines, err := reader(d.Dnum, filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		fmt.Println(string(line))
	}
	return 0
}
