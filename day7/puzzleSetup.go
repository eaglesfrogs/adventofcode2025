package day7

import "github.com/eaglesfrogs/adventofcode2025/util"

type Puzzle struct {
	InputFile  string
	puzzleData [][]rune
	cache      map[string]int64
}

func NewPuzzle(inputFile string) util.DailyPuzzle {
	return &Puzzle{InputFile: inputFile}
}

func (p *Puzzle) init() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	p.puzzleData = make([][]rune, len(lines))
	for i, line := range lines {
		p.puzzleData[i] = make([]rune, len(line))
		for j, char := range line {
			p.puzzleData[i][j] = char
		}
	}

	p.cache = make(map[string]int64)

	return nil
}
