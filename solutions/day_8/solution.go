package dayEight

import (
	"fmt"
	"log"
	"strings"
)

type vec2 struct {
	row int
	col int
}

func (v vec2) toString() string {
	return fmt.Sprintf("%d-%d", v.row, v.col)
}

func (v vec2) isOutsideBounds(minRow, minCol, maxRow, maxCol int) bool {
	isOutsideRowBounds := v.row < minRow || v.row > maxRow
	isOutsideColBounds := v.col < minCol || v.col > maxCol
	return isOutsideRowBounds || isOutsideColBounds
}

func getPositionVector(a, b vec2) vec2 {
	return vec2{row: b.row - a.row, col: b.col - a.col}
}

type city struct {
	rowLen                  int
	colLen                  int
	antennaGroupedPositions map[rune][]vec2
}

func isValidAntennaCharacter(char rune) bool {
	isLowercase := char >= 'a' && char <= 'z'
	isUppercase := char >= 'A' && char <= 'Z'
	isDigit := char >= '0' && char <= '9'
	return isLowercase || isUppercase || isDigit
}

func parseInput(input string) city {
	antennaGroupedPositions := make(map[rune][]vec2)
	rows := strings.Split(strings.TrimSpace(input), "\n")
	for rowIndex, row := range rows {
		for colIndex, freqChar := range row {
			if !isValidAntennaCharacter(freqChar) {
				continue
			}

			antennasOfSameFrequencies := antennaGroupedPositions[freqChar]
			pos := vec2{row: rowIndex, col: colIndex}
			antennaGroupedPositions[freqChar] = append(antennasOfSameFrequencies, pos)
		}
	}

	return city{
		rowLen:                  len(rows),
		colLen:                  len(rows[0]),
		antennaGroupedPositions: antennaGroupedPositions,
	}
}

func PartOne(input string) int {
	city := parseInput(input)
	antinodes := make(map[string]struct{})
	log.Printf("%v\n", city.antennaGroupedPositions)

	for freq, antennas := range city.antennaGroupedPositions {
		for antennaAIndex, antennaA := range antennas {
			for antennaBIndex, antennaB := range antennas {
				if antennaAIndex == antennaBIndex {
					continue
				}

				vecAToB := getPositionVector(antennaA, antennaB)
				antinode := vec2{row: antennaB.row + vecAToB.row, col: antennaB.col + vecAToB.col}
				vecBToAntinode := getPositionVector(antennaB, antinode)
				if vecAToB != vecBToAntinode {
					log.Panicf("vectors do not match: a(%v) b(%v) vecAToB(%v) vecBToAntinode(%v)\n", antennaA, antennaB, vecAToB, vecBToAntinode)
				}

				if antinode.isOutsideBounds(0, 0, city.rowLen, city.colLen) {
					continue
				}

				log.Printf("(%c) %v - %v = %v -> %v\n", freq, antennaA, antennaB, vecAToB, antinode)
				antinodes[antinode.toString()] = struct{}{}
			}
		}
	}

	return len(antinodes)
}

func PartTwo(input string) int {
	return 0
}
