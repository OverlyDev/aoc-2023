package main

import (
	"fmt"

	"github.com/OverlyDev/aoc-2023/day1/files"
	"github.com/OverlyDev/aoc-2023/day2/game"
)

func partOne(lines []string) {
	bag := game.BagContents{Red: 12, Green: 13, Blue: 14}
	games := game.CreateMultipleGames(lines, bag)
	fmt.Println("Game ID Sum:", game.SumGameIds(games))
}

func partTwo(lines []string) {
	games := game.CreateMultipleGames(lines, game.BagContents{})
	fmt.Println("Sum of Power:", game.GetPowerSum(games))
}

func main() {
	lines := files.GetLinesFromFile("input.txt")
	partOne(lines)
	partTwo(lines)
}
