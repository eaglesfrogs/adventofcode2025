package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

var Factors = map[int][]int{
	2:  {1},
	3:  {1},
	4:  {1, 2},
	5:  {1},
	6:  {1, 2, 3},
	7:  {1},
	8:  {1, 2, 4},
	9:  {1, 3},
	10: {1, 2, 5},
	11: {1},
	12: {1, 2, 3, 4, 6},
}

func (p *Puzzle) Puzzle2() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	invalidTotal := 0
	ranges := strings.Split(lines[0], ",")

	for _, r := range ranges {
		ends := strings.Split(r, "-")
		lower, err := strconv.Atoi(ends[0])
		if err != nil {
			return err
		}

		upper, err := strconv.Atoi(ends[1])
		if err != nil {
			return err
		}

		for num := lower; num <= upper; num++ {
			if num < 10 {
				continue
			}

			numStr := strconv.Itoa(num)
			factors := Factors[len(numStr)]

			for _, factor := range factors {
				prefix := numStr[0:factor]
				repeatedStr := strings.Repeat(prefix, len(numStr)/factor)

				if repeatedStr == numStr {
					invalidTotal += num
					break
				}
			}
		}
	}

	log.Printf("Invalid total for puzzle 2 is %d\n", invalidTotal)

	return nil
}
