package solvers

import (
	"bytes"
	"io"
)

func solveDay2(input io.ReadCloser) (string, string) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	return day2part1(tee), day2part2(&buf)
}

func day2part1(input io.Reader) string {
	return ""
}

func day2part2(input io.Reader) string {
	return ""
}
