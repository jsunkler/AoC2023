package solvers

import (
	"bufio"
	"bytes"
	"io"
	"math"
	"strconv"
	"strings"
)

func solveDay4(input io.ReadCloser) (string, string, error) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	part1, err := day4part1(tee)
	if err != nil {
		return "", "", err
	}

	part2, err := day4part2(&buf)
	if err != nil {
		return "", "", err
	}

	return part1, part2, nil
}

func day4part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	sum := 0

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		card, err := readCard(line)
		if err != nil {
			return "", err
		}

		sum += card.getScore()
	}

	return strconv.Itoa(sum), nil
}

func day4part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	cards := make([]*Day4Card, 0, 4096)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		card, err := readCard(line)
		if err != nil {
			return "", err
		}

		cards = append(cards, card)
	}

	for _, c := range cards {
		c.countHits()
	}

	length := len(cards)

	for i, c := range cards {
		for j := 1; j <= c.Hits && i+j < length; j++ {
			cards[i+j].Instances += c.Instances
		}
	}

	sum := 0

	for _, c := range cards {
		sum += c.Instances
	}

	return strconv.Itoa(sum), nil
}

type Day4Card struct {
	CardId         int
	WinningNumbers []int
	OwnedNumbers   []int
	Hits           int
	Instances      int
}

func (card *Day4Card) countHits() {
	hits := 0

	for _, ownedInt := range card.OwnedNumbers {
		for _, winningInt := range card.WinningNumbers {
			if ownedInt == winningInt {
				hits++
				break
			}
		}
	}

	card.Hits = hits
}

func (card *Day4Card) getScore() int {
	card.countHits()

	if card.Hits == 0 {
		return 0
	}

	return MathPow(2, card.Hits-1)
}

func MathPow(n, m int) int {
	return int(math.Pow(float64(n), float64(m)))
}

func readCard(line string) (*Day4Card, error) {

	split1 := strings.Split(line, ":")
	cardIdStr := strings.Fields(split1[0])[1]
	cardIdInt, err := strconv.Atoi(cardIdStr)
	if err != nil {
		return nil, err
	}

	split2 := strings.Split(split1[1], "|")
	partWinning := split2[0]
	partOwned := split2[1]

	winningNumbers := strings.Fields(partWinning)
	ownedNumbers := strings.Fields(partOwned)

	currentCard := Day4Card{
		CardId:         cardIdInt,
		WinningNumbers: make([]int, 0, 10),
		OwnedNumbers:   make([]int, 0, 25),
		Instances:      1,
	}

	for _, val := range winningNumbers {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		currentCard.WinningNumbers = append(currentCard.WinningNumbers, valInt)
	}

	for _, val := range ownedNumbers {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		currentCard.OwnedNumbers = append(currentCard.OwnedNumbers, valInt)
	}

	return &currentCard, nil
}
