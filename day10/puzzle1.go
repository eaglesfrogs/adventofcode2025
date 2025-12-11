package day10

import "log"

type Puzzle1Step struct {
	count int
	value int
}

func (p *Puzzle) Puzzle1() error {
	p.init()

	answers := make([]int, len(p.lights))

OuterLoop:
	for i := 0; i < len(p.lights); i++ {
		lights := p.lights[i]
		buttons := p.buttons[i]

		puzzleStepQueue := make([]Puzzle1Step, 0)

		for _, b := range buttons {
			currentLights := 0 ^ b

			if currentLights == lights {
				answers[i] = 1
				continue OuterLoop
			}

			puzzleStepQueue = append(puzzleStepQueue, Puzzle1Step{1, currentLights})
		}

		for len(puzzleStepQueue) > 0 {
			nextPuzzleStep := puzzleStepQueue[0]

			if len(puzzleStepQueue) > 1 {
				puzzleStepQueue = puzzleStepQueue[1:]
			} else {
				puzzleStepQueue = make([]Puzzle1Step, 0)
			}

			for _, b := range buttons {
				currentLights := nextPuzzleStep.value ^ b

				if currentLights == lights {
					answers[i] = nextPuzzleStep.count + 1
					continue OuterLoop
				}

				puzzleStepQueue = append(puzzleStepQueue, Puzzle1Step{nextPuzzleStep.count + 1, currentLights})
			}
		}
	}

	sum := 0
	for i := 0; i < len(answers); i++ {
		sum += answers[i]
	}

	log.Printf("Answer for puzzle 1 is %d", sum)

	return nil
}
