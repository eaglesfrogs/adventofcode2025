package day9

import (
	"log"
	"strconv"
	"strings"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle1() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	points := make([]*Point, len(lines))
	for i := 0; i < len(lines); i++ {
		lineChunks := strings.Split(lines[i], ",")

		x, err := strconv.ParseInt(lineChunks[0], 10, 64)
		if err != nil {
			return err
		}
		y, err := strconv.ParseInt(lineChunks[1], 10, 64)
		if err != nil {
			return err
		}

		points[i] = &Point{x, y}
	}

	maxArea := int64(0)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			xLen := points[j].x - points[i].x
			yLen := points[j].y - points[i].y

			area := (max(xLen, -xLen) + 1) * (max(yLen, -yLen) + 1)

			if area > maxArea {
				maxArea = area
			}
		}
	}

	log.Printf("Highest area is %d", maxArea)

	return nil
}
