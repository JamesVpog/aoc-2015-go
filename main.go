package main

import (
	"2015-aoc/d1"
	"2015-aoc/d2"
	"2015-aoc/d3"
	"2015-aoc/d4"
	"2015-aoc/d5"
	"2015-aoc/d6"
	"2015-aoc/d7"
	"2015-aoc/d8"
	"fmt"
	"log"
	"os"
)

// inputReader reads the input file for a given day.
func inputReader(day string) string {
	filePath := fmt.Sprintf("d%s/%s.txt", day, day)
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading input for day %s: %v", day, err)
	}
	return string(data)
}

// solutionRunner defines a function signature for a day's solution part.
type solutionRunner func(string)

// dayRunners maps each day to its part 1 and part 2 solutions.
var dayRunners = map[string][2]solutionRunner{
	"1": {d1.P1, d1.P2},
	"2": {d2.P1, d2.P2},
	"3": {d3.P1, d3.P2},
	"4": {d4.P1, d4.P2},
	"5": {d5.P1, d5.P2},
	"6": {d6.P1, d6.P2},
	"7": {d7.P1, d7.P2},
	"8": {d8.P1, d8.P2},
	// Add new days here
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <day>")
	}
	day := os.Args[1]

	runners, ok := dayRunners[day]
	if !ok {
		log.Fatalf("No solution found for day %s", day)
	}

	content := inputReader(day)

	fmt.Printf("--- Day %s ---\n", day)
	fmt.Println("Part 1:")
	runners[0](content)
	fmt.Println("Part 2:")
	runners[1](content)
}
