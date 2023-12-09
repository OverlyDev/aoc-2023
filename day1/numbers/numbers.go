package numbers

import (
	"regexp"
	"strconv"
	"unicode"
)

// Gets digits (1, 2, etc.) from string
func GetDigitsFromString(s string) int {
	// Slice holding found digits
	var digits []string

	// For each character in s, if it's a digit: add it to the digits slice as a string
	for i := range s {
		if unicode.IsDigit(rune(s[i])) {
			digits = append(digits, string(s[i]))
		}
	}
	// Smack first and last digit together and return it as an int
	final, _ := strconv.Atoi(string(digits[0] + string(digits[len(digits)-1])))
	return final
}

// We need to replace (word)numbers with (digit)numbers
// The tricky part is when first/last characters overlap
// That's why we just put the representative (digit)number in the middle of the (word)number
var replacements = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "th3ee",
	"four":  "fo4ur",
	"five":  "fi5ve",
	"six":   "s6x",
	"seven": "se7en",
	"eight": "ei8ht",
	"nine":  "ni9ne",
}

// Replace (word)numbers with (digit)numbers first, then GetDigitsFromString
func GetFancyDigitsFromString(s string) int {
	for word, num := range replacements {
		re := regexp.MustCompile(word)
		s = re.ReplaceAllString(s, num)
	}
	return GetDigitsFromString(s)
}
