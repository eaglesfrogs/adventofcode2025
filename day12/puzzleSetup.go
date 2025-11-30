package day12

import "github.com/eaglesfrogs/adventofcode2025/util"

type Puzzle struct {
	InputFile string
}

func NewPuzzle(inputFile string) util.DailyPuzzle {
	return &Puzzle{InputFile: inputFile}
}
