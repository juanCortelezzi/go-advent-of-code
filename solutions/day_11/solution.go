package dayEleven

import (
	"log"
	"strconv"
	"strings"
)

func parseInput(input string) map[int]int {
	stoneStrings := strings.Split(strings.TrimSpace(input), " ")
	stones := make(map[int]int)
	for _, stoneString := range stoneStrings {
		stone, err := strconv.Atoi(stoneString)
		if err != nil {
			log.Panicf("stone '%s' can not be converted to integer", stoneString)
		}
		stones[stone] += 1
	}

	return stones
}

func getNumberOfDigits(number int) int {
	if number == 0 {
		return 1
	}
	if number < 0 {
		number = -number
	}

	count := 0
	for number > 0 {
		number /= 10
		count++
	}
	return count
}

func splitFirstNDigitsString(num, n int) (int, int) {
	if n <= 0 {
		return 0, num
	}

	strNum := strconv.Itoa(num)

	firstPartStr := strNum[:n]
	secondPartStr := strNum[n:]

	firstPart, _ := strconv.Atoi(firstPartStr)
	secondPart, _ := strconv.Atoi(secondPartStr)

	return firstPart, secondPart
}

func updateStone(stone int) []int {
	newStones := make([]int, 0, 2)
	if stone == 0 {
		newStones = append(newStones, 1)
		return newStones
	}

	numberOfDigits := getNumberOfDigits(stone)
	hasEvenNumberOfDigits := (numberOfDigits & 1) == 0
	if hasEvenNumberOfDigits {
		splitStoneA, splitStoneB := splitFirstNDigitsString(stone, numberOfDigits>>1)
		newStones = append(newStones, splitStoneA, splitStoneB)
		return newStones
	}

	newStones = append(newStones, stone*2024)
	return newStones
}

func PartOne(input string) int {
	stones := parseInput(input)

	for range 25 {
		newStones := make(map[int]int, len(stones))
		for stone, currentQuantity := range stones {
			for _, newStone := range updateStone(stone) {
				newStones[newStone] += currentQuantity
			}
		}

		stones = newStones
	}

	result := 0
	for _, value := range stones {
		result += value
	}

	return result
}

func PartTwo(input string) int {
	stones := parseInput(input)

	for range 75 {
		newStones := make(map[int]int, len(stones))
		for stone, currentQuantity := range stones {
			for _, newStone := range updateStone(stone) {
				newStones[newStone] += currentQuantity
			}
		}

		stones = newStones
	}

	result := 0
	for _, value := range stones {
		result += value
	}

	return result
}
