package day4

import (
	"log"
	"strings"

	"github.com/DanielHakim98/aoc/utils"
)

type Day4 struct{}

func (d *Day4) PartOne(filename string, reader utils.AocReader) int {
	lines, err := reader(4, filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		card := strings.Split(line, ":")
		nums := strings.Split(card[1], "|")
		left := nums[0]

		// collect win numbers
		var wins []string
		num, isDigit := "", false
		for _, char := range left {
			if char >= '0' && char <= '9' {
				if !isDigit {
					isDigit = true
				}
				num += string(char)
			} else {
				if isDigit {
					wins = append(wins, num)
				}
				isDigit = false
				num = ""
			}
		}

	}
	return 0
}
