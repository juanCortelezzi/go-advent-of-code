package dayNine

import (
	"log"
	"strconv"
	"strings"
)

type disk struct {
	numberOfFiles int
	fileSystem    []int
}

func parseInput(input string) disk {
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

	return disk{
		numberOfFiles: fileCounter,
		fileSystem:    hugeAssArray,
	}
}

func problemToString(line []int, head, tail int) string {
	str := ""
	for _, item := range line {
		if item == -1 {
			str += "."
		} else {
			str += strconv.Itoa(item)
		}
	}

	str += "\n"
	for index := range line {
		if index == head {
			str += "h"
			continue
		}

		if index == tail {
			str += "t"
			continue
		}

		str += " "
	}

	return str
}

func PartOne(input string) int {
	disk := parseInput(input)

	head := 0
	tail := len(disk.fileSystem) - 1

	for head < tail {
		if disk.fileSystem[head] != -1 {
			head += 1
			continue
		}

		if disk.fileSystem[tail] == -1 {
			tail -= 1
			continue
		}

		disk.fileSystem[head], disk.fileSystem[tail] = disk.fileSystem[tail], disk.fileSystem[head]
		head += 1
		tail -= 1
	}

	result := 0
	for index, num := range disk.fileSystem {
		if num == -1 {
			break
		}
		result += index * num
	}
	return result
}

func indexOfGapWithSize(fileSystem []int, size int) int {
	start := -1
	for index, item := range fileSystem {
		isSpace := item == -1

		if !isSpace {
			start = -1
			continue
		}

		startIsEmpty := start == -1
		if startIsEmpty {
			start = index
		}

		if (index - start) == size-1 {
			return start
		}
	}
	return -1
}

func PartTwo(input string) int {
	disk := parseInput(input)

	seenFileIds := make(map[int]struct{})

	tail := len(disk.fileSystem) - 1

	for len(seenFileIds) != disk.numberOfFiles {
		fileId := disk.fileSystem[tail]
		if fileId == -1 {
			tail -= 1
			continue
		}

		seenFileIds[fileId] = struct{}{}

		fileLen := 0
		for tail-fileLen >= 0 && disk.fileSystem[tail-fileLen] == fileId {
			fileLen += 1
		}

		gapStartIndex := indexOfGapWithSize(disk.fileSystem[:tail], fileLen)
		if gapStartIndex == -1 {
			tail -= fileLen
			continue
		}

		for delta := range fileLen {
			disk.fileSystem[gapStartIndex+delta], disk.fileSystem[tail-delta] = disk.fileSystem[tail-delta], disk.fileSystem[gapStartIndex+delta]
		}

		tail -= fileLen

	}

	result := 0
	for index, num := range disk.fileSystem {
		if num == -1 {
			continue
		}
		result += index * num
	}
	return result
}
