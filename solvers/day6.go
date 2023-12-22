package solvers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"time"
)

func solveDay6(input io.ReadCloser) (string, string, error) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	start := time.Now()

	part1, err := day6part1(tee)
	if err != nil {
		return "", "", err
	}

	fmt.Printf("Part 1 took: %s\n", time.Since(start))
	start2 := time.Now()

	part2, err := day6part2(&buf)
	if err != nil {
		return "", "", err
	}

	fmt.Printf("Part 2 took: %s\n", time.Since(start2))
	fmt.Printf("Full solution took: %s\n", time.Since(start))

	return part1, part2, nil
}

func day6part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

	}

	return "", nil
}

func day6part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

	}

	return "", nil
}
