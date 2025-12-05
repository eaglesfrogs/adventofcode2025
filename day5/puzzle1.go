package day5

import (
	"log"
	"strconv"
	"strings"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle1() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	line := lines[0]
	i := 0

	rangePairs := make([]rangePair, 0)

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
		rangePairs = append(rangePairs, rp)

		i++
		line = lines[i]
	}

	i++

	totalFresh := 0

	for i < len(lines) {
		if lines[i] == "" {
			continue
		}

		num, err := strconv.ParseInt(lines[i], 10, 64)
		if err != nil {
			return err
		}

		for _, rangePair := range rangePairs {
			if rangePair.isBetween(num) {
				totalFresh++
				break
			}
		}
		i++
	}

	log.Printf("Answer to day 1 puzzle 1 is %d", totalFresh)

	return nil
}
