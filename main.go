package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/juancortelezzi/goaoc/solutions/day_1"
	"github.com/juancortelezzi/goaoc/solutions/day_10"
	"github.com/juancortelezzi/goaoc/solutions/day_11"
	"github.com/juancortelezzi/goaoc/solutions/day_2"
	"github.com/juancortelezzi/goaoc/solutions/day_3"
	"github.com/juancortelezzi/goaoc/solutions/day_4"
	"github.com/juancortelezzi/goaoc/solutions/day_5"
	"github.com/juancortelezzi/goaoc/solutions/day_6"
	"github.com/juancortelezzi/goaoc/solutions/day_7"
	"github.com/juancortelezzi/goaoc/solutions/day_8"
	"github.com/juancortelezzi/goaoc/solutions/day_9"
)

var dayAndPartToFunc = map[int]func(input string) int{
	11:  day1.PartOne,
	12:  day1.PartTwo,
	21:  day2.PartOne,
	22:  day2.PartTwo,
	31:  day3.PartOne,
	32:  day3.PartTwo,
	41:  day4.PartOne,
	42:  day4.PartTwo,
	51:  day5.PartOne,
	52:  day5.PartTwo,
	61:  day6.PartOne,
	62:  day6.PartTwo,
	71:  day7.PartOne,
	72:  day7.PartTwo,
	81:  day8.PartOne,
	82:  day8.PartTwo,
	91:  day9.PartOne,
	92:  day9.PartTwo,
	101: day10.PartOne,
	102: day10.PartTwo,
	111: day11.PartOne,
	112: day11.PartTwo,
}

func main() {
	testFlag := flag.Bool("test", false, "use test input")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		log.Panicf("Usage: aoc [--test] <day number> <part number>\n")
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		log.Panicf("day '%s' can not be converted to integer\n", args[0])
	}

	if day < 1 || day > 24 {
		log.Panicf("day should be between 1 and 24 but got: %d\n", day)
	}

	part, err := strconv.Atoi(args[1])
	if err != nil {
		log.Panicf("part '%s' can not be converted to integer\n", args[1])
	}

	if part != 1 && part != 2 {
		log.Panicf("part should be either 1 or 2 but got: '%d'\n", part)
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
		log.Panicf("could not read input file: '%s' - %s\n", inputFilePath, err.Error())
	}

	dayAndPart := day*10 + part
	fn, found := dayAndPartToFunc[dayAndPart]
	if !found {
		log.Panicf("Could not find implementation for day %d part %d\n", day, part)
	}

	result := fn(string(input))
	log.Printf("the result for day %d part %d is: %d\n", day, part, result)
}
