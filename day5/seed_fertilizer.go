package day5

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

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

	seeds := func() []int {
		seedStr := strings.Fields(strings.Split(lines[0], ":")[1])
		var seeds []int
		for _, s := range seedStr {
			num, _ := strconv.Atoi(s)
			seeds = append(seeds, num)
		}
		return seeds
	}()

	results := make([]int, len(seeds))
	copy(results, seeds)
	done := make(map[int]int)
	for _, line := range lines[1:] {
		if len(line) == 0 {
			clear(done)
			// fmt.Println()
			continue
		}
		if line[0] < '0' || line[0] > '9' {
			// fmt.Println(line)
			continue
		}

		srcDesMapper := strings.Fields(line)
		ranges, _ := strconv.Atoi(srcDesMapper[2])
		srcMin, _ := strconv.Atoi(srcDesMapper[1])
		srcMax := srcMin + ranges - 1
		desMin, _ := strconv.Atoi(srcDesMapper[0])
		diff := desMin - srcMin

		// fmt.Printf("src:[%v, %v] -> target[%v, %v]; diff: %v\n", srcMin, srcMax, desMin, desMin+ranges-1, diff)
		for i, src := range results {
			if src < srcMin || src > srcMax {
				continue
			}
			if _, ok := done[src]; !ok {
				target := src + diff
				done[target] = src
				// fmt.Println("source:", results[i])
				// fmt.Println("target:", target)
				results[i] = target
			}

		}
	}

	fmt.Println(results)
	return slices.Min(results)
}

func (d *Day5) PartTwo(filename string, reader utils.AocReader) int {
	lines, err := reader(d.Dnum, filename)
	if err != nil {
		log.Fatal(err)
	}
	seeds := func() []int {
		seedStr := strings.Fields(strings.Split(lines[0], ":")[1])
		var seeds []int
		// existing := make(map[int]struct{})
		for i := 0; i < len(seedStr); i += 2 {
			current, _ := strconv.Atoi(seedStr[i])
			total, _ := strconv.Atoi(seedStr[i+1])
			for j := 1; j <= total; j++ {
				// fmt.Println("start:", current)
				// if _, ok := existing[current]; ok {
				// 	continue
				// }
				// existing[current] = struct{}{}
				seeds = append(seeds, current)
				current += 1
			}

		}
		return seeds
	}()

	results := make([]int, len(seeds))
	copy(results, seeds)
	// type STM struct {
	// 	Src int
	// 	Des int
	// }
	type Index int
	done := make(map[Index]struct{})
	fmt.Println("before: ", results)
	// fmt.Println("results[3](before):", results[3])
	for _, line := range lines[1:] {

		if len(line) == 0 {
			clear(done)
			// fmt.Println()
			continue
		}
		if line[0] < '0' || line[0] > '9' {
			// fmt.Println(line)
			continue
		}

		srcDesMapper := strings.Fields(line)
		ranges, _ := strconv.Atoi(srcDesMapper[2])
		srcMin, _ := strconv.Atoi(srcDesMapper[1])
		srcMax := srcMin + ranges - 1
		desMin, _ := strconv.Atoi(srcDesMapper[0])
		diff := desMin - srcMin

		// fmt.Printf("src:[%v, %v] -> target[%v, %sv]; diff: %v\n", srcMin, srcMax, desMin, desMin+ranges-1, diff)
		for i, src := range results {
			if src < srcMin || src > srcMax {
				continue
			}

			if _, ok := done[Index(i)]; !ok {
				target := src + diff
				done[Index(i)] = struct{}{}
				results[i] = target
			}
		}
		// fmt.Println(results[3])
	}

	// fmt.Println(results[3])
	fmt.Println("after: ", results)
	return slices.Min(results)
}
