package main

import (
	"aoc/scanner"
	"fmt"
	"strings"
)

func main() {
	numbers := map[string]int{
		"zero":  0,
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

	stringSlice, err := scanner.ScanFile("day1/input.txt")

	total := 0

	if err != nil {
		fmt.Println(err)
	}

	for _, line := range stringSlice {
		for numberString, numberValue := range numbers {
			// do not replace first and last characters because those characters may be part of another number
			line = strings.ReplaceAll(line, numberString, (fmt.Sprintf(numberString[0:1]) + fmt.Sprintf("%d", numberValue) + numberString[len(numberString)-1:]))
		}

		numberResult := 0
		firstDigit := -1
		lastDigit := -1

		for _, char := range line {
			if char >= '0' && char <= '9' {
				if firstDigit == -1 {
					firstDigit = int(char - '0')
				} else {
					lastDigit = int(char - '0')
				}
			}

			if firstDigit != -1 {
				numberResult = firstDigit
			}

			if lastDigit == -1 {
				lastDigit = firstDigit
			}

			numberResult = numberResult*10 + lastDigit
		}

		total += numberResult

		fmt.Println(numberResult)
	}

	fmt.Println(total)
}
