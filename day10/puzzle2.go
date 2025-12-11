package day10

import (
	"fmt"
	"log"

	"gonum.org/v1/gonum/mat"
)

func (p *Puzzle) Puzzle2() error {
	p.init()

	answers := make([]int, len(p.lights))

	for i := 0; i < len(p.lights); i++ {
		matrix := mat.NewDense(len(p.joltageRequirement[i]), len(p.joltageSwitches[i]), nil)

		for j, btn := range p.joltageSwitches[i] {
			for k := 0; k < len(btn); k++ {
				matrix.Set(btn[k], j, 1)
			}
		}

		bMatrix := mat.NewVecDense(len(p.joltageRequirement[i]), p.joltageRequirement[i])

		fc := mat.Formatted(matrix, mat.Prefix("    "), mat.Squeeze())
		fmt.Printf("c = %v", fc)

		result := SolveBinary(matrix, bMatrix)

		log.Printf("%v", result)
	}

	sum := 0
	for i := 0; i < len(answers); i++ {
		sum += answers[i]
	}

	log.Printf("Answer for puzzle 1 is %d", sum)
	return nil
}

func SumVec(v *mat.VecDense) float64 {
	sum := 0.0
	for i := 0; i < v.Len(); i++ {
		sum += v.AtVec(i)
	}
	return sum
}

func SolveBinary(A *mat.Dense, b *mat.VecDense) [][]int {
	rows, cols := A.Dims()
	var results [][]int
	total := 1 << cols // 2^cols

	for mask := 0; mask < total; mask++ {
		// Build candidate vector x from the bitmask
		xVals := make([]float64, cols)
		for j := 0; j < cols; j++ {
			if (mask>>j)&1 == 1 {
				xVals[j] = 1
			}
		}
		x := mat.NewVecDense(cols, xVals)

		// Compute A*x
		var Ax mat.VecDense
		Ax.MulVec(A, x)

		// Check equality
		ok := true
		for i := 0; i < rows; i++ {
			if Ax.AtVec(i) != b.AtVec(i) {
				ok = false
				break
			}
		}

		if ok {
			sol := make([]int, cols)
			for i := range sol {
				sol[i] = int(xVals[i])
			}
			results = append(results, sol)
		}
	}

	return results
}
