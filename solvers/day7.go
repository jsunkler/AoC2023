package solvers

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
	"time"
)

func solveDay7(input io.ReadCloser) (string, string, error) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	start := time.Now()

	part1, err := day7part1(tee)
	if err != nil {
		return "", "", err
	}

	fmt.Printf("Part 1 took: %s\n", time.Since(start))
	start2 := time.Now()

	part2, err := day7part2(&buf)
	if err != nil {
		return "", "", err
	}

	fmt.Printf("Part 2 took: %s\n", time.Since(start2))
	fmt.Printf("Full solution took: %s\n", time.Since(start))

	return part1, part2, nil
}

func day7part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	hands := make([]*Day7Hand, 0)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		vals := strings.Fields(line)
		strHand := vals[0]
		strBid := vals[1]
		iBid, err := strconv.Atoi(strBid)
		if err != nil {
			return "", err
		}
		hands = append(hands, &Day7Hand{
			cards: []rune(strHand),
			bid:   iBid,
		})
	}

	slices.SortFunc(hands, compareHandsPart1)

	sum := 0

	for i, h := range hands {
		sum += (i + 1) * h.bid
	}

	return strconv.Itoa(sum), nil
}

func day7part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	hands := make([]*Day7Hand, 0)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		vals := strings.Fields(line)
		strHand := vals[0]
		strBid := vals[1]
		iBid, err := strconv.Atoi(strBid)
		if err != nil {
			return "", err
		}
		hands = append(hands, &Day7Hand{
			cards: []rune(strHand),
			bid:   iBid,
		})
	}

	slices.SortFunc(hands, compareHandsPart2)

	sum := 0

	for i, h := range hands {
		sum += (i + 1) * h.bid
	}

	return strconv.Itoa(sum), nil
}

var day7cardOrderingPart1 string = "23456789TJQKA"
var day7cardOrderingPart2 string = "J23456789TQKA"

type Day7Hand struct {
	cards []rune
	bid   int
}

func compareHandsPart1(a *Day7Hand, b *Day7Hand) int {
	if n := cmp.Compare(a.CalculateCompareValPart1(), b.CalculateCompareValPart1()); n != 0 {
		return n
	}

	for currIndex := 0; currIndex < 5; currIndex++ {
		rA := strings.IndexRune(day7cardOrderingPart1, a.cards[currIndex])
		rB := strings.IndexRune(day7cardOrderingPart1, b.cards[currIndex])
		if n := cmp.Compare(rA, rB); n != 0 {
			return n
		}
	}

	return 0
}

func compareHandsPart2(a *Day7Hand, b *Day7Hand) int {
	if n := cmp.Compare(a.CalculateCompareValPart2(), b.CalculateCompareValPart2()); n != 0 {
		return n
	}

	for currIndex := 0; currIndex < 5; currIndex++ {
		rA := strings.IndexRune(day7cardOrderingPart2, a.cards[currIndex])
		rB := strings.IndexRune(day7cardOrderingPart2, b.cards[currIndex])
		if n := cmp.Compare(rA, rB); n != 0 {
			return n
		}
	}

	return 0
}

func (hand *Day7Hand) CountCards() map[rune]int {
	counterMap := make(map[rune]int, 0)

	for _, r := range hand.cards {
		counterMap[r]++
	}

	return counterMap
}

func (hand *Day7Hand) CalculateCompareValPart1() int {
	m := hand.CountCards()

	values := make([]int, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	if len(m) == 1 {
		return 6
	}
	if len(m) == 2 {
		if slices.Max(values) == 4 {
			return 5
		} else {
			return 4
		}
	}
	if len(m) == 3 {
		if slices.Max(values) == 3 {
			return 3
		} else {
			return 2
		}
	}
	if len(m) == 4 {
		return 1
	}
	return 0
}

func (hand *Day7Hand) CalculateCompareValPart2() int {
	m := hand.CountCards()

	jokers := m['J']
	delete(m, 'J')

	values := make([]int, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	if len(m) <= 1 {
		return 6
	}
	if len(m) == 2 {
		if slices.Max(values)+jokers == 4 {
			return 5
		} else {
			return 4
		}
	}
	if len(m) == 3 {
		if slices.Max(values)+jokers == 3 {
			return 3
		} else {
			return 2
		}
	}
	if len(m) == 4 {
		return 1
	}
	return 0
}
