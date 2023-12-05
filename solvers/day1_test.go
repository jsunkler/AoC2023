package solvers

import (
	"io"
	"strings"
	"testing"
)

const testdata_day1_part1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const testdata_day1_part2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func Test_day1part1(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day 1 part 1", args: args{input: strings.NewReader(testdata_day1_part1)}, want: "142"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := day1part1(tt.args.input); got != tt.want {
				t.Errorf("day1part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day1part2(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day 1 part 2", args: args{input: strings.NewReader(testdata_day1_part2)}, want: "281"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := day1part2(tt.args.input); got != tt.want {
				t.Errorf("day1part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
