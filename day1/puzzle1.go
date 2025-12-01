package day1

import (
	"log"
	"strconv"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle1() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	dial := 50
	zeroCount := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		dir := line[0:1]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			return err
		}

		dist = dist % 100

		if dir == "L" {
			dial = dial - dist
		} else {
			dial = dial + dist
		}

		// some of the rotations are >100
		if dial >= 100 {
			dial = dial - 100
		}

		if dial < 0 {
			// add a negtive to 100 to subtract
			dial = 100 + dial
		}

		if dial == 0 {
			zeroCount++
		}
	}

	log.Printf("Answer to puzzle 1 is %d", zeroCount)

	return nil
}
