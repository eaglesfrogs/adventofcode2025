package day7

import (
	"fmt"
	"log"
)

func (p *Puzzle) Puzzle2() error {
	err := p.init()
	if err != nil {
		return err
	}

	startCol := 0
	for col, r := range p.puzzleData[0] {
		if r == 'S' {
			startCol = col
			break
		}
	}

	total := p.goDownAndSplit2(1, startCol)

	log.Printf("Answer to puzzle 2 is %d", total)

	return nil
}

func (p *Puzzle) goDownAndSplit2(row int, col int) int64 {
	// we can do this here because final row has no splits in it anyway
	if row+1 == len(p.puzzleData) {
		return int64(1)
	}

	if val, ok := p.cache[fmt.Sprintf("%d-%d", row, col)]; ok {
		return val
	}

	total := int64(0)

	if p.puzzleData[row][col] == '.' {
		total += p.goDownAndSplit2(row+1, col)
	} else if p.puzzleData[row][col] == '^' {
		total += p.goDownAndSplit2(row+1, col-1)
		total += p.goDownAndSplit2(row+1, col+1)
	}

	p.cache[fmt.Sprintf("%d-%d", row, col)] = total
	return total
}
