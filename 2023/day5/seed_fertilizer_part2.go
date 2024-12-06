package day5

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/DanielHakim98/aoc/utils"
)

func (d *Day5) PartTwo(filename string, reader utils.AocReader) int {
	lines, err := reader(d.Dnum, filename)
	if err != nil {
		log.Fatal(err)
	}
	last := len(lines)
	mappers := make([]Mapper, 0)
	titleIdx := make(map[int]struct{})
	for i := last - 1; i >= 0; i-- {
		line := lines[i]
		// Skip if it's whitespace or title
		if len(line) == 0 || line[0] < '0' || line[0] > '9' {

			if len(line) > 0 && i != 0 {
				titleIdx[i] = struct{}{}
			}
			// fmt.Println()
			continue
		}
		mappers = append(mappers, d.GenerateMapper(line))
	}
	sort.Slice(mappers, func(i, j int) bool {
		return mappers[i].interval.Start < mappers[j].interval.End
	})

	fmt.Println(titleIdx)
	seeds := d.GenerateSeedRange(lines[0])
	for i, mapper := range mappers {
		if _, ok := titleIdx[last-1+i]; ok {
			fmt.Println("Finshing mapper group")
		}
		for _, seed := range seeds {

			d.IntervalCalculation(mapper.interval, Interval(seed), mapper.shift)
		}

	}
	// fmt.Println()
	// fmt.Println("seeds:", seeds)
	return len(mappers)
}

type SeedRange Interval

func (d *Day5) GenerateSeedRange(line string) (seedRange []SeedRange) {
	s := strings.Fields(strings.Split(line, ":")[1])
	for i := 0; i < len(s); i += 2 {
		start, _ := strconv.Atoi(s[i])
		ranges, _ := strconv.Atoi(s[i+1])
		end := start + ranges - 1
		seedRange = append(seedRange, SeedRange{start, end})
	}
	return
}

type Mapper struct {
	interval Interval // both inclusive start and end
	shift    int
}

func (d *Day5) GenerateMapper(line string) Mapper {
	nums := strings.Fields(line)
	srcMin, _ := strconv.Atoi(nums[1])
	ranges, _ := strconv.Atoi(nums[2])
	srcMax := srcMin + ranges - 1
	desMin, _ := strconv.Atoi(nums[0])
	diff := desMin - srcMin

	return Mapper{
		interval: Interval{srcMin, srcMax},
		shift:    diff,
	}
}
