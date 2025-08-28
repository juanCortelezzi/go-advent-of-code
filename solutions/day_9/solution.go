package dayNine

import (
	"log"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	hugeAssArray := make([]int, 0, 100_000)
	isFile := true
	fileCounter := 0
	for _, char := range strings.TrimSpace(input) {

		num, err := strconv.Atoi(string(char))
		if err != nil {
			log.Panicf("char is not an int: '%c'\n", char)
		}

		toAppend := 0
		if isFile {
			toAppend = fileCounter
			fileCounter += 1
		} else {
			toAppend = -1
		}

		for range num {
			hugeAssArray = append(hugeAssArray, toAppend)
		}

		isFile = !isFile
	}

	return hugeAssArray
}

func printProblem(line []int) string {
	str := ""
	for _, item := range line {
		if item == -1 {
			str += "."
		} else {
			str += strconv.Itoa(item)
		}
	}
	return str
}

func PartOne(input string) int {
	arr := parseInput(input)

	head := 0
	tail := len(arr) - 1

	for head < tail {
		if arr[head] != -1 {
			head += 1
			continue
		}

		if arr[tail] == -1 {
			tail -= 1
			continue
		}

		arr[head], arr[tail] = arr[tail], arr[head]
		head += 1
		tail -= 1
	}

	result := 0
	for index, num := range arr {
		if num == -1 {
			break
		}
		result += index * num
	}
	return result
}

func PartTwo(input string) int {
	return 0
}
