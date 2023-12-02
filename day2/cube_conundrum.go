package day2

import (
	"bufio"
	"fmt"
	"os"
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

func Day2(filename string, reader func(string) ([]string, error)) int {
	return 0
}
