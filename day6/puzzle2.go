package day6

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle2() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	ops := extractOperators(lines[4])
	numbers := make([]int64, 0)
	total := int64(0)

	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == ' ' && lines[1][i] == ' ' && lines[2][i] == ' ' && lines[3][i] == ' ' {
			ans := int64(0)
			if ops[0] == "*" {
				ans = int64(1)
			}

			for _, n := range numbers {
				if ops[0] == "+" {
					ans += n
				} else {
					ans *= n
				}
			}

			ops = ops[1:]
			total += ans
			numbers = make([]int64, 0)

			continue
		}

		numStr := strings.TrimSpace(fmt.Sprintf("%c%c%c%c", lines[0][i], lines[1][i], lines[2][i], lines[3][i]))
		num, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			return err
		}

		numbers = append(numbers, num)
	}

	ans := int64(0)
	if ops[0] == "*" {
		ans = int64(1)
	}
	for _, n := range numbers {
		if ops[0] == "+" {
			ans += n
		} else {
			ans *= n
		}
	}
	total += ans

	log.Printf("Answer for puzzle 2 is %d", total)

	return nil
}
