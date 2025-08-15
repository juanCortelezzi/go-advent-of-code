package dayFour

import (
	"fmt"
	"strings"
)

func parseInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	return strings.Split(trimmed, "\n")
}

func coordsToString(letters []string, relativeDirection [][]int, xRow, xCol int) (string, error) {
	result := ""
	for _, coord := range relativeDirection {
		row := coord[0] + xRow
		col := coord[1] + xCol

		if row < 0 || row >= len(letters) {
			return "", fmt.Errorf("Coordinate row[%d] is out of bounds: len(letters) = %d", row, len(letters))
		}

		if col < 0 || col >= len(letters[row]) {
			return "", fmt.Errorf("Coordinate col[%d] is out of bounds: len(letters[%d]) = %d", col, row, len(letters[row]))
		}

		result += string(letters[row][col])
	}

	return result, nil
}

var clockWiseRelativeDirections = [][][]int{
	{{-1, 0}, {-2, 0}, {-3, 0}},
	{{-1, 1}, {-2, 2}, {-3, 3}},
	{{0, 1}, {0, 2}, {0, 3}},
	{{1, 1}, {2, 2}, {3, 3}},
	{{1, 0}, {2, 0}, {3, 0}},
	{{1, -1}, {2, -2}, {3, -3}},
	{{0, -1}, {0, -2}, {0, -3}},
	{{-1, -1}, {-2, -2}, {-3, -3}},
}

func findXmasCountAround(letters []string, xRow, xCol int) int {
	count := 0
	for _, relativeDirection := range clockWiseRelativeDirections {
		result, err := coordsToString(letters, relativeDirection, xRow, xCol)
		if err != nil {
			continue
		}

		if result == "MAS" {
			count += 1
		}
	}

	return count
}

func PartOne(input string) int {
	letters := parseInput(input)

	count := 0
	for row := range letters {
		for col := range letters[row] {
			char := letters[row][col]
			if char == 'X' {
				count += findXmasCountAround(letters, row, col)
			}
		}
	}

	return count
}
