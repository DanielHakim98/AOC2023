package day1

import (
	"fmt"
	"testing"
)

type Test struct {
	filename string
	reader   func(string) ([]string, error)
	want     int
}

func TestDay1Part2(t *testing.T) {

	testCases := []Test{
		{
			filename: "example_part1",
			reader: func(string) ([]string, error) {
				return []string{
					"1abc2",
					"pqr3stu8vwx",
					"a1b2c3d4e5f",
					"treb7uchet",
				}, nil
			},
			want: 142,
		},
		{
			filename: "fake_file",
			reader: func(s string) ([]string, error) {
				return []string{
					"two1nine",
					"eightwothree",
				}, nil
			},
			want: 112,
		},
		{
			filename: "overlap",
			reader: func(s string) ([]string, error) {
				return []string{
					"vqmoneight9tknqtcsmb",
				}, nil
			},
			want: 19,
		},
		{
			filename: "very_edgy",
			reader: func(s string) ([]string, error) {
				return []string{
					"oneeightwone",
				}, nil
			},
			want: 11,
		},
	}

	challenges := Day1{}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("testing '%v'", tc.filename), func(t *testing.T) {
			got := challenges.PartTwo(tc.filename, tc.reader)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
