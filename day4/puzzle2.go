package day4

import (
	"log"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle2() error {
	fileLines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	if fileLines[len(fileLines)-1] == "" {
		fileLines = fileLines[:len(fileLines)-1]
	}

	lines := make([][]rune, len(fileLines))
	for row := 0; row < len(fileLines); row++ {
		lines[row] = make([]rune, len(fileLines[row]))

		for col := 0; col < len(fileLines[row]); col++ {
			lines[row][col] = rune(fileLines[row][col])
		}
	}

	totalRemovalCount := 0
	removalCount := 1

	for removalCount > 0 {
		removalCount = 0

		newLines := make([][]rune, len(lines))

		for row := 0; row < len(lines); row++ {
			newLines[row] = make([]rune, len(lines[row]))

			for col := 0; col < len(lines[row]); col++ {
				if lines[row][col] != '@' {
					newLines[row][col] = '.'
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
					newLines[row][col] = '.'
					removalCount++
				} else {
					newLines[row][col] = '@'
				}
			}
		}

		totalRemovalCount += removalCount
		lines = newLines
	}

	log.Printf("Puzzle 2 answer is %d", totalRemovalCount)

	return nil
}
