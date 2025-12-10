package day9

import "github.com/eaglesfrogs/adventofcode2025/util"

type Puzzle struct {
	InputFile string
}

func NewPuzzle(inputFile string) util.DailyPuzzle {
	return &Puzzle{InputFile: inputFile}
}

type Point struct {
	x int64
	y int64
}

type PointPairs struct {
	p1 *Point
	p2 *Point
}
