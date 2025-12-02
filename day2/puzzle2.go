package day2

import (
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
			// numStr := strconv.Itoa(num)

			// // middly digits don't count.  so get rid of odd length numbers
			// if len(numStr)%2 == 1 {
			// 	continue
			// }

			// numHalf := numStr[0 : len(numStr)/2]

			// if strings.HasSuffix(numStr, numHalf) {
			// 	log.Printf("Range %d-%d hash invalid ID %d", lower, upper, num)
			// 	invalidTotal += num
			// }
		}
	}

	log.Printf("Invalid total for puzzle 2 is %d\n", invalidTotal)

	return nil
}
