package day8

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle2() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	points := make([]*Point, len(lines))
	unconnectedPoints := make(map[*Point]bool)

	for i := 0; i < len(lines); i++ {
		lineChunks := strings.Split(lines[i], ",")

		x, err := strconv.ParseFloat(lineChunks[0], 64)
		if err != nil {
			return err
		}
		y, err := strconv.ParseFloat(lineChunks[1], 64)
		if err != nil {
			return err
		}
		z, err := strconv.ParseFloat(lineChunks[2], 64)
		if err != nil {
			return err
		}

		points[i] = &Point{x, y, z, make([]*Point, 0)}
		unconnectedPoints[points[i]] = true
	}

	pointDistances := make([]*PointDistance, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			dist := p1.calcDistance(p2)

			pointDistances = append(pointDistances, &PointDistance{p1, p2, dist})
		}
	}

	sort.Slice(pointDistances, func(i, j int) bool {
		return pointDistances[i].distance < pointDistances[j].distance
	})

	answer := 0

OuterLoop:
	for _, pointDistance := range pointDistances {
		p1 := pointDistance.p1
		p2 := pointDistance.p2

		p1.connections = append(p1.connections, p2)
		p2.connections = append(p2.connections, p1)

		if len(unconnectedPoints) > 0 {
			delete(unconnectedPoints, p1)
			delete(unconnectedPoints, p2)
		}

		if len(unconnectedPoints) == 0 {
			visited := make(map[*Point]bool)

			for _, point := range points {
				if _, v := visited[point]; v {
					continue
				}

				if len(point.connections) == 0 {
					continue
				}

				count := 1
				visited[point] = true
				pointsToCheck := make([]*Point, 0)
				pointsToCheck = append(pointsToCheck, point.connections...)

				for len(pointsToCheck) > 0 {
					nextPoint := pointsToCheck[0]
					pointsToCheck = pointsToCheck[1:]

					if _, v := visited[nextPoint]; v {
						continue
					}

					count++
					visited[nextPoint] = true

					for _, c := range nextPoint.connections {
						if _, v := visited[c]; !v {
							pointsToCheck = append(pointsToCheck, c)
						}
					}
				}

				if count == len(points) {
					answer = int(p1.x) * int(p2.x)
					break OuterLoop
				}
			}
		}
	}

	log.Printf("The answer to puzzle 2 is %d", answer)

	return nil
}
