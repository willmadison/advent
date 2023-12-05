package advent2023

import (
	"bufio"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

type categoryMapKey struct {
	source, destination string
}

type categoryMap struct {
	key                        categoryMapKey
	srcStart, destStart, width uint64
}

func (c *categoryMap) intersect(other categoryMap) categoryMap {
	return categoryMap{}
}

type SeedStrategy int

type seedRange struct {
	start, end uint64
}

func (s seedRange) inRange(seed uint64) bool {
	return seed >= s.start && seed <= s.end
}

const (
	Direct SeedStrategy = iota
	Pairwise
)

func FindLowestLocation(r io.Reader, seedStrategies ...SeedStrategy) uint64 {
	scanner := bufio.NewScanner(r)

	var seeds []uint64
	var seedRanges []seedRange

	currentKey := categoryMapKey{}
	categoryMapsByKey := map[categoryMapKey][]categoryMap{}

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.TrimSpace(line) == "":
			continue
		case strings.HasPrefix(line, "seeds: "):
			line = strings.TrimPrefix(line, "seeds: ")
			allSeedValues := parseBigNumbers(line)

			var strategy SeedStrategy

			if len(seedStrategies) == 0 {
				strategy = Direct
			} else {
				strategy = seedStrategies[0]
			}

			switch strategy {
			case Pairwise:
				rangeStartIndex, lengthIndex := 0, 1

				for rangeStartIndex <= len(allSeedValues)-2 {
					start, length := allSeedValues[rangeStartIndex], allSeedValues[lengthIndex]

					seedRanges = append(seedRanges, seedRange{start: start, end: start + length - 1})

					rangeStartIndex, lengthIndex = rangeStartIndex+2, lengthIndex+2
				}
			default:
				seeds = allSeedValues
			}

		case strings.HasSuffix(line, " map:"):
			line = strings.TrimSuffix(line, " map:")
			mapKeyParts := strings.Split(line, "-to-")
			currentKey = categoryMapKey{source: mapKeyParts[0], destination: mapKeyParts[1]}
		default:
			numbers := parseBigNumbers(line)
			categoryMapsByKey[currentKey] = append(categoryMapsByKey[currentKey],
				categoryMap{
					key:       currentKey,
					destStart: numbers[0],
					srcStart:  numbers[1],
					width:     numbers[2],
				})
		}
	}

	for _, categoryMaps := range categoryMapsByKey {
		sort.Slice(categoryMaps, func(i, j int) bool {
			return categoryMaps[i].srcStart < categoryMaps[j].srcStart
		})
	}

	var minLocation uint64

	minLocation = math.MaxInt64

	if len(seeds) > 0 {
		for _, seed := range seeds {
			location := determineLocation(seed, categoryMapsByKey)

			if location < minLocation {
				minLocation = location
			}
		}
	} else {
		minLocation = determineMinLocationGivenRanges(seedRanges, categoryMapsByKey)
	}

	return minLocation
}

func determineLocation(seed uint64, categoryMapsByKey map[categoryMapKey][]categoryMap) uint64 {
	mapKeyProgression := []categoryMapKey{
		{source: "seed", destination: "soil"},
		{source: "soil", destination: "fertilizer"},
		{source: "fertilizer", destination: "water"},
		{source: "water", destination: "light"},
		{source: "light", destination: "temperature"},
		{source: "temperature", destination: "humidity"},
		{source: "humidity", destination: "location"},
	}

	var src, destination uint64

	src, destination = 0, seed

	for _, key := range mapKeyProgression {
		src, destination = destination, 0

		categoryMaps := categoryMapsByKey[key]

		var mappingFound bool
		var categoryMap categoryMap

		for _, c := range categoryMaps {
			if c.srcStart <= src && src < c.srcStart+c.width {
				categoryMap = c
				mappingFound = true
				break
			}
		}

		if !mappingFound {
			destination = src
		} else {
			offset := src - categoryMap.srcStart
			destination = categoryMap.destStart + offset
		}
	}

	return destination
}

func determineMinLocationGivenRanges(seedRanges []seedRange, categoryMapsByKey map[categoryMapKey][]categoryMap) uint64 {
	for location := 0; location < math.MaxInt; location++ {
		seed, err := deriveSeedFromLocation(uint64(location), categoryMapsByKey)

		if err == nil {
			for _, sr := range seedRanges {
				if sr.inRange(seed) {
					return uint64(location)
				}
			}
		}
	}

	return math.MaxInt64
}

func deriveSeedFromLocation(location uint64, categoryMapsByKey map[categoryMapKey][]categoryMap) (uint64, error) {
	mapKeyProgression := []categoryMapKey{
		{source: "humidity", destination: "location"},
		{source: "temperature", destination: "humidity"},
		{source: "light", destination: "temperature"},
		{source: "water", destination: "light"},
		{source: "fertilizer", destination: "water"},
		{source: "soil", destination: "fertilizer"},
		{source: "seed", destination: "soil"},
	}

	var src, destination uint64

	src, destination = 0, location

	for _, key := range mapKeyProgression {
		categoryMaps := categoryMapsByKey[key]

		var mappingFound bool
		var categoryMap categoryMap

		for _, c := range categoryMaps {
			if c.destStart <= destination && destination < c.destStart+c.width {
				categoryMap = c
				mappingFound = true
				break
			}
		}

		if !mappingFound {
			src = destination
		} else {
			offset := destination - categoryMap.destStart
			src = categoryMap.srcStart + offset
		}

		src, destination = destination, src
	}

	return destination, nil
}

func parseBigNumbers(spaceDelimitedNumbers string) []uint64 {
	numbers := []uint64{}

	for _, rawNumber := range strings.Fields(spaceDelimitedNumbers) {
		n, _ := strconv.ParseUint(rawNumber, 10, 64)
		numbers = append(numbers, n)
	}

	return numbers
}
