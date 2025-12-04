package day4

import (
	"log"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle1() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	forkliftCount := 0

	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] != '@' {
				continue
			}

			rollCount := 0

			if row > 0 {
				if col > 0 && lines[row-1][col-1] == '@' {
					rollCount++
				}
				if lines[row-1][col] == '@' {
					rollCount++
				}
				if col < len(lines[row])-1 && lines[row-1][col+1] == '@' {
					rollCount++
				}
			}
			if row < len(lines)-1 {
				if col > 0 && lines[row+1][col-1] == '@' {
					rollCount++
				}
				if lines[row+1][col] == '@' {
					rollCount++
				}
				if col < len(lines[row])-1 && lines[row+1][col+1] == '@' {
					rollCount++
				}
			}
			if col > 0 && lines[row][col-1] == '@' {
				rollCount++
			}
			if col < len(lines[row])-1 && lines[row][col+1] == '@' {
				rollCount++
			}

			if rollCount < 4 {
				forkliftCount++
			}
		}
	}

	log.Printf("Puzzle 1 answer is %d", forkliftCount)

	return nil
}
