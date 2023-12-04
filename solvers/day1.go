package solvers

import (
	"bufio"
	"bytes"
	"io"
	"log/slog"
	"regexp"
	"strconv"
)

func solveDay1(input io.ReadCloser) (string, string) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	return day1part1(tee), day1part2(&buf)
}

func day1part1(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	regex := regexp.MustCompile(`\d`)

	var sum int

	for scanner.Scan() {
		line := scanner.Text()

		matches := regex.FindAllString(line, -1)

		s1 := matches[0]
		s2 := matches[len(matches)-1]

		number, err := strconv.Atoi(s1 + s2)

		if err != nil {
			slog.Error("Cannot convert number.", "s1", s1, "s2", s2)
		}

		sum += number
	}

	return strconv.Itoa(sum)
}

func day1part2(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	stringToNumberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	stringNumbers := "one|two|three|four|five|six|seven|eight|nine"
	stringNumbersReverse := reverseString(stringNumbers)

	regexForward := regexp.MustCompile(`\d|` + stringNumbers)
	regexReverse := regexp.MustCompile(`\d|` + stringNumbersReverse)

	var sum int

	for scanner.Scan() {
		line := scanner.Text()

		s1 := regexForward.FindString(line)
		s2 := regexReverse.FindString(reverseString(line))

		if len([]rune(s1)) > 1 {
			s1 = stringToNumberMap[s1]
		}

		if len([]rune(s2)) > 1 {
			s2 = stringToNumberMap[reverseString(s2)]
		}

		number, err := strconv.Atoi(s1 + s2)

		if err != nil {
			slog.Error("Cannot convert number.", "s1", s1, "s2", s2)
		}

		sum += number
	}

	return strconv.Itoa(sum)
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
