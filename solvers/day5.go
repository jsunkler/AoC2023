package solvers

import (
	"bufio"
	"bytes"
	"io"
	"slices"
	"strconv"
	"strings"
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

	seeds := make([]int64, 0, 16)
	seedToSoilMap := Day5TranslateMap{
		Ranges: make([]*Day5TranslateRange, 0, 32),
	}
	soilToFertilizerMap := Day5TranslateMap{
		Ranges: make([]*Day5TranslateRange, 0, 32),
	}
	fertilizerToWaterMap := Day5TranslateMap{
		Ranges: make([]*Day5TranslateRange, 0, 32),
	}
	waterToLightMap := Day5TranslateMap{
		Ranges: make([]*Day5TranslateRange, 0, 32),
	}
	lightToTemperatureMap := Day5TranslateMap{
		Ranges: make([]*Day5TranslateRange, 0, 32),
	}
	temperatureToHumidityMap := Day5TranslateMap{
		Ranges: make([]*Day5TranslateRange, 0, 32),
	}
	humidityToLocationMap := Day5TranslateMap{
		Ranges: make([]*Day5TranslateRange, 0, 32),
	}

	var currentEditedMap *Day5TranslateMap

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			currentEditedMap = nil
			continue
		}

		if strings.HasPrefix(line, "seed-to-soil") {
			currentEditedMap = &seedToSoilMap
			continue
		}

		if strings.HasPrefix(line, "soil-to-fertilizer") {
			currentEditedMap = &soilToFertilizerMap
			continue
		}

		if strings.HasPrefix(line, "fertilizer-to-water") {
			currentEditedMap = &fertilizerToWaterMap
			continue
		}

		if strings.HasPrefix(line, "water-to-light") {
			currentEditedMap = &waterToLightMap
			continue
		}

		if strings.HasPrefix(line, "light-to-temperature") {
			currentEditedMap = &lightToTemperatureMap
			continue
		}

		if strings.HasPrefix(line, "temperature-to-humidity") {
			currentEditedMap = &temperatureToHumidityMap
			continue
		}

		if strings.HasPrefix(line, "humidity-to-location") {
			currentEditedMap = &humidityToLocationMap
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			splittedLine := strings.Split(line, ":")
			for _, seed := range strings.Fields(splittedLine[1]) {
				seedInt, err := strconv.ParseInt(seed, 10, 64)
				if err != nil {
					return "", err
				}
				seeds = append(seeds, seedInt)
			}

			continue
		}

		numbers := strings.Fields(line)
		destStr := numbers[0]
		srcStr := numbers[1]
		lenStr := numbers[2]

		dest, err := strconv.ParseInt(destStr, 10, 64)
		if err != nil {
			return "", err
		}
		src, err := strconv.ParseInt(srcStr, 10, 64)
		if err != nil {
			return "", err
		}
		length, err := strconv.ParseInt(lenStr, 10, 64)
		if err != nil {
			return "", err
		}

		currentEditedMap.Ranges = append(currentEditedMap.Ranges, &Day5TranslateRange{
			DestinationStart: dest,
			SourceStart:      src,
			Length:           length,
		})

	}

	results := make([]int64, 0, len(seeds))

	for _, s := range seeds {
		results = append(results,
			humidityToLocationMap.Translate(
				temperatureToHumidityMap.Translate(
					lightToTemperatureMap.Translate(
						waterToLightMap.Translate(
							fertilizerToWaterMap.Translate(
								soilToFertilizerMap.Translate(
									seedToSoilMap.Translate(s))))))))

	}

	return strconv.FormatInt((slices.Min(results)), 10), nil
}

func day5part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	sum := 0

	for row := 0; scanner.Scan(); row++ {
		//line := scanner.Text()

	}

	return strconv.Itoa(sum), nil
}

type Day5TranslateRange struct {
	SourceStart      int64
	DestinationStart int64
	Length           int64
}

type Day5TranslateMap struct {
	Ranges []*Day5TranslateRange
}

func (trange *Day5TranslateRange) Translate(input int64) (wasTranslated bool, result int64) {
	if input >= trange.SourceStart && input < trange.SourceStart+trange.Length {
		offset := input - trange.SourceStart
		return true, trange.DestinationStart + offset
	}
	return false, input
}

func (tmap *Day5TranslateMap) Translate(input int64) (result int64) {
	for _, rng := range tmap.Ranges {
		wasTranslated, result := rng.Translate(input)
		if wasTranslated {
			return result
		}
	}
	return input
}
