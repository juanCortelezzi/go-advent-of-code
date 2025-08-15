package dayThree

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func PartOne(input string) int {
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := regex.FindAllString(input, -1)

	result := 0
	for _, match := range matches {
		a, b, found := strings.Cut(match, ",")
		if !found {
			log.Fatalf("could not find \",\" in match: \"%s\"\n", match)
		}

		first, err := strconv.Atoi(a[4:])
		if err != nil {
			log.Fatalf("first item of match '%s' is not an int: '%s'", match, a[4:])
		}

		second, err := strconv.Atoi(b[:len(b)-1])
		if err != nil {
			log.Fatalf("second item of match '%s' is not an int: '%s'", match, b[:len(b)-1])
		}

		result += first * second
	}

	return result
}

func PartTwo(input string) int {
	return 0
}
