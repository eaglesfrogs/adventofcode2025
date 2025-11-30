package main

import (
	"flag"
	"log"
	"os"

	"github.com/eaglesfrogs/adventofcode2025/day1"
	"github.com/eaglesfrogs/adventofcode2025/day10"
	"github.com/eaglesfrogs/adventofcode2025/day11"
	"github.com/eaglesfrogs/adventofcode2025/day12"
	"github.com/eaglesfrogs/adventofcode2025/day2"
	"github.com/eaglesfrogs/adventofcode2025/day3"
	"github.com/eaglesfrogs/adventofcode2025/day4"
	"github.com/eaglesfrogs/adventofcode2025/day5"
	"github.com/eaglesfrogs/adventofcode2025/day6"
	"github.com/eaglesfrogs/adventofcode2025/day7"
	"github.com/eaglesfrogs/adventofcode2025/day8"
	"github.com/eaglesfrogs/adventofcode2025/day9"
	"github.com/eaglesfrogs/adventofcode2025/util"
	"github.com/joho/godotenv"
)

func main() {
	day := flag.Int("day", 1, "Day of the Advent of Code challenge")
	flag.Parse()

	if *day < 1 || *day > 12 {
		log.Fatalf("Invalid day: %d. Please provide a day between 1 and 12.", *day)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	session := os.Getenv("SESSION_TOKEN")

	inputFile, err := util.GetInputPath(*day, session)
	if err != nil {
		log.Fatalf("Error getting input file: %v", err)
	}

	var puzzle util.DailyPuzzle
	switch *day {
	case 1:
		puzzle = day1.NewPuzzle(inputFile)
	case 2:
		puzzle = day2.NewPuzzle(inputFile)
	case 3:
		puzzle = day3.NewPuzzle(inputFile)
	case 4:
		puzzle = day4.NewPuzzle(inputFile)
	case 5:
		puzzle = day5.NewPuzzle(inputFile)
	case 6:
		puzzle = day6.NewPuzzle(inputFile)
	case 7:
		puzzle = day7.NewPuzzle(inputFile)
	case 8:
		puzzle = day8.NewPuzzle(inputFile)
	case 9:
		puzzle = day9.NewPuzzle(inputFile)
	case 10:
		puzzle = day10.NewPuzzle(inputFile)
	case 11:
		puzzle = day11.NewPuzzle(inputFile)
	case 12:
		puzzle = day12.NewPuzzle(inputFile)
	default:
		log.Fatalf("Day %d not implemented yet.", *day)
	}

	log.Printf("Executing Puzzle 1 for day %d", *day)
	err = puzzle.Puzzle1()
	if err != nil {
		log.Fatalf("Error executing Puzzle 1 for day %d: %v", *day, err)
	}
	log.Printf("Executing Puzzle 2 for day %d", *day)
	err = puzzle.Puzzle2()
	if err != nil {
		log.Fatalf("Error executing Puzzle 2 for day %d: %v", *day, err)
	}
}
