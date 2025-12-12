package day11

import (
	"log"
	"maps"
	"regexp"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func countPaths1(nodeMap map[string][]string, visited map[string]bool, node string, cache map[string]int) int {
	if node == "out" {
		return 1
	}

	if visited[node] {
		return 0
	}

	if _, ok := cache[node]; ok {
		return cache[node]
	}

	visited[node] = true

	total := 0

	for _, n := range nodeMap[node] {
		total += countPaths1(nodeMap, maps.Clone(visited), n, cache)
	}

	cache[node] = total

	return total
}

func (p *Puzzle) Puzzle1() error {
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

	total := countPaths1(nodeMap, visitedMap, "you", cache)

	log.Printf("Answer is %d", total)

	return nil
}
