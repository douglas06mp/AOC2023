package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"AOC2023/helper"
)

func main() {
	lines, err := helper.ReadInput("solutions/day1/input.txt")
	if err != nil {
		panic(err)
	}

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	sum := 0

	for _, line := range lines {
		first, last := findFirstAndLastDigit(line)

		if first != 0 && last != 0 {
			num, _ := strconv.Atoi(string(first) + string(last))
			sum += num
		}
	}

	fmt.Println(sum)
}

func findFirstAndLastDigit(str string) (byte, byte) {
	var first, last byte

	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			if first == 0 {
				first = str[i]
			}

			last = str[i]
		}
	}

	return first, last
}

func part2(lines []string) {
	sum := 0

	for _, line := range lines {
		first, last := findFirstAndLastPattern(line)
		sum += first*10 + last
	}

	fmt.Println(sum)
}

func findFirstAndLastPattern(str string) (int, int) {
	wordToNum := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var firstNum, lastNum, firstIndex int = 0, 0, math.MaxInt32

	for i := 0; i < len(str); i++ {
		for word, num := range wordToNum {
			if strings.HasPrefix(str[i:], word) {
				if i < firstIndex {
					firstIndex, firstNum = i, num
				}
				lastNum = num
			}
		}

		if str[i] >= '0' && str[i] <= '9' {
			num, _ := strconv.Atoi(string(str[i]))

			if i < firstIndex {
				firstIndex, firstNum = i, num
			}
			lastNum = num
		}
	}

	return firstNum, lastNum
}
