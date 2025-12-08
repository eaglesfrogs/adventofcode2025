package day8

import (
	"log"
	"sort"
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

	pointDistances = pointDistances[0:1000]
	for _, pointDistance := range pointDistances {
		p1 := pointDistance.p1
		p2 := pointDistance.p2

		p1.connections = append(p1.connections, p2)
		p2.connections = append(p2.connections, p1)
	}

	visited := make(map[*Point]bool)
	graphLengths := make([]int, 0)

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

		graphLengths = append(graphLengths, count)
	}

	sort.Slice(graphLengths, func(i, j int) bool {
		return graphLengths[i] > graphLengths[j]
	})

	log.Printf("The answer to puzzle 1 is %d", graphLengths[0]*graphLengths[1]*graphLengths[2])

	return nil
}
