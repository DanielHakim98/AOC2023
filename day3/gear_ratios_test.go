package day3

import (
	"fmt"
	"testing"
)

type Test struct {
	filename string
	reader   func(string) ([]string, error)
	want     int
}

func TestDay3Part1(t *testing.T) {
	testCases := []Test{
		{
			filename: "example_part1",
			reader: func(string) ([]string, error) {
				return []string{
					"467..114..",
					"...*......",
					"..35..633.",
					"......#...",
					"617*......",
					".....+.58.",
					"..592.....",
					"......755.",
					"...$.*....",
					".664.598..",
				}, nil
			},
			want: 4361,
		},
	}

	challenges := Day3{}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testing '%v'", tc.filename), func(t *testing.T) {
			got := challenges.PartOne(tc.filename, tc.reader)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}

}
