package solvers

import (
	"io"
	"strings"
	"testing"
)

const testdata_day8_part1_1 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const testdata_day8_part1_2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const testdata_day8_part2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func Test_day8part1(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day8 part 1 (1)", args: args{input: strings.NewReader(testdata_day8_part1_1)}, want: "2"},
		{name: "testingdata day8 part 1 (2)", args: args{input: strings.NewReader(testdata_day8_part1_2)}, want: "6"},
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
		{name: "testingdata day8 part 2", args: args{input: strings.NewReader(testdata_day8_part2)}, want: "6"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day8part2(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day8part2() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}
