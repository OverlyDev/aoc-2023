package main

import (
	"fmt"

	"github.com/OverlyDev/aoc-2023/day1/files"
	"github.com/OverlyDev/aoc-2023/day1/numbers"
)

func partOne(lines []string) {
	total := 0
	for _, x := range lines {
		total += numbers.GetDigitsFromString(x)
	}
	fmt.Println("Part One Total:", total)
}

func partTwo(lines []string) {
	total := 0
	for _, x := range lines {
		total += numbers.GetFancyDigitsFromString(x)
	}
	fmt.Println("Part Two Total:", total)
}

func main() {
	lines := files.GetLinesFromFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
