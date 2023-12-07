package solvers

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
)

func solveDay3(input io.ReadCloser) (string, string, error) {
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

func day3part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	markers := make([]Day3MarkerPoint, 0, 512)
	values := make([]Day3ValuePoint, 0, 512)

	var currentValuePoint *Day3ValuePoint = nil

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		for col, r := range line {
			if strings.Contains("0123456789", string(r)) {
				if currentValuePoint == nil {
					currentValuePoint = &Day3ValuePoint{
						row:   row,
						col:   col,
						value: string(r),
					}
				} else {
					currentValuePoint.value += string(r)
				}
			} else {
				if currentValuePoint != nil {
					values = append(values, *currentValuePoint)
					currentValuePoint = nil
				}
			}

			if !strings.Contains(".0123456789", string(r)) {
				markers = append(markers, Day3MarkerPoint{
					row: row,
					col: col,
				})
			}
		}
	}

	sum := 0

	for _, val := range values {
		if val.isActive(markers) {
			valInt, err := strconv.Atoi(val.value)
			if err != nil {
				return "", err
			}
			sum += valInt
		}
	}

	return strconv.Itoa(sum), nil
}

func day3part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	markers := make([]Day3MarkerPoint, 0, 512)
	values := make([]Day3ValuePoint, 0, 512)

	var currentValuePoint *Day3ValuePoint = nil

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		for col, r := range line {
			if strings.Contains("0123456789", string(r)) {
				if currentValuePoint == nil {
					currentValuePoint = &Day3ValuePoint{
						row:   row,
						col:   col,
						value: string(r),
					}
				} else {
					currentValuePoint.value += string(r)
				}
			} else {
				if currentValuePoint != nil {
					values = append(values, *currentValuePoint)
					currentValuePoint = nil
				}
			}

			if r == '*' {
				markers = append(markers, Day3MarkerPoint{
					row: row,
					col: col,
				})
			}
		}
	}

	sum := 0

	for _, val1 := range values {
		for _, val2 := range values {
			if val1.isEqual(&val2) {
				continue
			}

			interception := intercept(val1.getPossibleMarkerPoints(), val2.getPossibleMarkerPoints())

			if len(interception) == 0 {
				continue
			}

			gears := intercept(interception, markers)

			if len(gears) == 0 {
				continue
			}

			val1Int, err := strconv.Atoi(val1.value)
			if err != nil {
				return "", err
			}
			val2Int, err := strconv.Atoi(val2.value)
			if err != nil {
				return "", err
			}

			sum += val1Int * val2Int

		}
	}

	return strconv.Itoa(sum / 2), nil
}

type Day3MarkerPoint struct {
	row int
	col int
}

type Day3ValuePoint struct {
	row   int
	col   int
	value string
}

func intercept(points []Day3MarkerPoint, otherPoints []Day3MarkerPoint) []Day3MarkerPoint {
	resultSlice := []Day3MarkerPoint{}

	for _, iVal := range points {
		for _, jVal := range otherPoints {
			if iVal.isEqual(&jVal) {
				resultSlice = append(resultSlice, iVal)
			}
		}
	}

	return resultSlice
}

func (me *Day3MarkerPoint) isEqual(other *Day3MarkerPoint) bool {
	if me.col == other.col && me.row == other.row {
		return true
	}
	return false
}

func (me *Day3ValuePoint) isEqual(other *Day3ValuePoint) bool {
	if me.col == other.col && me.row == other.row && me.value == other.value {
		return true
	}
	return false
}

func (vp *Day3ValuePoint) getPossibleMarkerPoints() []Day3MarkerPoint {
	possibleMarkerPoints := make([]Day3MarkerPoint, 0, 16)
	for i := range vp.value {
		possibleMarkerPoints = append(possibleMarkerPoints, Day3MarkerPoint{
			col: vp.col + i,
			row: vp.row - 1,
		}, Day3MarkerPoint{
			col: vp.col + i,
			row: vp.row + 1,
		}, Day3MarkerPoint{
			col: vp.col + i - 1,
			row: vp.row,
		}, Day3MarkerPoint{
			col: vp.col + i + 1,
			row: vp.row,
		})
	}

	possibleMarkerPoints = append(possibleMarkerPoints, Day3MarkerPoint{
		col: vp.col - 1,
		row: vp.row - 1,
	}, Day3MarkerPoint{
		col: vp.col - 1,
		row: vp.row + 1,
	}, Day3MarkerPoint{
		col: vp.col + len(vp.value),
		row: vp.row - 1,
	}, Day3MarkerPoint{
		col: vp.col + len(vp.value),
		row: vp.row + 1,
	})

	return possibleMarkerPoints
}

func (vp *Day3ValuePoint) isActive(markers []Day3MarkerPoint) bool {

	possibleMarkerPoints := vp.getPossibleMarkerPoints()

	for _, m1 := range markers {
		for _, m2 := range possibleMarkerPoints {
			if m1.isEqual(&m2) {
				return true
			}
		}
	}

	return false
}
