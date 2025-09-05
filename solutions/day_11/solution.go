package dayEleven

import (
	"log"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	stoneStrings := strings.Split(strings.TrimSpace(input), " ")
	stones := make([]int, 0, len(stoneStrings))
	for _, stoneString := range stoneStrings {
		stone, err := strconv.Atoi(stoneString)
		if err != nil {
			log.Panicf("stone '%s' can not be converted to integer", stoneString)
		}
		stones = append(stones, stone)
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

func PartOne(input string) int {
	stones := parseInput(input)

	for range 25 {
		newStones := make([]int, 0, len(stones)*2)

		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
				continue
			}

			numberOfDigits := getNumberOfDigits(stone)
			hasEvenNumberOfDigits := (numberOfDigits & 1) == 0
			if hasEvenNumberOfDigits {
				splitStoneA, splitStoneB := splitFirstNDigitsString(stone, numberOfDigits>>1)
				newStones = append(newStones, splitStoneA, splitStoneB)
				continue
			}

			newStones = append(newStones, stone*2024)
		}

		stones = newStones
	}

	return len(stones)
}

func PartTwo(input string) int {
	return 0
}
