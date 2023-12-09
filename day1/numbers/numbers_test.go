package numbers_test

import (
	"testing"

	"github.com/OverlyDev/aoc-2023/day1/numbers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PartOneSuite struct {
	suite.Suite
	NumberStrings  []string
	CorrectNumbers []int
}

func (suite *PartOneSuite) SetupTest() {
	suite.NumberStrings = []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	suite.CorrectNumbers = []int{12, 38, 15, 77}
}

func (suite *PartOneSuite) TestGetDigitsFromString() {
	for i := range suite.NumberStrings {
		assert.Equal(suite.T(), suite.CorrectNumbers[i], numbers.GetDigitsFromString(suite.NumberStrings[i]))
	}
}

type PartTwoSuite struct {
	suite.Suite
	NumberStrings  []string
	CorrectNumbers []int
}

func (suite *PartTwoSuite) SetupTest() {
	suite.NumberStrings = []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	suite.CorrectNumbers = []int{29, 83, 13, 24, 42, 14, 76}
}

func (suite *PartTwoSuite) TestGetFancyDigitsFromString() {
	for i := range suite.NumberStrings {
		assert.Equal(suite.T(), suite.CorrectNumbers[i], numbers.GetFancyDigitsFromString(suite.NumberStrings[i]))
	}
}

func TestNumbers(t *testing.T) {
	suite.Run(t, new(PartOneSuite))
	suite.Run(t, new(PartTwoSuite))
}
