package day4

import (
	"fmt"
	"log"
	"strconv"
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

type ScratchCard struct {
	Id      int
	Matches int
}

type Pair struct {
	Current ScratchCard
	Next    *[]Pair
}

type CardId int

var count int

func (d *Day4) PartTwo(filename string, reader utils.AocReader) int {
	lines, err := reader(4, filename)
	if err != nil {
		log.Fatal(err)
	}

	scratchCards := make(map[CardId]ScratchCard)
	for _, line := range lines {
		card := strings.Split(line, ":")
		nums := strings.Split(card[1], "|")
		// collect total wins
		left := nums[0]
		wins := make(map[string]string)
		num, isDigit := "", false
		for _, char := range left {
			if char >= '0' && char <= '9' {
				if !isDigit {
					isDigit = true
				}
				num += string(char)
				// we dont' consider last index here because
				// it's always empty spaces
			} else {
				if isDigit {
					wins[num] = num
				}
				isDigit = false
				num = ""
			}
		}

		right := nums[1]
		var count int
		num, isDigit = "", false
		for k, char := range right {
			if char >= '0' && char <= '9' {
				// if number then set to number
				if !isDigit {
					isDigit = true
				}
				// collect number
				num += string(char)
				// check if a number locates at last index
				if k == len(right)-1 {
					_, ok := wins[num]
					if ok {
						count++
					}
				}
			} else {
				// If current 'char' is not not a number
				// then check isDigit is true (if true then previous is number)
				if isDigit {
					_, ok := wins[num]
					if ok {
						count++
					}
				}
				// reset isDigit and num
				isDigit = false
				num = ""
			}
		}

		id := d.GetCardNumber(card)
		scratchCards[CardId(id)] = ScratchCard{id, count}

	}

	for _, c := range scratchCards {
		d.CountCards(c, scratchCards)

	}

	return count
}

func (d *Day4) GetCardNumber(card []string) (id int) {
	id, _ = strconv.Atoi(strings.TrimSpace(strings.Split(card[0], " ")[1]))
	return
}

func (d *Day4) CountCards(c ScratchCard, ref map[CardId]ScratchCard) {
	fmt.Println(c)
	if c.Matches == 0 {
		return
	}

	limit := c.Id + c.Matches
	cur := c.Id + 1

	for ; cur <= limit; cur++ {
		curCard, ok := ref[CardId(cur)]
		fmt.Println(curCard)
		if ok {
			if c.Matches > 0 {
				d.CountCards(curCard, ref)
			}

		}
	}

}
