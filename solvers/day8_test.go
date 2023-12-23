package solvers

import (
	"io"
	"strings"
	"testing"
)

const testdata_day8_part1 = ``

const testdata_day8_part2 = ``

func Test_day8part1(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day8 part 1", args: args{input: strings.NewReader(testdata_day8_part1)}, want: "-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day8part1(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day8part1() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}

func Test_day8part2(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day8 part 2", args: args{input: strings.NewReader(testdata_day8_part2)}, want: "-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day8part2(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day8part2() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}
