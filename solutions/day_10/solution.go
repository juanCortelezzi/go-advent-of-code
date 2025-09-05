package day10

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type vec2 struct {
	row int
	col int
}

func (v vec2) isOutsideBounds(minRow, minCol, maxRow, maxCol int) bool {
	isOutsideRowBounds := v.row < minRow || v.row > maxRow
	isOutsideColBounds := v.col < minCol || v.col > maxCol
	return isOutsideRowBounds || isOutsideColBounds
}

type inputMap struct {
	rowLen     int
	colLen     int
	trailheads []vec2
	topography [][]int
}

func parseInput(input string) inputMap {
	rawRows := strings.Split(strings.TrimSpace(input), "\n")
	trailheads := make([]vec2, 0, 10)
	topography := make([][]int, 0, len(rawRows))
	for rowIndex, rawRow := range rawRows {
		row := make([]int, 0, len(rawRow))
		for colIndex, rawCol := range rawRow {
			height, err := strconv.Atoi(string(rawCol))
			if err != nil {
				log.Panicf("day '%c' can not be converted to integer", rawCol)
			}

			if height == 0 {
				trailheads = append(trailheads, vec2{row: rowIndex, col: colIndex})
			}

			row = append(row, height)
		}
		topography = append(topography, row)
	}

	return inputMap{
		rowLen:     len(topography),
		colLen:     len(topography[0]),
		trailheads: trailheads,
		topography: topography,
	}

}

func findNextPostions(input inputMap, start vec2) []vec2 {
	directions := []vec2{
		{row: -1, col: 0},
		{row: 0, col: 1},
		{row: 1, col: 0},
		{row: 0, col: -1},
	}

	nextPositions := make([]vec2, 0, 2)

	currentHeight := input.topography[start.row][start.col]
	for _, direction := range directions {
		nextPosition := vec2{
			row: start.row + direction.row,
			col: start.col + direction.col,
		}

		if nextPosition.isOutsideBounds(0, 0, input.rowLen-1, input.colLen-1) {
			continue
		}

		nextHeight := input.topography[nextPosition.row][nextPosition.col]
		if nextHeight == currentHeight+1 {
			nextPositions = append(nextPositions, nextPosition)
		}
	}

	return nextPositions
}

func getTrailtailsFromTrailheads(input inputMap, seenTrailtails map[vec2]struct{}, trailhead vec2) {
	currentHeight := input.topography[trailhead.row][trailhead.col]
	if currentHeight == 9 {
		seenTrailtails[trailhead] = struct{}{}
	}

	for _, position := range findNextPostions(input, trailhead) {
		getTrailtailsFromTrailheads(input, seenTrailtails, position)
	}
}

func PartOne(input string) int {
	x := parseInput(input)
	log.Println(x)

	result := 0
	for _, trailhead := range x.trailheads {
		seenTrailtails := make(map[vec2]struct{})
		getTrailtailsFromTrailheads(x, seenTrailtails, trailhead)
		fmt.Println(trailhead, seenTrailtails)
		result += len(seenTrailtails)
	}

	return result
}

func getTrailheadScore(input inputMap, trailhead vec2) int {
	currentHeight := input.topography[trailhead.row][trailhead.col]
	if currentHeight == 9 {
		return 1
	}

	result := 0
	for _, position := range findNextPostions(input, trailhead) {
		result += getTrailheadScore(input, position)
	}

	return result
}

func PartTwo(input string) int {
	x := parseInput(input)
	log.Println(x)

	result := 0
	for _, trailhead := range x.trailheads {
		result += getTrailheadScore(x, trailhead)
	}

	return result
}
