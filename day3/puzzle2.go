package day3

import (
	"log"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle2() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	var totalJoltage int64 = 0

	for _, line := range lines {
		var joltage int64 = 0

		nextStartingIdx := 0
		for i := 0; i < 12; i++ {
			highest := '0'

			idx := nextStartingIdx

			for idx < len(line)-11+i {
				if line[idx] == '9' {
					highest = '9'
					nextStartingIdx = idx + 1
					break
				}

				if rune(line[idx]) > highest {
					highest = rune(line[idx])
					nextStartingIdx = idx + 1
				}

				idx++
			}

			joltage = joltage*10 + int64(highest-'0')
		}

		log.Printf("Highest voltage is %d\n", joltage)

		totalJoltage += joltage
	}

	log.Printf("Total joltage for puzzle 2 is %d", totalJoltage)

	return nil
}
