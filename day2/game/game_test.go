package game_test

import (
	"testing"

	"github.com/OverlyDev/aoc-2023/day2/game"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var testGames = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

type GameIdSuite struct {
	suite.Suite
	GameStrings     []string
	ExpectedGameIds []int
}

func (suite *GameIdSuite) SetupTest() {
	suite.GameStrings = testGames
	suite.ExpectedGameIds = []int{1, 2, 3, 4, 5}
}

// Test getting game id from valid game strings
func (suite *GameIdSuite) TestGetGameIdGood() {
	for i := range suite.GameStrings {
		assert.Equal(suite.T(), suite.ExpectedGameIds[i], game.GetGameIdFromString(suite.GameStrings[i]))
	}
}

// Test func returns 0 on invalid game id
func (suite *GameIdSuite) TestGetGameIdBad() {
	x := []string{
		"AGasdm 234: saf sadf sa f",
		"Gamer69",
		"",
	}
	for i := range x {
		assert.Equal(suite.T(), 0, game.GetGameIdFromString(x[i]))
	}
}

type HandSuite struct {
	suite.Suite
	GameStrings   []string
	ExpectedHands [][]game.Hand
}

func (suite *HandSuite) SetupTest() {
	suite.GameStrings = testGames
	suite.ExpectedHands = [][]game.Hand{
		{
			{Red: 4, Green: 0, Blue: 3},
			{Red: 1, Green: 2, Blue: 6},
			{Red: 0, Green: 2, Blue: 0},
		},
		{
			{Red: 0, Green: 2, Blue: 1},
			{Red: 1, Green: 3, Blue: 4},
			{Red: 0, Green: 1, Blue: 1},
		},
		{
			{Red: 20, Green: 8, Blue: 6},
			{Red: 4, Green: 13, Blue: 5},
			{Red: 1, Green: 5, Blue: 0},
		},
		{
			{Red: 3, Green: 1, Blue: 6},
			{Red: 6, Green: 3, Blue: 0},
			{Red: 14, Green: 3, Blue: 15},
		},
		{
			{Red: 6, Green: 3, Blue: 1},
			{Red: 1, Green: 2, Blue: 2},
		},
	}
}

// Test valid red/green/blue extraction
func (suite *HandSuite) TestGetHands() {
	for i := range suite.GameStrings {
		assert.Equal(suite.T(), suite.ExpectedHands[i], game.GetHandsFromString(suite.GameStrings[i]))
	}
}

type GameValidSuite struct {
	suite.Suite
	GameStrings        []string
	Bag                game.BagContents
	ExpectedValidGames []bool
	ExpectedIdSum      int
}

func (suite *GameValidSuite) SetupTest() {
	suite.GameStrings = testGames
	suite.Bag = game.BagContents{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	suite.ExpectedValidGames = []bool{
		true,
		true,
		false,
		false,
		true,
	}
	suite.ExpectedIdSum = 8

}

// Test CreateGame and marking games invalid
func (suite *GameValidSuite) TestGameValidSingle() {
	for i, s := range suite.GameStrings {
		g := game.CreateGame(s, suite.Bag)
		assert.Equal(suite.T(), suite.ExpectedValidGames[i], g.Valid)
	}
}

// Test CreateMultipleGames and marking games invalid
func (suite *GameValidSuite) TestGameValidMulti() {
	games := game.CreateMultipleGames(suite.GameStrings, suite.Bag)
	for i, g := range games {
		assert.Equal(suite.T(), suite.ExpectedValidGames[i], g.Valid)
	}
}

// Test summing game ids
func (suite *GameValidSuite) TestSumGameIds() {
	games := make([]game.Game, len(suite.GameStrings))
	for i, s := range suite.GameStrings {
		games[i] = game.CreateGame(s, suite.Bag)
	}
	assert.Equal(suite.T(), suite.ExpectedIdSum, game.SumGameIds(games))
}

func TestGame(t *testing.T) {
	suite.Run(t, new(GameIdSuite))
	suite.Run(t, new(HandSuite))
	suite.Run(t, new(GameValidSuite))
}
