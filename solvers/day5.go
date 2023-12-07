package solvers

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

func solveDay5(input io.ReadCloser) (string, string, error) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	part1, err := day5part1(tee)
	if err != nil {
		return "", "", err
	}

	part2, err := day5part2(&buf)
	if err != nil {
		return "", "", err
	}

	return part1, part2, nil
}

func day5part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	sum := 0

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

	}

	return strconv.Itoa(sum), nil
}

func day5part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	sum := 0

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

	}

	return strconv.Itoa(sum), nil
}
