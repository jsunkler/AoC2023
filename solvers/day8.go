package solvers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func solveDay8(input io.ReadCloser) (string, string, error) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	start := time.Now()

	part1, err := day8part1(tee)
	if err != nil {
		return "", "", err
	}

	fmt.Printf("Part 1 took: %s\n", time.Since(start))
	start2 := time.Now()

	part2, err := day8part2(&buf)
	if err != nil {
		return "", "", err
	}

	fmt.Printf("Part 2 took: %s\n", time.Since(start2))
	fmt.Printf("Full solution took: %s\n", time.Since(start))

	return part1, part2, nil
}

func day8part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var commands []rune

	sdMap := make(map[string]*Day8Destinations, 0)

	regex := regexp.MustCompile(`^(?P<source>[A-Z]{3}) = \((?P<destleft>[A-Z]{3}), (?P<destright>[A-Z]{3})\)$`)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		if row == 0 {
			commands = []rune(line)
			continue
		}

		if row == 1 {
			continue
		}

		match := regex.FindStringSubmatch(line)

		source := match[regex.SubexpIndex("source")]
		destleft := match[regex.SubexpIndex("destleft")]
		destright := match[regex.SubexpIndex("destright")]

		sdMap[source] = &Day8Destinations{
			left:  destleft,
			right: destright,
		}

	}

	sum := 0
	commandIndex := 0
	currentPos := "AAA"

	for {
		if commandIndex >= len(commands) {
			commandIndex = 0
		}

		if commands[commandIndex] == 'L' {
			currentPos = sdMap[currentPos].left
		}
		if commands[commandIndex] == 'R' {
			currentPos = sdMap[currentPos].right
		}

		commandIndex++
		sum++

		if currentPos == "ZZZ" {
			break
		}
	}

	return strconv.Itoa(sum), nil
}

func day8part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	var commands []rune

	sdMap := make(map[string]*Day8Destinations, 0)

	regex := regexp.MustCompile(`^(?P<source>[1-9A-Z]{3}) = \((?P<destleft>[1-9A-Z]{3}), (?P<destright>[1-9A-Z]{3})\)$`)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		if row == 0 {
			commands = []rune(line)
			continue
		}

		if row == 1 {
			continue
		}

		match := regex.FindStringSubmatch(line)

		source := match[regex.SubexpIndex("source")]
		destleft := match[regex.SubexpIndex("destleft")]
		destright := match[regex.SubexpIndex("destright")]

		sdMap[source] = &Day8Destinations{
			left:  destleft,
			right: destright,
		}

	}

	currentPositions := make([]string, 0)

	for k := range sdMap {
		if strings.HasSuffix(k, "A") {
			currentPositions = append(currentPositions, k)
		}
	}

	sums := make([]int, 0, len(currentPositions))

	for _, currentPos := range currentPositions {
		commandIndex := 0
		sum := 0

		for {
			if commandIndex >= len(commands) {
				commandIndex = 0
			}

			if commands[commandIndex] == 'L' {
				currentPos = sdMap[currentPos].left
			}
			if commands[commandIndex] == 'R' {
				currentPos = sdMap[currentPos].right
			}

			commandIndex++
			sum++

			if strings.HasSuffix(currentPos, "Z") {
				sums = append(sums, sum)
				break
			}
		}
	}

	val := sums[0]
	for i := 1; i < len(sums); i++ {
		val = lcm(val, sums[i])
	}

	return strconv.Itoa(val), nil
}

type Day8Destinations struct {
	left  string
	right string
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
