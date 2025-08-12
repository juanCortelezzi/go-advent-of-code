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
					isValidReport = false
					break
				}
				difference = prevLevel - level
			} else {
				if direction == "" {
					direction = "asc"
				}
				if direction != "asc" {
					isValidReport = false
					break
				}
				difference = level - prevLevel
			}

			if difference < 1 || difference > 3 {
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
	reports := strings.SplitSeq(strings.TrimSpace(input), "\n")
	safeCount := 0
	for reportString := range reports {
		levelStrings := strings.Split(reportString, " ")
		report := make([]int, 0, len(levelStrings))
		for _, levelString := range levelStrings {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				log.Fatalf("level of report '%s' is not an int: '%s'", reportString, levelString)
			}

			report = append(report, level)
		}

		isValid, errorIndex := isValidReport(report)
		if isValid {
			safeCount += 1
			continue
		}

		for index := range errorIndex + 1 {
			x := make([]int, 0, len(report)-1)
			x = append(x, report[0:index]...)
			x = append(x, report[index+1:]...)
			if valid, _ := isValidReport(x); valid {
				safeCount += 1
				break
			}
		}

	}

	return safeCount
}

func isValidReport(report []int) (bool, int) {
	log.Println("-------------------------------")
	log.Printf("report: %#v\n", report)

	if len(report) == 0 {
		log.Fatalf("there are no levels in this report: '%#v'\n", report)
	}

	if len(report) == 1 {
		log.Printf("there is only one level in this report '%#v'\n", report)
		return true, 0
	}

	if len(report) == 2 {
		log.Printf("there is only two levels in this report '%#v'\n", report)
		return true, 0
	}

	direction := ""
	index := 1
	prevLevel := report[0]
	for index < len(report) {
		level := report[index]
		log.Printf("p: %d | c: %d\n", prevLevel, level)

		difference := 0
		if prevLevel > level {
			if direction == "" {
				direction = "desc"
			}
			if direction != "desc" {
				log.Printf("mistake not desc @: %d\n", index)
				return false, index
			}
			difference = prevLevel - level
		} else {
			if direction == "" {
				direction = "asc"
			}
			if direction != "asc" {
				log.Printf("mistake not asc @: %d\n", index)
				return false, index
			}
			difference = level - prevLevel
		}

		if difference < 1 || difference > 3 {
			log.Printf("mistake step count @: %d\n", index)
			return false, index
		}

		prevLevel = level
		index += 1
	}

	return true, 0
}
