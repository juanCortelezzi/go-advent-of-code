package day1

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func getListsFromInput(input string) ([]int, []int) {

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
			log.Panicf("item '%s' can not be converted to integer\n", itemLeft)
		}

		intRight, err := strconv.Atoi(itemRight)
		if err != nil {
			log.Panicf("item '%s' can not be converted to intege\n", itemRight)
		}

		listLeft = append(listLeft, intLeft)
		listRight = append(listRight, intRight)
	}

	if len(listLeft) != len(listRight) {
		log.Panicf(
			"list length mismatch: len(listA) = %d and len(listB) = %d\n",
			len(listLeft),
			len(listRight),
		)
	}

	return listLeft, listRight
}

func PartOne(input string) int {
	listLeft, listRight := getListsFromInput(input)

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

	return result
}

func PartTwo(input string) int {
	listLeft, listRight := getListsFromInput(input)

	listRightFrequency := make(map[int]int, len(listRight))
	for _, item := range listRight {
		listRightFrequency[item] += 1
	}

	result := 0
	for _, item := range listLeft {
		result += item * listRightFrequency[item]
	}

	return result
}
