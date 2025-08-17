package dayFive

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

// format: `X|Y`
// if both X and Y are present in the pages to be changed then:
//     X must be some time before Y

type Input = struct {
	rulesLength int
	xs          []int
	ys          []int
	manuals     [][]int
}

func parseInput(input string) Input {
	rulesSection, manualsSection, found := strings.Cut(input, "\n\n")
	if !found {
		log.Fatalln("could not split input between rules and manuals")
	}

	xs, ys := parseRules(rulesSection)
	if len(xs) != len(ys) {
		log.Fatalln("xs and ys should be the same length! crash and burn!!")
	}

	manuals := parseManuals(manualsSection)

	return Input{
		rulesLength: len(xs),
		xs:          xs,
		ys:          ys,
		manuals:     manuals,
	}
}

func parseManuals(manuals string) [][]int {
	manualStrings := strings.Split(strings.TrimSpace(manuals), "\n")
	result := make([][]int, 0, len(manualStrings))
	for _, manualString := range manualStrings {

		pageStrings := strings.Split(manualString, ",")
		pages := make([]int, 0, len(pageStrings))
		for _, pageString := range pageStrings {
			page, err := strconv.Atoi(pageString)
			if err != nil {
				log.Fatalf("page of manual '%s' is not an int: '%s'\n", manualString, pageString)
			}
			pages = append(pages, page)
		}
		result = append(result, pages)
	}

	return result
}

func parseRules(rules string) ([]int, []int) {
	ruleStrings := strings.Split(rules, "\n")
	xs := make([]int, 0, len(ruleStrings))
	ys := make([]int, 0, len(ruleStrings))
	for _, ruleString := range ruleStrings {
		xString, yString, found := strings.Cut(ruleString, "|")
		if !found {
			log.Fatalf("could not split x and y from rule: \"%s\"\n", ruleString)
		}

		x, err := strconv.Atoi(xString)
		if err != nil {
			log.Fatalf("first number of rule '%s' is not an int: '%s'\n", ruleString, xString)
		}

		y, err := strconv.Atoi(yString)
		if err != nil {
			log.Fatalf("second number of rule '%s' is not an int: '%s'\n", ruleString, yString)
		}

		xs = append(xs, x)
		ys = append(ys, y)
	}
	return xs, ys
}

func getValidRulesForManual(manual []int, xs, ys []int) ([]int, []int) {
	xsValid := make([]int, 0)
	ysValid := make([]int, 0)
	for ruleIndex := range len(xs) {
		x := xs[ruleIndex]
		y := ys[ruleIndex]
		if slices.Contains(manual, x) && slices.Contains(manual, y) {
			log.Printf("rule: %d|%d\n", x, y)
			xsValid = append(xsValid, x)
			ysValid = append(ysValid, y)
		}
	}
	return xsValid, ysValid
}

func isValidManual(manual []int, xs []int, ys []int) bool {
	log.Printf("checking manual: %#v\n", manual)
	for pageIndex, page := range manual {
		for xIndex, x := range xs {
			if x != page {
				continue
			}
			y := ys[xIndex]
			if slices.Contains(manual[0:pageIndex], y) {
				log.Printf("page %d failed before rule: %d|%d\n", page, x, y)
				return false
			}
		}

		// for yIndex, y := range ys {
		// 	if y != page {
		// 		continue
		// 	}
		// 	x := xs[yIndex]
		// 	if slices.Contains(manual[pageIndex+1:], x) {
		// 		log.Printf("page %d failed after rule: %d|%d\n", page, x, y)
		// 		return false
		// 	}
		// }
	}

	return true
}

func PartOne(input string) int {
	data := parseInput(input)

	result := 0
	for _, manual := range data.manuals {
		xsValid, ysValid := getValidRulesForManual(manual, data.xs, data.ys)
		if isValidManual(manual, xsValid, ysValid) {
			midNumber := manual[(len(manual)-1)>>1]
			log.Printf("manual %#v is valid, mid number: %d\n", manual, midNumber)
			result += midNumber
		}
	}

	return result
}

func PartTwo(input string) int {
	data := parseInput(input)

	result := 0
	for _, manual := range data.manuals {
		xsValid, ysValid := getValidRulesForManual(manual, data.xs, data.ys)
		if isValidManual(manual, xsValid, ysValid) {
			continue
		}

		slices.SortFunc(manual, func(a, b int) int {
			for index := range len(xsValid) {
				x := xsValid[index]
				y := ysValid[index]

				if x == a && y == b {
					return -1
				}
			}
			return 0
		})

		midNumber := manual[(len(manual)-1)>>1]
		log.Printf("manual %#v is valid, mid number: %d\n", manual, midNumber)
		result += midNumber
	}

	return result
}
