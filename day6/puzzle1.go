package day6

import (
	"log"
	"regexp"
	"strconv"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

func (p *Puzzle) Puzzle1() error {
	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	num0, err := extractNumbers(lines[0])
	if err != nil {
		return err
	}
	num1, err := extractNumbers(lines[1])
	if err != nil {
		return err
	}
	num2, err := extractNumbers(lines[2])
	if err != nil {
		return err
	}
	num3, err := extractNumbers(lines[3])
	if err != nil {
		return err
	}
	ops := extractOperators(lines[4])

	total := int64(0)

	for i := 0; i < len(num0); i++ {
		if ops[i] == "+" {
			total += num0[i] + num1[i] + num2[i] + num3[i]
		} else {
			total += num0[i] * num1[i] * num2[i] * num3[i]
		}
	}

	log.Printf("Answer for puzzle 1 is %d", total)

	return nil
}

func extractNumbers(line string) ([]int64, error) {
	pattern := regexp.MustCompile(`\d+`)
	matches := pattern.FindAllString(line, -1)

	nums := make([]int64, len(matches))

	var err error
	for i := 0; i < len(matches); i++ {
		nums[i], err = strconv.ParseInt(matches[i], 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return nums, nil
}

func extractOperators(line string) []string {
	pattern := regexp.MustCompile(`[+|*]`)
	matches := pattern.FindAllString(line, -1)

	return matches
}
