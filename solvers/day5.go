package solvers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"
)

func solveDay5(input io.ReadCloser) (string, string, error) {
	defer input.Close()

	var buf bytes.Buffer

	tee := io.TeeReader(input, &buf)

	start := time.Now()

	part1, err := day5part1(tee)
	if err != nil {
		return "", "", err
	}

	fmt.Printf("Part 1 took: %s\n", time.Since(start))
	start2 := time.Now()

	part2, err := day5part2(&buf)
	if err != nil {
		return "", "", err
	}

	fmt.Printf("Part 2 took: %s\n", time.Since(start2))
	fmt.Printf("Full solution took: %s\n", time.Since(start))

	return part1, part2, nil
}

func day5part1(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	seeds := make([]uint64, 0, 16)
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
				seedInt, err := strconv.ParseUint(seed, 10, 64)
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

		dest, err := strconv.ParseUint(destStr, 10, 64)
		if err != nil {
			return "", err
		}
		src, err := strconv.ParseUint(srcStr, 10, 64)
		if err != nil {
			return "", err
		}
		length, err := strconv.ParseUint(lenStr, 10, 64)
		if err != nil {
			return "", err
		}

		currentEditedMap.Ranges = append(currentEditedMap.Ranges, &Day5TranslateRange{
			DestinationStart: dest,
			SourceStart:      src,
			Length:           length,
		})

	}

	results := uint64(math.MaxUint64)

	for _, s := range seeds {
		results = min(results,
			humidityToLocationMap.Translate(
				temperatureToHumidityMap.Translate(
					lightToTemperatureMap.Translate(
						waterToLightMap.Translate(
							fertilizerToWaterMap.Translate(
								soilToFertilizerMap.Translate(
									seedToSoilMap.Translate(s))))))))

	}

	return strconv.FormatUint(results, 10), nil
}

func day5part2(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	seeds := make([]uint64, 0, 1024*1024*1024)
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

		prevIndex := uint64(0)
		if strings.HasPrefix(line, "seeds:") {
			splittedLine := strings.Split(line, ":")
			for index, seed := range strings.Fields(splittedLine[1]) {
				seedInt, err := strconv.ParseUint(seed, 10, 64)
				if err != nil {
					return "", err
				}

				if index%2 == 0 {
					prevIndex = seedInt
				}
				if index%2 == 1 {
					for i := prevIndex; i < prevIndex+seedInt; i++ {
						seeds = append(seeds, i)
					}
					prevIndex = 0
				}
			}

			continue
		}

		numbers := strings.Fields(line)
		destStr := numbers[0]
		srcStr := numbers[1]
		lenStr := numbers[2]

		dest, err := strconv.ParseUint(destStr, 10, 64)
		if err != nil {
			return "", err
		}
		src, err := strconv.ParseUint(srcStr, 10, 64)
		if err != nil {
			return "", err
		}
		length, err := strconv.ParseUint(lenStr, 10, 64)
		if err != nil {
			return "", err
		}

		currentEditedMap.Ranges = append(currentEditedMap.Ranges, &Day5TranslateRange{
			DestinationStart: dest,
			SourceStart:      src,
			Length:           length,
		})

	}

	results := uint64(math.MaxUint64)

	for _, s := range seeds {
		result := humidityToLocationMap.Translate(
			temperatureToHumidityMap.Translate(
				lightToTemperatureMap.Translate(
					waterToLightMap.Translate(
						fertilizerToWaterMap.Translate(
							soilToFertilizerMap.Translate(
								seedToSoilMap.Translate(s)))))))

		results = min(results, result)

	}

	return strconv.FormatUint(results, 10), nil
}

type Day5TranslateRange struct {
	SourceStart      uint64
	DestinationStart uint64
	Length           uint64
}

type Day5TranslateMap struct {
	Ranges []*Day5TranslateRange
}

func (trange *Day5TranslateRange) Translate(input uint64) (wasTranslated bool, result uint64) {
	if input >= trange.SourceStart && input < trange.SourceStart+trange.Length {
		offset := input - trange.SourceStart
		return true, trange.DestinationStart + offset
	}
	return false, input
}

func (tmap *Day5TranslateMap) Translate(input uint64) (result uint64) {
	for _, rng := range tmap.Ranges {
		wasTranslated, result := rng.Translate(input)
		if wasTranslated {
			return result
		}
	}
	return input
}
