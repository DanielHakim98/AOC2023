package day5

import (
	"fmt"
	"testing"
)

type Test struct {
	filename string
	reader   func(int, string) ([]string, error)
	want     int
}

func TestDay5Part1(t *testing.T) {
	testCases := []Test{
		{
			filename: "example_part1",
			reader: func(int, string) ([]string, error) {
				return []string{
					"seeds: 79 14 55 13",
					"",
					"seed-to-soil map:",
					"50 98 2",
					"52 50 48",
					"soil-to-fertilizer map:",
					"0 15 37",
					"37 52 2",
					"39 0 15",
					"",
					"fertilizer-to-water map:",
					"49 53 8",
					"0 11 42",
					"42 0 7",
					"57 7 4",
					"",
					"water-to-light map:",
					"88 18 7",
					"18 25 70",
					"",
					"light-to-temperature map:",
					"45 77 23",
					"81 45 19",
					"68 64 13",
					"",
					"temperature-to-humidity map:",
					"0 69 1",
					"1 0 69",
					"",
					"humidity-to-location map:",
					"60 56 37",
					"56 93 4",
				}, nil
			},
			want: 35,
		},
		{
			// This is nasty because
			// at fertilizer-to-water map,
			// current seed 14 has soil 14, fertilizer 53,
			// and 54 will be shifted to 49 based on this mapping
			// but the catch is...49 is also matches in current map
			// so previous logic will shift 49 to 38 despite having
			// done mapping previously.
			filename: "example_single",
			reader: func(int, string) ([]string, error) {
				return []string{
					"seeds: 14",
					"",
					"seed-to-soil map:",
					"50 98 2",
					"52 50 48",
					"",
					"soil-to-fertilizer map:",
					"0 15 37",
					"37 52 2",
					"39 0 15",
					"",
					"fertilizer-to-water map:",
					"49 53 8",
					"0 11 42",
					"42 0 7",
					"57 7 4",
				}, nil
			},
			want: 49,
		},
	}

	challenges := Day5{}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testing '%v'", tc.filename), func(t *testing.T) {
			got := challenges.PartOne(tc.filename, tc.reader)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
