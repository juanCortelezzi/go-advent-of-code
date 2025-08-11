package main

import (
	"log"
	"strconv"
	"strings"
)

func DayTwoPartOne(input string) int {
	reports := strings.SplitSeq(strings.TrimSpace(input), "\n")
	safeCount := 0
	for report := range reports {
		levels := strings.Split(report, " ")

		if len(levels) == 0 {
			log.Fatalf("there are no levels in this report: '%s'\n", report)
		}

		if len(levels) == 1 {
			log.Printf("there is only one level in this report '%s'\n", report)
			safeCount += 1
			continue
		}

		log.Printf("levels: %#v\n", levels)

		firstLevel, err := strconv.Atoi(levels[0])
		if err != nil {
			log.Fatalf("first level of report '%s' is not an int: '%s'", report, levels[0])
		}

		prevLevel := firstLevel
		var direction string
		isValidReport := true
		for _, levelString := range levels[1:] {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				log.Fatalf("level of report '%s' is not an int: '%s'", report, levelString)
			}

			log.Printf("prev: %d | curr: %d\n", prevLevel, level)

			difference := 0
			if prevLevel > level {
				if direction == "" {
					direction = "desc"
				}
				if direction != "desc" {
					prevLevel = level
					isValidReport = false
					break
				}
				difference = prevLevel - level
			} else {
				if direction == "" {
					direction = "asc"
				}
				if direction != "asc" {
					prevLevel = level
					isValidReport = false
					break
				}
				difference = level - prevLevel
			}

			if difference < 1 || difference > 3 {
				prevLevel = level
				isValidReport = false
				break
			}

			prevLevel = level
		}

		if isValidReport {
			safeCount += 1
		}
	}

	return safeCount
}

func DayTwoPartTwo(input string) int {
	return 1
}
