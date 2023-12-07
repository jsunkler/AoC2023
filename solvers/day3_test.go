package solvers

import (
	"io"
	"strings"
	"testing"
)

const testdata_day3_part1 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

const testdata_day3_part2 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_day3part1(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day 3 part 1", args: args{input: strings.NewReader(testdata_day3_part1)}, want: "4361"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day3part1(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day3part1() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}

func Test_day3part2(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day 3 part 2", args: args{input: strings.NewReader(testdata_day3_part2)}, want: "467835"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day3part2(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day3part2() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}
