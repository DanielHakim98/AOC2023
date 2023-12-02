package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Day1() {
	// 1. Open file
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// 2. Read line by line
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// digitInLetters := map[string]string{
	// 	"one":   "1",
	// 	"two":   "2",
	// 	"three": "3",
	// 	"four":  "4",
	// 	"five":  "5",
	// 	"six":   "6",
	// 	"seven": "7",
	// 	"eight": "8",
	// 	"nine":  "9",
	// }
	// keys := make([]string, 0)
	// for k := range digitInLetters {
	// 	keys = append(keys, k)
	// }
	// cleanedLines := make([]string, 0)
	// for _, line := range lines {
	// 	temp := line
	// 	for _, key := range keys {
	// 		if strings.Contains(temp, key) {
	// 			fmt.Printf("line '%v' contains %v\n", temp, digitInLetters[key])
	// 			v := strings.ReplaceAll(temp, key, digitInLetters[key])
	// 			fmt.Println(v)
	// 			temp = v
	// 		}
	// 	}
	// 	fmt.Println("done: ", temp, "\n")
	// 	cleanedLines = append(cleanedLines, temp)
	// }

	// for _, v := range cleanedLines {
	// 	fmt.Println(v)
	// }
	// 3. Get first occurence of first digit and last occurences of last digit
	calibrators := make([]string, 0)
	for _, line := range lines {
		temp := []string{}
		for _, char := range line {
			if unicode.IsDigit(char) {
				temp = append(temp, string(char))
			}
		}
		if len(temp) == 0 {
			continue
		} else if len(temp) == 1 {
			// 3.1. If there is only 1 digit, then it's "<digit><digit>"
			calibrators = append(calibrators, temp[0]+temp[0])
		} else {
			// 3.2. Concat first digit and last digit as "<first digit><last digit>"
			calibrators = append(calibrators, temp[0]+temp[len(temp)-1])
		}
	}

	// // 4. Sum of all calibrators
	sum := 0
	for _, calibrator := range calibrators {
		v, _ := strconv.Atoi(calibrator)
		sum += v
	}

	fmt.Println("sum of calibration values: ", sum)
}
