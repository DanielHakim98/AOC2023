package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput(filename string) ([]string, error) {
	path := fmt.Sprintf("./day2/%v.txt", filename)
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

type Day2 struct{}

func (d *Day2) PartOne(filename string, reader func(string) ([]string, error)) int {
	games, err := reader(filename)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, game := range games {
		g := d.ParseGame(game)
		if valid := d.CheckGame(g); valid {
			sum += g.Id
		}
	}

	return sum
}

type Game struct {
	Id   int
	Sets []map[string]int
}

func (day2 *Day2) ParseGame(game string) Game {
	g := strings.Split(game, ":")

	// Get Id
	title := strings.Split(strings.TrimSpace(g[0]), " ")
	id, _ := strconv.Atoi(title[len(title)-1])

	// Collect set
	collected := make([]map[string]int, 0)
	sets := strings.Split(strings.TrimSpace(g[1]), "; ")
	for _, set := range sets {
		collect := make(map[string]int)
		subsets := strings.Split(set, ", ")
		for _, subset := range subsets {
			s := strings.Split(subset, " ")
			color := s[len(s)-1]
			amount, _ := strconv.Atoi(s[0])
			collect[color] = amount
		}
		collected = append(collected, collect)
	}

	return Game{
		Id:   id,
		Sets: collected,
	}
}

// only 12 red cubes, 13 green cubes, and 14 blue cubes
var bags = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func (day2 *Day2) CheckGame(game Game) bool {
	for _, set := range game.Sets {
		greenInvalid := set["green"] > bags["green"]
		blueInvalid := set["blue"] > bags["blue"]
		redInvalid := set["red"] > bags["red"]

		if greenInvalid || blueInvalid || redInvalid {
			return false
		}
	}
	return true
}

func (d *Day2) PartTwo(filename string, reader func(string) ([]string, error)) int {
	return 0
}
