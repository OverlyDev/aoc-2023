package game

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Regex expressions for finding things
var (
	GameNumberRe = regexp.MustCompile(`(?i)game\s(\d+)`)
	RedRe        = regexp.MustCompile(`(?i)(\d+)\sred`)
	GreenRe      = regexp.MustCompile(`(?i)(\d+)\sgreen`)
	BlueRe       = regexp.MustCompile(`(?i)(\d+)\sblue`)
)

// Holds the contents of the bag for game validation
type BagContents struct {
	Red   int
	Green int
	Blue  int
}

// Returns bag colors multiplied together
func (b BagContents) GetPower() int {
	return b.Red * b.Blue * b.Green
}

// Holds a single "handful result"
type Hand struct {
	Red   int
	Green int
	Blue  int
}

// Holds all handfuls for a given game as well as the game id
type Game struct {
	Id    int
	Valid bool
	Hands []Hand
}

// Calculates what the minimum bag could be for a given game
func (g *Game) GetMinBag() BagContents {
	bag := BagContents{}
	for _, x := range g.Hands {

		if x.Red > bag.Red {
			bag.Red = x.Red
		}

		if x.Green > bag.Green {
			bag.Green = x.Green
		}

		if x.Blue > bag.Blue {
			bag.Blue = x.Blue
		}
	}
	return bag
}

// Creates slice of Game from given slice of game strings
func CreateMultipleGames(s []string, b BagContents) []Game {
	games := make([]Game, len(s))
	for i, x := range s {
		games[i] = CreateGame(x, b)
	}
	return games
}

// Create a single Game from the given game string
func CreateGame(s string, b BagContents) Game {
	g := Game{
		Id:    GetGameIdFromString(s),
		Hands: GetHandsFromString(s),
		Valid: true,
	}

	// Mark game as invalid if any hand is invalid
	for _, x := range g.Hands {
		if !ValidHand(x, b) {
			g.Valid = false
		}
	}
	return g
}

// Extracts the game id from a given game string
func GetGameIdFromString(s string) int {
	x := GameNumberRe.FindStringSubmatch(s)
	if len(x) == 0 {
		return 0
	}
	val, err := strconv.Atoi(x[1])
	if err != nil {
		fmt.Println(err, s)
		return 0

	}
	return val
}

// Extracts the red/green/blue values from each round of a given game string
func GetHandsFromString(s string) []Hand {
	rounds := strings.Split(strings.Split(s, ": ")[1], "; ")
	hands := make([]Hand, len(rounds))

	for i, x := range rounds {
		hand := Hand{}

		r := RedRe.FindStringSubmatch(x)
		if len(r) >= 1 {
			val, _ := strconv.Atoi(r[1])
			hand.Red = val
		}

		g := GreenRe.FindStringSubmatch(x)
		if len(g) >= 1 {
			val, _ := strconv.Atoi(g[1])
			hand.Green = val
		}

		b := BlueRe.FindStringSubmatch(x)
		if len(b) >= 1 {
			val, _ := strconv.Atoi(b[1])
			hand.Blue = val
		}

		hands[i] = hand
	}
	return hands
}

// Checks if a hand is valid by ensuring none of the colors are greater than the bag contents
func ValidHand(hand Hand, bag BagContents) bool {
	if hand.Red > bag.Red {
		return false
	}

	if hand.Green > bag.Green {
		return false
	}

	if hand.Blue > bag.Blue {
		return false
	}
	return true
}

// Adds up the game ids for each valid game in given Game slice
func SumGameIds(games []Game) int {
	total := 0
	for _, x := range games {
		if x.Valid {
			total += x.Id
		}
	}
	return total
}

// Get the sum of all games' minimum bag
func GetPowerSum(games []Game) int {
	total := 0
	for _, x := range games {
		total += x.GetMinBag().GetPower()
	}
	return total
}
