package main

import (
	"aoc/scanner"
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Copy struct {
	index int
	value int
}

func main() {
	stringSlice, err := scanner.ScanFile("day4/input.txt")
	points := 0
	instances := 0
	instancesArary := make([]int, len(stringSlice))
	cardCopies := make([]Copy, 0)

	if err != nil {
		panic(err)
	}

	for lineIndex, line := range stringSlice {
		numbers := strings.Split(line, ":")[1]
		winningNumbersStr, myNumbersStr := strings.Split(numbers, "|")[0], strings.Split(numbers, "|")[1]
		winningNumbers := getNumbers(winningNumbersStr)
		myNumbers := getNumbers(myNumbersStr)
		myWinningNumbers := make([]int, 0)

		for _, winningNumber := range winningNumbers {

			for _, myNumber := range myNumbers {

				if winningNumber == myNumber {
					myWinningNumbers = append(myWinningNumbers, winningNumber)
				}
			}
		}

		if len(myWinningNumbers) > 0 {
			//TODO: part 2
			cardIndex := slices.IndexFunc(cardCopies, func(c Copy) bool {
				return c.index == lineIndex
			})

			if cardIndex == -1 {
				instances++
				instancesArary[lineIndex] = 1
				for idx := range myWinningNumbers {
					copiesIndex := slices.IndexFunc(cardCopies, func(c Copy) bool {
						return c.index == lineIndex+idx+1
					})

					if copiesIndex == -1 {
						cardCopies = append(cardCopies, Copy{index: lineIndex + idx + 1, value: 1})
					} else {
						cardCopies[copiesIndex].value++
					}
				}
			} else {
				for cardCopies[cardIndex].value > 0 {
					instances++
					instancesArary[lineIndex]++
					for idx := range myWinningNumbers {
						copiesIndex := slices.IndexFunc(cardCopies, func(c Copy) bool {
							return c.index == lineIndex+idx+1
						})

						if copiesIndex == -1 {
							cardCopies = append(cardCopies, Copy{index: lineIndex + idx + 1, value: 1})
						} else {
							cardCopies[copiesIndex].value++
						}
					}
					cardCopies[cardIndex].value--
				}
			}
			//END TODO

			pts := 1 * int(math.Pow(2, float64(len(myWinningNumbers)-1)))
			fmt.Printf("Game %d won %d points\n", lineIndex+1, pts)
			fmt.Println(myWinningNumbers)
			points += pts
		}

	}

	fmt.Println(points)
	fmt.Println(cardCopies)
	fmt.Println(instancesArary)
	fmt.Println(instances)
}

func getNumbers(numString string) []int {
	index := 1
	numbers := make([]int, 0)

	str := ""

	for index < len(numString) {
		if index%3 != 0 {
			str = str + string(numString[index])
		} else {
			n, err := strconv.Atoi(strings.TrimSpace(str))

			if err != nil {
				panic(err)
			}

			numbers = append(numbers, n)
			str = ""
		}
		index++
	}

	if str != "" {
		n, err := strconv.Atoi(strings.TrimSpace(str))

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, n)
	}

	return numbers
}
