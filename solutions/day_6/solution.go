package day6

import (
	"fmt"
	"log"
	"strings"
)

type direction = int

const (
	up direction = iota
	right
	down
	left
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
}

func stateToString(lab lab, player position, visited map[string]struct{}) string {
	str := ""
	for rowIndex := range lab.rowLen {
		for colIndex := range lab.colLen {
			position := position{row: rowIndex, col: colIndex}
			if position == player {
				str += "^"
				continue
			}

			if _, isObstacle := lab.obstacles[position.toString()]; isObstacle {
				str += "#"
				continue
			}

			if _, hasVisited := visited[position.toString()]; hasVisited {
				str += "X"
				continue
			}

			str += "."

		}
		str += "\n"
	}
	return str
}

func parseInput(input string) (lab, position) {
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
		rowLen:    len(labStrings),
		colLen:    len(labStrings[0]), // Lab should always exist, it's a puzzle after all
		obstacles: obstacles,
	}, player
}

func getNextPositionInDirection(player position, direction direction) position {
	switch direction {
	case up:
		return position{row: player.row - 1, col: player.col}
	case right:
		return position{row: player.row, col: player.col + 1}
	case down:
		return position{row: player.row + 1, col: player.col}
	case left:
		return position{row: player.row, col: player.col - 1}
	default:
		log.Printf("there is no such direction as: %d\n", direction)
		panic("crash and burn")
	}
}

func turnRight(direction direction) direction {
	switch direction {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		log.Printf("there is no such direction as: %d\n", direction)
		panic("crash and burn")

	}

}

func PartOne(input string) int {
	lab, player := parseInput(input)
	log.Printf("lab: %v\n", lab)

	visited := make(map[string]struct{})
	direction := up
	for {
		nextPosition := getNextPositionInDirection(player, direction)
		if nextPosition.isOutsideBounds(0, 0, lab.rowLen-1, lab.colLen-1) {
			log.Printf("next position is out of bounds! %#v\n", nextPosition)
			fmt.Println(stateToString(lab, player, visited))
			break
		}

		if _, isObstacle := lab.obstacles[nextPosition.toString()]; isObstacle {
			direction = turnRight(direction)
			log.Printf("changing direction obstacle detected: %#v\n", nextPosition)
			fmt.Println(stateToString(lab, player, visited))
			continue
		}

		player = nextPosition
		visited[nextPosition.toString()] = struct{}{}
	}

	log.Printf("lab: %v\n", lab)
	return len(visited)
}

func isLoop(
	lab lab,
	player position,
	playerDirection direction,
) bool {
	turnPositionWithDirection := make(map[string]struct{})
	for {
		nextPosition := getNextPositionInDirection(player, playerDirection)
		if nextPosition.isOutsideBounds(0, 0, lab.rowLen-1, lab.colLen-1) {
			return false
		}

		if _, isObstacle := lab.obstacles[nextPosition.toString()]; isObstacle {
			positionWithDirection := fmt.Sprintf("%d-%d-%d", player.row, player.col, playerDirection)
			if _, hasVisited := turnPositionWithDirection[positionWithDirection]; hasVisited {
				return true
			}
			turnPositionWithDirection[positionWithDirection] = struct{}{}
			playerDirection = turnRight(playerDirection)
			continue
		}

		player = nextPosition
	}
}

func PartTwo(input string) int {
	lab, player := parseInput(input)
	visited := make(map[string]struct{})
	direction := up
	loopPositions := make([]position, 0, 10)
	for {
		nextPosition := getNextPositionInDirection(player, direction)

		if nextPosition.isOutsideBounds(0, 0, lab.rowLen-1, lab.colLen-1) {
			break
		}

		if _, isObstacle := lab.obstacles[nextPosition.toString()]; isObstacle {
			direction = turnRight(direction)
			continue
		}

		if _, hasVisited := visited[nextPosition.toString()]; !hasVisited {
			lab.obstacles[nextPosition.toString()] = struct{}{}
			if isLoop(lab, player, direction) {
				loopPositions = append(loopPositions, nextPosition)
			}
			delete(lab.obstacles, nextPosition.toString())
		}

		player = nextPosition
		visited[nextPosition.toString()] = struct{}{}
	}

	log.Printf("lab: %v\n", lab)
	log.Printf("loops: %v\n", loopPositions)
	return len(loopPositions)
}
