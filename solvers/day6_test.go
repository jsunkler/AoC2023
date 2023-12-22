package solvers

import (
	"io"
	"strings"
	"testing"
)

const testdata_day6_part1 = `Time:      7  15   30
Distance:  9  40  200`

const testdata_day6_part2 = `Time:      7  15   30
Distance:  9  40  200`

func Test_day6part1(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day 6 part 1", args: args{input: strings.NewReader(testdata_day6_part1)}, want: "288"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day6part1(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day6part1() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}

func Test_day6part2(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day 6 part 2", args: args{input: strings.NewReader(testdata_day6_part2)}, want: "71503"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day6part2(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day6part2() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}
