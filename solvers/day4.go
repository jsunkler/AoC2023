package solvers

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

func solveDay4(input io.ReadCloser) (string, string, error) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	part1, err := day3part1(tee)
	if err != nil {
		return "", "", err
	}

	part2, err := day3part2(&buf)
	if err != nil {
		return "", "", err
	}

	return part1, part2, nil
}

func day4part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

	}

	sum := 0

	return strconv.Itoa(sum), nil
}

func day4part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

	}

	sum := 0

	return strconv.Itoa(sum), nil
}
