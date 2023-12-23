package solvers

import (
	"io"
	"strings"
	"testing"
)

const testdata_day7_part1 = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

const testdata_day7_part2 = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func Test_day7part1(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day 7 part 1", args: args{input: strings.NewReader(testdata_day7_part1)}, want: "6440"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day7part1(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day7part1() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}

func Test_day7part2(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "testingdata day 7 part 2", args: args{input: strings.NewReader(testdata_day7_part2)}, want: "5905"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := day7part2(tt.args.input); got != tt.want || err != nil {
				t.Errorf("day7part2() = %v, want %v, err %v", got, tt.want, err)
			}
		})
	}
}
