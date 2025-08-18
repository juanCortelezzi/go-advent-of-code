package daySix

import (
	"fmt"
	"log"
	"strings"
)

type direction = int

const (
	UP direction = iota
	RIGTH
	DOWN
	LEFT
)

type position struct {
	row int
	col int
}

func (p position) toString() string {
	return fmt.Sprintf("%d-%d", p.row, p.col)
}

func (p position) isOutsideBounds(minRow, minCol, maxRow, maxCol int) bool {
	isOutsideRowBounds := p.row < minRow || p.row > maxRow
	isOutsideColBounds := p.col < minCol || p.col > maxCol
	return isOutsideRowBounds || isOutsideColBounds
}

type lab struct {
	rowLen    int
	colLen    int
	obstacles map[string]struct{}
	visited   map[string]struct{}
	player    position
}

func (lab lab) toString() string {
	str := ""
	for rowIndex := range lab.rowLen {
		for colIndex := range lab.colLen {
			position := position{row: rowIndex, col: colIndex}
			if position == lab.player {
				str += "^"
				continue
			}

			if _, isObstacle := lab.obstacles[position.toString()]; isObstacle {
				str += "#"
				continue
			}

			if _, hasVisited := lab.visited[position.toString()]; hasVisited {
				str += "X"
				continue
			}

			str += "."

		}
		str += "\n"
	}
	return str
}

func parseInput(input string) lab {
	labStrings := strings.Split(strings.TrimSpace(input), "\n")

	player := position{}
	obstacles := make(map[string]struct{})
	visited := make(map[string]struct{})

	for rowIndex, string := range labStrings {
		for colIndex, item := range string {
			if item == '#' {
				obstacle := position{row: rowIndex, col: colIndex}
				obstacles[obstacle.toString()] = struct{}{}
			}
			if item == '^' {
				player.row = rowIndex
				player.col = colIndex
			}
		}
	}

	visited[player.toString()] = struct{}{}

	return lab{
		rowLen: len(labStrings),
		// Lab should always exist, it's a puzzle after all
		colLen:    len(labStrings[0]),
		obstacles: obstacles,
		player:    player,
		visited:   visited,
	}
}

func getNextPositionInDirection(player position, direction direction) position {
	switch direction {
	case UP:
		return position{row: player.row - 1, col: player.col}
	case RIGTH:
		return position{row: player.row, col: player.col + 1}
	case DOWN:
		return position{row: player.row + 1, col: player.col}
	case LEFT:
		return position{row: player.row, col: player.col - 1}
	default:
		log.Printf("there is no such direction as: %d\n", direction)
		panic("crash and burn")
	}
}

func turnRight(direction direction) direction {
	switch direction {
	case UP:
		return RIGTH
	case RIGTH:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	default:
		log.Printf("there is no such direction as: %d\n", direction)
		panic("crash and burn")

	}

}

func PartOne(input string) int {
	lab := parseInput(input)
	log.Printf("lab: %v\n", lab)

	direction := UP
	for {
		nextPosition := getNextPositionInDirection(lab.player, direction)
		if nextPosition.isOutsideBounds(0, 0, lab.rowLen-1, lab.colLen-1) {
			log.Printf("next position is out of bounds! %#v\n", nextPosition)
			break
		}

		if _, isObstacle := lab.obstacles[nextPosition.toString()]; isObstacle {
			direction = turnRight(direction)
			log.Printf("changing direction obstacle detected: %#v\n", nextPosition)
			fmt.Println(lab.toString())
			continue
		}

		lab.player = nextPosition
		lab.visited[nextPosition.toString()] = struct{}{}
	}

	log.Printf("lab: %v\n", lab)
	return len(lab.visited)
}

func PartTwo(input string) int {
	return 0
}
