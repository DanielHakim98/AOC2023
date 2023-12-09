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
