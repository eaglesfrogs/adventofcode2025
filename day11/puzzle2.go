package day11

import (
	"fmt"
	"log"
	"maps"
	"regexp"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func countPaths2(nodeMap map[string][]string, visited map[string]bool, node string, cache map[string]int, dac bool, fft bool) int {
	if node == "out" {
		if dac && fft {
			return 1
		}

		return 0
	}

	if visited[node] {
		return 0
	}

	key := fmt.Sprintf("%s%v%v", node, dac, fft)

	if _, ok := cache[key]; ok {
		return cache[key]
	}

	visited[node] = true

	total := 0

	if node == "dac" {
		dac = true
	}

	if node == "fft" {
		fft = true
	}

	for _, n := range nodeMap[node] {
		total += countPaths2(nodeMap, maps.Clone(visited), n, cache, dac, fft)
	}

	cache[key] = total
	return total
}

func (p *Puzzle) Puzzle2() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`[a-z]{3}`)
	nodeMap := make(map[string][]string)
	visitedMap := make(map[string]bool)
	cache := make(map[string]int)

	for _, line := range lines {
		nodes := re.FindAllString(line, -1)
		nodeMap[nodes[0]] = nodes[1:]
		visitedMap[nodes[0]] = false
	}

	total := countPaths2(nodeMap, visitedMap, "svr", cache, false, false)

	log.Printf("Answer is %d", total)

	return nil
}
