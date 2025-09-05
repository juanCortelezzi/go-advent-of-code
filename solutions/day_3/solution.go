package day3

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func parseMul(input string) (int, int) {

	a, b, found := strings.Cut(input, ",")
	if !found {
		log.Panicf("could not find \",\" in input: \"%s\"\n", input)
	}

	first, err := strconv.Atoi(a[4:])
	if err != nil {
		log.Panicf("first item of input '%s' is not an int: '%s'\n", input, a[4:])
	}

	second, err := strconv.Atoi(b[:len(b)-1])
	if err != nil {
		log.Panicf("second item of input '%s' is not an int: '%s'\n", input, b[:len(b)-1])
	}

	return first, second
}

func PartOne(input string) int {
	matches := regexp.
		MustCompile(`mul\(\d{1,3},\d{1,3}\)`).
		FindAllString(input, -1)

	result := 0
	for _, match := range matches {
		a, b := parseMul(match)
		result += a * b
	}

	return result
}

func PartTwo(input string) int {
	matches := regexp.
		MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`).
		FindAllString(input, -1)

	result := 0
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		}

		if match == "don't()" {
			enabled = false
			continue
		}

		if enabled {
			a, b := parseMul(match)
			result += a * b
		}
	}

	return result
}
