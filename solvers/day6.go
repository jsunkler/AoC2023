package solvers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
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
	timesMs := make([]int, 0)
	distances := make([]int, 0)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		if row == 0 {
			strTimesMs := strings.Fields(line)[1:]
			for _, val := range strTimesMs {
				currVal, err := strconv.Atoi(val)
				if err != nil {
					return "", err
				}
				timesMs = append(timesMs, currVal)
			}
		}

		if row == 1 {
			strDistances := strings.Fields(line)[1:]
			for _, val := range strDistances {
				currVal, err := strconv.Atoi(val)
				if err != nil {
					return "", err
				}
				distances = append(distances, currVal)
			}
		}

	}

	sum := 1

	for i := range timesMs {
		currentRaceDurationMs := timesMs[i]
		currentRaceRecordDistance := distances[i]

		winnings := 0

		for holdDurationMs := 1; holdDurationMs < currentRaceDurationMs; holdDurationMs++ {
			distance := holdDurationMs * (currentRaceDurationMs - holdDurationMs)
			if distance > currentRaceRecordDistance {
				winnings++
			}
		}

		sum *= winnings
	}

	return strconv.Itoa(sum), nil
}

func day6part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)
	timesMs := make([]int, 0)
	distances := make([]int, 0)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		if row == 0 {
			strTimesMs := strings.Fields(line)[1:]
			strOneTime := strings.Join(strTimesMs, "")
			currVal, err := strconv.Atoi(strOneTime)
			if err != nil {
				return "", err
			}
			timesMs = append(timesMs, currVal)
		}

		if row == 1 {
			strDistances := strings.Fields(line)[1:]
			strOneDistance := strings.Join(strDistances, "")
			currVal, err := strconv.Atoi(strOneDistance)
			if err != nil {
				return "", err
			}
			distances = append(distances, currVal)
		}

	}

	sum := 1

	for i := range timesMs {
		currentRaceDurationMs := timesMs[i]
		currentRaceRecordDistance := distances[i]

		winnings := 0

		for holdDurationMs := 1; holdDurationMs < currentRaceDurationMs; holdDurationMs++ {
			distance := holdDurationMs * (currentRaceDurationMs - holdDurationMs)
			if distance > currentRaceRecordDistance {
				winnings++
			}
		}

		sum *= winnings
	}

	return strconv.Itoa(sum), nil
}
