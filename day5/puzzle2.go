package day5

import (
	"log"
	"strconv"
	"strings"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle2() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	line := lines[0]
	i := 0

	rangePairs := make([]*rangePair, 0)

	for line != "" {
		lineSegments := strings.Split(line, "-")
		lower, err := strconv.ParseInt(lineSegments[0], 10, 64)
		if err != nil {
			return err
		}
		upper, err := strconv.ParseInt(lineSegments[1], 10, 64)
		if err != nil {
			return err
		}

		rp := rangePair{lower, upper}
		rangePairs = append(rangePairs, &rp)

		i++
		line = lines[i]
	}

	for {
		foundOverlaps := false

		normalizedRangePairs := make([]*rangePair, 0)

		for _, rangePair := range rangePairs {
			extendedRange := false

			for _, normalizedRangePair := range normalizedRangePairs {
				if normalizedRangePair.isBetween(rangePair.lower) && rangePair.upper > normalizedRangePair.upper {
					normalizedRangePair.upper = rangePair.upper
					extendedRange = true
				}

				if normalizedRangePair.isBetween(rangePair.upper) && rangePair.lower < normalizedRangePair.lower {
					normalizedRangePair.lower = rangePair.lower
					extendedRange = true
				}

				if rangePair.lower <= normalizedRangePair.lower && rangePair.upper >= normalizedRangePair.upper {
					normalizedRangePair.lower = rangePair.lower
					normalizedRangePair.upper = rangePair.upper
					extendedRange = true
				}

				if rangePair.lower >= normalizedRangePair.lower && rangePair.upper <= normalizedRangePair.upper {
					extendedRange = true
				}
			}

			if extendedRange {
				foundOverlaps = true
			} else {
				copy := *rangePair
				normalizedRangePairs = append(normalizedRangePairs, &copy)
			}
		}

		rangePairs = normalizedRangePairs
		if !foundOverlaps {
			break
		}
	}

	total := int64(0)
	for _, rangePair := range rangePairs {
		total = total + (rangePair.upper - rangePair.lower) + 1
	}

	log.Printf("Answer to day 1 puzzle 2 is %d", total)

	return nil
}
