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
	defer file.Close()

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

// only 12 red cubes, 13 green cubes, and 14 blue cubes?
var bags = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Day2(filename string, reader func(string) ([]string, error)) int {
	games, err := reader(filename)
	if err != nil {
		log.Fatal(err)
	}
	g := ParseGame(games[0])
	fmt.Println(g)
	return 0
}

type Game struct {
	Id   int
	Sets []map[string]int
}

func ParseGame(game string) Game {
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
