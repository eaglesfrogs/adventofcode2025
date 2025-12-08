package day8

import (
	"math"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

type Puzzle struct {
	InputFile string
}

func NewPuzzle(inputFile string) util.DailyPuzzle {
	return &Puzzle{InputFile: inputFile}
}

type Point struct {
	x           float64
	y           float64
	z           float64
	connections []*Point
}

func (p *Point) calcDistance(p2 *Point) float64 {
	x := p2.x - p.x
	y := p2.y - p.y
	z := p2.z - p.z
	sum := x*x + y*y + z*z

	return math.Sqrt(sum)
}

type PointDistance struct {
	p1       *Point
	p2       *Point
	distance float64
}
