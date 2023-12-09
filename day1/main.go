package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/OverlyDev/aoc-2023/day1/numbers"
)

func getLinesFromFile(filename string) []string {
	// Open the file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Setup some things
	scanner := bufio.NewScanner(f)
	result := []string{}

	// Read lines to slice
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

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
	lines := getLinesFromFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
