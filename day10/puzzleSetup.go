package day10

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/eaglesfrogs/adventofcode2025/util"
)

type Puzzle struct {
	InputFile string
	lights    []int
	buttons   [][]int

	joltageRequirement [][]float64
	joltageSwitches    [][][]int
}

func NewPuzzle(inputFile string) util.DailyPuzzle {
	return &Puzzle{InputFile: inputFile}
}

func (p *Puzzle) init() error {

	lines, err := util.ReadFileLines(p.InputFile)
	if err != nil {
		return err
	}

	lightsRegex := regexp.MustCompile(`\[[\.|#]+\]`)
	buttonsRegex := regexp.MustCompile(`\((\d,?)*\)`)
	joltageRegex := regexp.MustCompile(`\{(\d,?)*\}`)

	lights := make([]int, len(lines))
	buttons := make([][]int, len(lines))
	joltageRequirement := make([][]float64, len(lines))
	joltageSwitches := make([][][]int, len(lines))

	for i, line := range lines {
		lightBank := lightsRegex.FindString(line)
		buttonMatches := buttonsRegex.FindAllString(line, -1)
		joltageMatch := joltageRegex.FindString(line)

		tmpLight := 0
		for j := 1; j < len(lightBank)-1; j++ {
			c := rune(lightBank[j])
			if c == '#' {
				tmpLight = tmpLight ^ (1 << (j - 1))
			}
		}

		lights[i] = tmpLight
		buttons[i] = make([]int, 0)
		joltageSwitches[i] = make([][]int, 0)

		for _, buttonMatch := range buttonMatches {
			buttonMask := 0

			buttonMatch = strings.Replace(buttonMatch, "(", "", 1)
			buttonMatch = strings.Replace(buttonMatch, ")", "", 1)
			buttonIndexes := strings.Split(buttonMatch, ",")

			switches := make([]int, len(buttonIndexes))

			for j, buttonIdx := range buttonIndexes {
				b, err := strconv.Atoi(buttonIdx)
				if err != nil {
					return err
				}

				switches[j] = b
				buttonMask = buttonMask ^ (1 << b)
			}

			joltageSwitches[i] = append(joltageSwitches[i], switches)
			buttons[i] = append(buttons[i], buttonMask)
		}

		joltageMatch = strings.Replace(joltageMatch, "{", "", 1)
		joltageMatch = strings.Replace(joltageMatch, "}", "", 1)
		joltageValues := strings.Split(joltageMatch, ",")

		joltageReq := make([]float64, len(joltageValues))

		for j, v := range joltageValues {
			val, err := strconv.Atoi(v)
			if err != nil {
				return err
			}

			joltageReq[j] = float64(val)
		}

		joltageRequirement[i] = joltageReq
	}

	p.lights = lights
	p.buttons = buttons
	p.joltageSwitches = joltageSwitches
	p.joltageRequirement = joltageRequirement

	return nil
}
