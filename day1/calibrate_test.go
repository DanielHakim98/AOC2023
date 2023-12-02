package day1

import "testing"

func TestDay1(t *testing.T) {
	type test struct {
		filename string
		reader   func(string) ([]string, error)
		want     int
	}

	tests := []test{
		{
			filename: "example part 1",
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
	}

	for _, uc := range tests {
		got := Day1(uc.filename, uc.reader)
		if got != uc.want {
			t.Fatalf("expected: %v, got: %v", uc.want, got)
		}
	}
}
