package solvers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
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

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		fmt.Println(line)
	}

	sum := 0

	return strconv.Itoa(sum), nil
}

func day7part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		fmt.Println(line)
	}

	sum := 0

	return strconv.Itoa(sum), nil
}
