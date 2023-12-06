package main

import (
	"fmt"
	"strconv"
	"strings"

	"AOC2023/helper"
)

func main() {
	lines, err := helper.ReadInput("solutions/day2/input.txt")
	if err != nil {
		panic(err)
	}

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	ans := 0

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		game, data := parts[0], parts[1]

		if checkPossible(data) {
			gameIdx, _ := strconv.Atoi(strings.Split(game, " ")[1])
			ans += gameIdx
		}

	}

	fmt.Printf("part1: %d\n", ans)
}

func checkPossible(data string) bool {
	limit := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	rounds := strings.Split(data, "; ")

	for _, round := range rounds {
		for _, r := range strings.Split(round, ",") {
			r = strings.Trim(r, " ")
			parts := strings.Split(r, " ")
			rColor := parts[1]
			rValue, _ := strconv.Atoi(parts[0])

			if rValue > limit[rColor] {
				return false
			}
		}
	}

	return true
}

func part2(lines []string) {
	ans := 0

	for _, line := range lines {
		data := strings.Split(line, ": ")[1]
		rounds := strings.Split(data, "; ")

		redMax, greenMax, blueMax := findMaxForEachColor(rounds)
		ans += redMax * greenMax * blueMax
	}

	fmt.Printf("part2: %d\n", ans)
}

func findMaxForEachColor(rounds []string) (int, int, int) {
	max := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, round := range rounds {
		for _, r := range strings.Split(round, ",") {
			r = strings.Trim(r, " ")
			parts := strings.Split(r, " ")
			rColor := parts[1]
			rValue, _ := strconv.Atoi(parts[0])

			if rValue > max[rColor] {
				max[rColor] = rValue
			}
		}
	}

	return max["red"], max["green"], max["blue"]
}
