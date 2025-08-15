package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	dayOne "github.com/juancortelezzi/goaoc/solutions/day_1"
	dayTwo "github.com/juancortelezzi/goaoc/solutions/day_2"
	dayThree "github.com/juancortelezzi/goaoc/solutions/day_3"
)

var dayAndPartToFunc = map[int]func(input string) int{
	11: dayOne.PartOne,
	12: dayOne.PartTwo,
	21: dayTwo.PartOne,
	22: dayTwo.PartTwo,
	31: dayThree.PartOne,
	32: dayThree.PartTwo,
}

func main() {
	testFlag := flag.Bool("test", false, "use test input")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		log.Fatalf("Usage: aoc [--test] <day number> <part number>\n")
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("day '%s' can not be converted to integer\n", args[0])
	}

	if day < 1 || day > 24 {
		log.Fatalf("day should be between 1 and 24 but got: %d\n", day)
	}

	part, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("part '%s' can not be converted to integer\n", args[1])
	}

	if part != 1 && part != 2 {
		log.Fatalf("part should be either 1 or 2 but got: '%d'\n", part)
	}

	var inputFileExtension string
	if *testFlag {
		inputFileExtension = "test.txt"
	} else {
		inputFileExtension = "txt"
	}

	inputFilePath := fmt.Sprintf("./inputs/day_%d.%s", day, inputFileExtension)
	input, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("could not read input file: '%s' - %s\n", inputFilePath, err.Error())
	}

	dayAndPart := day*10 + part
	fn, found := dayAndPartToFunc[dayAndPart]
	if !found {
		log.Fatalf("Could not find implementation for day %d part %d\n", day, part)
	}

	result := fn(string(input))
	log.Printf("the result for day %d part %d is: %d\n", day, part, result)
}
