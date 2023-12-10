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

	type Interval struct {
		Start int
		End   int
	}
	seedStr := strings.Fields(strings.Split(lines[0], ":")[1])
	seedInterval := []Interval{}
	for i := 0; i < len(seedStr); i += 2 {
		start, _ := strconv.Atoi(seedStr[i])
		ranges, _ := strconv.Atoi(seedStr[i+1])
		end := start + ranges - 1
		seedInterval = append(seedInterval, Interval{start, end})
	}

	for _, line := range lines[1:] {
		if len(line) == 0 || line[0] < '0' || line[0] > '9' {
			continue
		}

		srcDesMapper := strings.Fields(line)
		ranges, _ := strconv.Atoi(srcDesMapper[2])
		srcMin, _ := strconv.Atoi(srcDesMapper[1])
		srcMax := srcMin + ranges - 1
		// desMin, _ := strconv.Atoi(srcDesMapper[0])
		// diff := desMin - srcMin
		srcInterval := Interval{srcMin, srcMax}
		expandedIntervals := make([][3]Interval, 0)
		for _, current := range seedInterval {
			overlap := (current.Start <= srcInterval.End) && (srcInterval.Start <= current.End)
			if overlap {
				fmt.Println("overlapp")
				intervalNums := []int{
					current.Start,
					current.End,
					srcInterval.Start,
					srcInterval.End,
				}
				slices.Sort(intervalNums)

				// if perfect overlapping
				if intervalNums[0] == intervalNums[1] && intervalNums[2] == intervalNums[3] {
					newIntervals := [3]Interval{
						{intervalNums[1], intervalNums[2]},
						{intervalNums[1], intervalNums[2]},
						{intervalNums[1], intervalNums[2]},
					}
					expandedIntervals = append(expandedIntervals, newIntervals)
				} else {
					newIntervals := [3]Interval{
						{intervalNums[0], intervalNums[1] - 1},
						{intervalNums[1], intervalNums[2]},
						{intervalNums[2] + 1, intervalNums[3]},
					}
					expandedIntervals = append(expandedIntervals, newIntervals)
				}
			}
		}
		fmt.Println(expandedIntervals)
		fmt.Println()
	}
	return 0
}

/*func (d *Day5) PartTwo(filename string, reader utils.AocReader) int {
	lines, err := reader(d.Dnum, filename)
	if err != nil {
		log.Fatal(err)
	}
	seeds := func() []int {
		seedStr := strings.Fields(strings.Split(lines[0], ":")[1])
		var seeds []int
		for i := 0; i < len(seedStr); i += 2 {
			start, _ := strconv.Atoi(seedStr[i])
			ranges, _ := strconv.Atoi(seedStr[i+1])
			end := start + ranges - 1

			fmt.Printf("(%v,%v)\n", start, end)

		}
		return seeds
	}()

	return 0
	results := make([]int, len(seeds))
	copy(results, seeds)
	type Index int
	done := make(map[Index]struct{})
	fmt.Println("before: ", results)
	for _, line := range lines[1:] {
		if len(line) == 0 {
			clear(done)
			continue
		}
		if line[0] < '0' || line[0] > '9' {
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
	}
	fmt.Println("after: ", results)
	return slices.Min(results)
}
*/
