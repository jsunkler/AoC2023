package solvers

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func solveDay2(input io.ReadCloser) (string, string, error) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	part1, err := day2part1(tee)
	if err != nil {
		return "", "", err
	}

	part2, err := day2part2(&buf)
	if err != nil {
		return "", "", err
	}

	return part1, part2, nil
}

func day2part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	regexGameId := regexp.MustCompile(`^Game (?P<gameid>\d+): (?P<draws>.*)$`)
	regexDraw := regexp.MustCompile(`(?P<num>\d+) (?P<type>red|green|blue)`)

	games := []Day2Game{}

	for scanner.Scan() {
		line := scanner.Text()

		match := regexGameId.FindStringSubmatch(line)

		gameId, err := strconv.Atoi(match[regexGameId.SubexpIndex("gameid")])
		if err != nil {
			return "", err
		}

		currentGame := Day2Game{
			id:    gameId,
			draws: []Day2Draw{},
		}

		draws := strings.Split(match[regexGameId.SubexpIndex("draws")], ";")

		for _, v := range draws {
			currentDraw := Day2Draw{
				reds:   0,
				greens: 0,
				blues:  0,
			}

			typeDraws := regexDraw.FindAllStringSubmatch(v, -1)
			for _, v1 := range typeDraws {
				val, err := strconv.Atoi(v1[regexDraw.SubexpIndex("num")])
				if err != nil {
					return "", err
				}
				switch v1[regexDraw.SubexpIndex("type")] {
				case "red":
					currentDraw.reds = val
				case "green":
					currentDraw.greens = val
				case "blue":
					currentDraw.blues = val
				}
			}

			currentGame.draws = append(currentGame.draws, currentDraw)
		}

		games = append(games, currentGame)
	}

	var sum int

	for _, g := range games {
		if g.IsPossible() {
			sum += g.id
		}
	}

	return strconv.Itoa(sum), nil
}

func day2part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	regexGameId := regexp.MustCompile(`^Game (?P<gameid>\d+): (?P<draws>.*)$`)
	regexDraw := regexp.MustCompile(`(?P<num>\d+) (?P<type>red|green|blue)`)

	games := []Day2Game{}

	for scanner.Scan() {
		line := scanner.Text()

		match := regexGameId.FindStringSubmatch(line)

		gameId, err := strconv.Atoi(match[regexGameId.SubexpIndex("gameid")])
		if err != nil {
			return "", err
		}

		currentGame := Day2Game{
			id:    gameId,
			draws: []Day2Draw{},
		}

		draws := strings.Split(match[regexGameId.SubexpIndex("draws")], ";")

		for _, v := range draws {
			currentDraw := Day2Draw{
				reds:   0,
				greens: 0,
				blues:  0,
			}

			typeDraws := regexDraw.FindAllStringSubmatch(v, -1)
			for _, v1 := range typeDraws {
				val, err := strconv.Atoi(v1[regexDraw.SubexpIndex("num")])
				if err != nil {
					return "", err
				}
				switch v1[regexDraw.SubexpIndex("type")] {
				case "red":
					currentDraw.reds = val
				case "green":
					currentDraw.greens = val
				case "blue":
					currentDraw.blues = val
				}
			}

			currentGame.draws = append(currentGame.draws, currentDraw)
		}

		games = append(games, currentGame)
	}

	var sum int

	for _, g := range games {
		maxDraw := g.Max()
		sum += maxDraw.blues * maxDraw.greens * maxDraw.reds
	}

	return strconv.Itoa(sum), nil
}

type Day2Draw struct {
	reds   int
	greens int
	blues  int
}

type Day2Game struct {
	id    int
	draws []Day2Draw
}

func (g *Day2Game) Max() Day2Draw {
	maxDraw := Day2Draw{}

	for _, v := range g.draws {
		if v.reds > maxDraw.reds {
			maxDraw.reds = v.reds
		}
		if v.blues > maxDraw.blues {
			maxDraw.blues = v.blues
		}
		if v.greens > maxDraw.greens {
			maxDraw.greens = v.greens
		}
	}

	return maxDraw
}

func (g *Day2Game) IsPossible() bool {
	maxDraw := g.Max()

	if maxDraw.reds > 12 || maxDraw.greens > 13 || maxDraw.blues > 14 {
		return false
	}

	return true
}
