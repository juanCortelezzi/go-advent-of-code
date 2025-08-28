package daySeven

import (
	"log"
	"strconv"
	"strings"
)

type operator = int

const (
	add operator = iota
	mul
	con
)

type Calculation struct {
	target  int
	numbers []int
}

func parseInput(input string) []Calculation {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	calculations := make([]Calculation, 0, len(lines))
	for _, line := range lines {
		targetString, rest, found := strings.Cut(line, ": ")
		if !found {
			log.Panicf("could not separate target from numbers in string \"%s\"", line)
		}

		target, err := strconv.Atoi(targetString)
		if err != nil {
			log.Panicf("target '%s' can not be converted to integer\n", targetString)
		}

		numberStrings := strings.Split(rest, " ")
		numbers := make([]int, 0, len(numberStrings))
		for _, numberString := range numberStrings {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				log.Panicf("number '%s' can not be converted to integer\n", numberString)
			}

			numbers = append(numbers, number)
		}

		calculations = append(calculations, Calculation{target: target, numbers: numbers})
	}

	return calculations
}

func cartesianProduct(alphabet []operator, n int) [][]operator {
	k := len(alphabet)
	total := 1 // n ^ k
	for range n {
		total *= k
	}

	out := make([][]operator, 0, total)
	for x := range total {
		y := x
		buf := make([]operator, n)
		for pos := range n {
			buf[pos] = alphabet[y%k]
			y /= k
		}
		out = append(out, buf)
	}
	return out
}

func PartOne(input string) int {
	calculations := parseInput(input)

	result := 0

	for _, calculation := range calculations {
		numberCount := len(calculation.numbers)
		if numberCount <= 1 {
			log.Panicf("there is less than 2 numbers in calculation: %#v\n", calculation)
		}

		variations := cartesianProduct([]operator{add, mul}, numberCount-1)
		for _, variation := range variations {
			variationResult := calculation.numbers[0]
			for i := 1; i < numberCount; i++ {
				num := calculation.numbers[i]
				op := variation[i-1]

				switch op {
				case add:
					variationResult += num
				case mul:
					variationResult *= num
				}
			}

			if variationResult == calculation.target {
				result += variationResult
				break
			}
		}
	}
	return result
}

func joinInts(a, b int) int {
	padding := 1
	for padding <= b {
		padding *= 10
	}

	return a*padding + b
}

func PartTwo(input string) int {
	calculations := parseInput(input)

	result := 0

	for _, calculation := range calculations {
		numberCount := len(calculation.numbers)
		if numberCount <= 1 {
			log.Panicf("there is less than 2 numbers in calculation: %#v\n", calculation)
		}

		variations := cartesianProduct([]operator{add, mul, con}, numberCount-1)
		for _, variation := range variations {
			variationResult := calculation.numbers[0]
			for i := 1; i < numberCount; i++ {
				num := calculation.numbers[i]
				op := variation[i-1]

				switch op {
				case add:
					variationResult += num
				case mul:
					variationResult *= num
				case con:
					variationResult = joinInts(variationResult, num)
				}
			}

			if variationResult == calculation.target {
				result += variationResult
				break
			}
		}
	}

	return result
}
