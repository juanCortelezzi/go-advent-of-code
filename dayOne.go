package main

import (
	_ "embed"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed inputs/day_one.txt
var input string

func getListsFromInput() ([]int, []int) {

	listLeft := make([]int, 0, 100)
	listRight := make([]int, 0, 100)

	inputSplitted := strings.SplitSeq(strings.TrimSpace(input), "\n")
	for line := range inputSplitted {
		itemLeft, itemRight, found := strings.Cut(line, "   ")
		if !found {
			panic("input file is broken!!")
		}

		intLeft, err := strconv.Atoi(itemLeft)
		if err != nil {
			log.Fatalf("item '%s' can not be converted to integer", itemLeft)
		}

		intRight, err := strconv.Atoi(itemRight)
		if err != nil {
			log.Fatalf("item '%s' can not be converted to integer", itemRight)
		}

		listLeft = append(listLeft, intLeft)
		listRight = append(listRight, intRight)
	}

	if len(listLeft) != len(listRight) {
		log.Fatalf(
			"list length mismatch: len(listA) = %d and len(listB) = %d",
			len(listLeft),
			len(listRight),
		)
	}

	return listLeft, listRight
}

func DayOnePartOne() {
	listLeft, listRight := getListsFromInput()

	slices.Sort(listLeft)
	slices.Sort(listRight)

	result := 0
	for i := range len(listLeft) {
		itemLeft := listLeft[i]
		itemRight := listRight[i]

		if itemLeft > itemRight {
			result += itemLeft - itemRight
			continue
		}

		result += itemRight - itemLeft
	}

	log.Printf("result: %d", result)
}

func DayOnePartTwo() {
	listLeft, listRight := getListsFromInput()

	listRightFrequency := make(map[int]int, len(listRight))
	for _, item := range listRight {
		listRightFrequency[item] += 1
	}

	result := 0
	for _, item := range listLeft {
		result += item * listRightFrequency[item]
	}

	log.Printf("result: %d", result)
}
