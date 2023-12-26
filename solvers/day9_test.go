package solvers

import (
	"io"
	"strings"
	"testing"
)

const testdata_day9_part1 = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

const testdata_day9_part2 = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func Test_day9part1(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day9 part 1", args: args{input: strings.NewReader(testdata_day9_part1)}, want: "114"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day9part1(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day9part1() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}

func Test_day9part2(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day9 part 2", args: args{input: strings.NewReader(testdata_day9_part2)}, want: "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day9part2(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day9part2() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}
