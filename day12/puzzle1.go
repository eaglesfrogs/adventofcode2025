package day12

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle1() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	re1 := regexp.MustCompile(`^[0-5]:`)
	re2 := regexp.MustCompile(`^\d+x\d+:`)

	// going to count number of hashes in a present for now
	shapes := [6]int{0, 0, 0, 0, 0, 0}
	currentIdx := -1

	total := 0

	for _, line := range lines {
		if currentIdx > -1 {
			if line == "" {
				currentIdx = -1
			} else {
				for i := 0; i < len(line); i++ {
					if line[i] == '#' {
						shapes[currentIdx] += 1
					}
				}
			}

			continue
		}

		if re1.MatchString(line) {
			currentIdx, err = strconv.Atoi(line[0:1])
			if err != nil {
				return err
			}

			continue
		}

		if re2.MatchString(line) {
			parts := strings.Split(line, ": ")
			giftArea := strings.Split(parts[0], "x")

			l, err := strconv.Atoi(giftArea[0])
			if err != nil {
				return err
			}
			w, err := strconv.Atoi(giftArea[1])
			if err != nil {
				return err
			}
			area := l * w

			giftPiles := strings.Split(parts[1], " ")
			piles := [6]int{0, 0, 0, 0, 0, 0}

			for i := 0; i < len(giftPiles); i++ {
				cnt, err := strconv.Atoi(giftPiles[i])
				if err != nil {
					return err
				}

				piles[i] = cnt
			}

			cnt9x9 := 0
			for i := 0; i < len(piles); i++ {
				cnt9x9 += piles[i] * 9
			}

			if area >= cnt9x9 { // all presents can just lay side by side
				total += 1
				continue
			}

			cntShapes := 0
			for i := 0; i < len(piles); i++ {
				cntShapes += piles[i] * shapes[i]
			}

			if area > cntShapes { // if there are too many # to fit, then it wont count but this will list the ones that are harder
				log.Printf("Check %s\n", line)
			}
		}
	}

	log.Printf("Answer maybe? %d", total)

	return nil
}
