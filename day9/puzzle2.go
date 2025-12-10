package day9

import (
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

const MAX_SIZE = 100000

func (p *Puzzle) Puzzle2() error {
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

	grid := make([][]int, MAX_SIZE)
	for i := 0; i < MAX_SIZE; i++ {
		grid[i] = make([]int, MAX_SIZE)
	}

	for i := 0; i < len(points); i++ {
		p1 := points[i]
		var p2 *Point

		if i == len(points)-1 {
			p2 = points[0]
		} else {
			p2 = points[i+1]
		}

		if p1.x == p2.x {
			//going vertical
			for y := min(p1.y, p2.y); y <= max(p1.y, p2.y); y++ {
				grid[p1.x][y] = 1
			}
		} else {
			//going horizontal
			for x := min(p1.x, p2.x); x <= max(p1.x, p2.x); x++ {
				grid[x][p1.y] = 1
			}
		}
	}

	floodFill(grid, 0, 0)

	maxArea := int64(0)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			if p1.x == p2.x || p1.y == p2.y {
				continue
			}

			xLen := p2.x - p1.x
			yLen := p2.y - p1.y

			area := (max(xLen, -xLen) + 1) * (max(yLen, -yLen) + 1)

			otherP1 := &Point{p1.x, p2.y}
			otherP2 := &Point{p2.x, p1.y}

			if !pointInOnShape(otherP1, grid) || !pointInOnShape(otherP2, grid) {
				continue
			}

			var wg sync.WaitGroup
			results := make(chan bool, 4)

			wg.Add(4)

			go worker(p1, otherP1, grid, results, &wg)
			go worker(otherP1, p2, grid, results, &wg)
			go worker(p2, otherP2, grid, results, &wg)
			go worker(otherP2, p1, grid, results, &wg)

			go func() {
				wg.Wait()
				close(results)
			}()

			answer := true

			for res := range results {
				answer = answer && res
			}

			if answer && area > maxArea {
				maxArea = area
			}

		}
	}

	log.Printf("Highest area is %d", maxArea)

	return nil
}

func worker(p1, p2 *Point, grid [][]int, results chan<- bool, wg *sync.WaitGroup) {

	defer wg.Done()

	if p1.x == p2.x {
		//going vertical
		for y := min(p1.y, p2.y); y <= max(p1.y, p2.y); y++ {
			inShape := pointInOnShape(&Point{p1.x, y}, grid)
			if !inShape {
				results <- false
				return
			}
		}
	} else {
		//going horizontal
		for x := min(p1.x, p2.x); x <= max(p1.x, p2.x); x++ {
			inShape := pointInOnShape(&Point{x, p1.y}, grid)
			if !inShape {
				results <- false
				return
			}
		}
	}

	results <- true
}

func pointInOnShape(p *Point, grid [][]int) bool {
	if grid[p.x][p.y] == 1 {
		return true
	}

	hitEdge := false

	crosses := 0
	for x := p.x; x >= 0; x-- {
		if !hitEdge && grid[x][p.y] == 1 {
			hitEdge = true
		} else if hitEdge && grid[x][p.y] == 0 {
			crosses++
			hitEdge = false
		}
	}

	return crosses%2 == 1
}

func floodFill(grid [][]int, sx, sy int64) {
	start := &Point{sx, sy}

	stack := []*Point{start}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if p.x < 0 || p.x >= MAX_SIZE || p.y < 0 || p.y >= MAX_SIZE || grid[p.x][p.y] != 0 {
			continue
		}

		grid[p.x][p.y] = 2

		stack = append(stack, &Point{p.x + 1, p.y})
		stack = append(stack, &Point{p.x - 1, p.y})
		stack = append(stack, &Point{p.x, p.y + 1})
		stack = append(stack, &Point{p.x, p.y - 1})
	}
}
