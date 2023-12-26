package solvers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
	"time"
)

func solveDay9(input io.ReadCloser) (string, string, error) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	start := time.Now()

	part1, err := day9part1(tee)
	if err != nil {
		return "", "", err
	}

	fmt.Printf("Part 1 took: %s\n", time.Since(start))
	start2 := time.Now()

	part2, err := day9part2(&buf)
	if err != nil {
		return "", "", err
	}

	fmt.Printf("Part 2 took: %s\n", time.Since(start2))
	fmt.Printf("Full solution took: %s\n", time.Since(start))

	return part1, part2, nil
}

func day9part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	histories := make([]Day9History, 0)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		nums := strings.Fields(line)
		numsInt := make([]int, 0, len(nums))

		for _, num := range nums {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				return "", err
			}
			numsInt = append(numsInt, numInt)
		}

		histories = append(histories, Day9History(numsInt))
	}

	sum := 0

	for _, h := range histories {
		ext := h.extrapolate()
		sum += ext[len(ext)-1]
	}

	return strconv.Itoa(sum), nil
}

type Day9History []int

func (history Day9History) differentiate() Day9History {
	slice := []int(history)

	resultSlice := make([]int, 0, len(slice)-1)

	for i := 0; i < len(slice)-1; i++ {
		resultSlice = append(resultSlice, slice[i+1]-slice[i])
	}

	return Day9History(resultSlice)
}

func (history Day9History) isZeroed() bool {
	slice := []int(history)

	retval := true

	for _, v := range slice {
		if v != 0 {
			retval = false
			break
		}
	}

	return retval
}

func (history Day9History) extrapolate() Day9History {
	if history.isZeroed() {
		return history.appendZero()
	}

	ext := history.differentiate().extrapolate()

	slice := []int(history)

	lastAdd := ext[len(ext)-1]

	slice = append(slice, slice[len(slice)-1]+lastAdd)

	return Day9History(slice)
}

func (history Day9History) appendZero() Day9History {
	slice := []int(history)

	slice = append(slice, 0)

	return Day9History(slice)
}

func day9part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	histories := make([]Day9History, 0)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		nums := strings.Fields(line)

		slices.Reverse(nums)

		numsInt := make([]int, 0, len(nums))

		for _, num := range nums {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				return "", err
			}
			numsInt = append(numsInt, numInt)
		}

		histories = append(histories, Day9History(numsInt))
	}

	sum := 0

	for _, h := range histories {
		ext := h.extrapolate()
		sum += ext[len(ext)-1]
	}

	return strconv.Itoa(sum), nil
}
