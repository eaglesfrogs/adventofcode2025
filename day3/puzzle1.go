package day3

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle1() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	totalJoltage := 0

	for _, line := range lines {
		firstHighest := '0'
		firstHighestIdx := 0

		for i, char := range line[0 : len(line)-1] {
			if char == '9' {
				firstHighest = '9'
				firstHighestIdx = i
				break
			}

			if char > firstHighest {
				firstHighest = char
				firstHighestIdx = i
			}
		}

		secondHighest := '0'
		for _, char := range line[firstHighestIdx+1 : len(line)] {
			if char == '9' {
				secondHighest = '9'
				break
			}

			if char > secondHighest {
				secondHighest = char
			}
		}

		joltage, err := strconv.Atoi(fmt.Sprintf("%c%c", firstHighest, secondHighest))
		if err != nil {
			return err
		}
		log.Printf("Highest joltage is %d\n", joltage)

		totalJoltage += joltage
	}

	log.Printf("Total joltage for puzzle 1 is %d", totalJoltage)

	return nil
}
