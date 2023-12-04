package day4

import (
	"fmt"
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

	var totalPoints int
	for _, line := range lines {
		card := strings.Split(line, ":")
		nums := strings.Split(card[1], "|")

		// collect win numbers
		left := nums[0]
		wins := make(map[string]string)
		num, isDigit := "", false
		for _, char := range left {
			if char >= '0' && char <= '9' {
				if !isDigit {
					isDigit = true
				}
				num += string(char)
			} else {
				if isDigit {
					wins[num] = num
				}
				isDigit = false
				num = ""
			}
		}

		// check available numbers
		right := nums[1]
		var count, sum int
		num, isDigit = "", false
		for k, char := range right {
			if char >= '0' && char <= '9' {

				// if number then set to number
				if !isDigit {
					isDigit = true
				}
				// collect number
				num += string(char)

				// check if number locates at last index
				if k == len(right)-1 {
					_, ok := wins[num]
					if ok {
						if count == 0 {
							sum += 1
						} else {
							sum = sum * 2
						}
						count++
					}
				}
			} else {
				// If current 'char' is not not a number
				// then check isDigit is true (if true then previous is number)
				if isDigit {
					_, ok := wins[num]
					if ok {
						if count == 0 {
							sum += 1
						} else {
							sum = sum * 2
						}
						count++
					}
				}
				// reset isDigit and num
				isDigit = false
				num = ""
			}

		}

		fmt.Println(wins)
		fmt.Println("count: ", count)
		fmt.Println("sum: ", sum)
		fmt.Println()
		totalPoints += sum
	}
	return totalPoints
}
