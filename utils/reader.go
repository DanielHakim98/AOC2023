package utils

import (
	"bufio"
	"fmt"
	"os"
)

type AocReader func(int, string) ([]string, error)

func GetInput(day int, filename string) ([]string, error) {
	path := fmt.Sprintf("./day%v/%v.txt", day, filename)
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
