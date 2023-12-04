package solvers

import (
	"io"
	"strings"
	"testing"
)

const testdata_day2_part1 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

const testdata_day2_part2 = ``

func Test_day2part1(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day 2 part 1", args: args{input: strings.NewReader(testdata_day2_part1)}, want: "8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day2part1(tt.args.input); got != tt.want {
				t.Errorf("day2part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day2part2(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day 2 part 2", args: args{input: strings.NewReader(testdata_day2_part2)}, want: "-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day2part2(tt.args.input); got != tt.want {
				t.Errorf("day2part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
