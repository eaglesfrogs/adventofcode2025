package day7

import "log"

func (p *Puzzle) Puzzle1() error {
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

	total := p.goDownAndSplit(1, startCol)

	log.Printf("Answer to puzzle 1 is %d", total)

	return nil
}

func (p *Puzzle) goDownAndSplit(row int, col int) int {
	// we can do this here because final row has no splits in it anyway
	if row+1 == len(p.puzzleData) {
		return 0
	}

	total := 0

	if p.puzzleData[row][col] == '.' {
		p.puzzleData[row][col] = '|'
		total += p.goDownAndSplit(row+1, col)
	} else if p.puzzleData[row][col] == '^' {
		total++

		if p.puzzleData[row][col-1] != '|' {
			p.puzzleData[row][col-1] = '|'
			total += p.goDownAndSplit(row+1, col-1)
		}

		if p.puzzleData[row][col+1] != '|' {
			p.puzzleData[row][col+1] = '|'
			total += p.goDownAndSplit(row+1, col+1)
		}
	}

	return total
}
