package files

import (
	"bufio"
	"os"
)

func GetLinesFromFile(filename string) []string {
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
