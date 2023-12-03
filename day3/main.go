package main

import (
	"aoc/scanner"
	"fmt"
	"strconv"
)

func main() {
	stringSlice, err := scanner.ScanFile("day3/input.txt")
	sumAdjacents := 0
	gearAdjacentNumber := 0

	if err != nil {
		panic(err)
	}

	// map each to a character array

	for lineIndex, line := range stringSlice {
		lineArray := []byte(line)
		currentNumber := 0
		for charIndex := 0; charIndex < len(lineArray); charIndex++ {
			if lineArray[charIndex] == '*' {
				adjacentNumbers := checkGearAdjacentNumbers(stringSlice, lineIndex, charIndex)

				if len(adjacentNumbers) == 2 {
					if adjacentNumbers[0] > adjacentNumbers[1] {
						for i, j := 0, len(adjacentNumbers)-1; i < j; i, j = i+1, j-1 {
							adjacentNumbers[i], adjacentNumbers[j] = adjacentNumbers[j], adjacentNumbers[i]
						}
					}

					fmt.Println("gear has adjacent", adjacentNumbers)
					gearAdjacentNumber += adjacentNumbers[0] * adjacentNumbers[1]
				}
			}

			if isDigit(lineArray[charIndex]) {
				currentNumber = currentNumber*10 + int(lineArray[charIndex]-'0')
			} else if currentNumber != 0 {
				hasAdjacent := checkAdjacentCharacters(stringSlice, lineIndex, charIndex, currentNumber)

				if hasAdjacent {
					// fmt.Println("has adjacent", currentNumber)
					sumAdjacents += currentNumber
				}

				currentNumber = 0
			}
		}

		if currentNumber != 0 {
			hasAdjacent := checkAdjacentCharacters(stringSlice, lineIndex, len(lineArray), currentNumber)

			if hasAdjacent {
				sumAdjacents += currentNumber
			}

			currentNumber = 0
		}

	}

	fmt.Println(sumAdjacents)
	fmt.Println(gearAdjacentNumber)
}

func checkAdjacentCharacters(stringSlice []string, lineIndex int, charIndex int, currentNumber int) bool {
	currentNumberLength := len(fmt.Sprintf("%d", currentNumber))
	hasAdjacent := false

	// check left
	if charIndex-currentNumberLength-1 > 0 {
		if stringSlice[lineIndex][charIndex-currentNumberLength-1] != '.' && (stringSlice[lineIndex][charIndex-currentNumberLength-1] < '0' || stringSlice[lineIndex][charIndex-currentNumberLength-1] > '9') {
			hasAdjacent = true
		}
	}

	// check right
	if charIndex+currentNumberLength < len(stringSlice[lineIndex]) {
		if stringSlice[lineIndex][charIndex] != '.' && (stringSlice[lineIndex][charIndex] < '0' || stringSlice[lineIndex][charIndex] > '9') {
			hasAdjacent = true
		}
	}

	// check up and up corners
	if lineIndex-1 > 0 {
		for i := 0; i < currentNumberLength+2; i++ {
			if charIndex-i >= 0 && charIndex-i < len(stringSlice[lineIndex-1]) {
				if stringSlice[lineIndex-1][charIndex-i] != '.' && (stringSlice[lineIndex-1][charIndex-i] < '0' || stringSlice[lineIndex-1][charIndex-i] > '9') {
					hasAdjacent = true
				}
			}
		}
	}

	// check down and down corners
	if lineIndex+1 < len(stringSlice) {
		for i := 0; i < currentNumberLength+2; i++ {
			if charIndex-i >= 0 && charIndex-i < len(stringSlice[lineIndex+1]) {
				if stringSlice[lineIndex+1][charIndex-i] != '.' && (stringSlice[lineIndex+1][charIndex-i] < '0' || stringSlice[lineIndex+1][charIndex-i] > '9') {
					hasAdjacent = true
				}
			}
		}
	}

	return hasAdjacent
}

func checkGearAdjacentNumbers(stringSlice []string, lineIndex int, charIndex int) []int {
	adjacentNumbers := make([]int, 0)
	// check left
	if charIndex-1 > 0 {
		if isDigit(stringSlice[lineIndex][charIndex-1]) {
			numberStr := string(stringSlice[lineIndex][charIndex-1])
			isNumber := true
			index := 1

			for isNumber {
				if charIndex-1-index >= 0 && isDigit(stringSlice[lineIndex][charIndex-1-index]) {
					numberStr = string(stringSlice[lineIndex][charIndex-1-index]) + numberStr
					index++
				} else {
					isNumber = false
				}
			}
			number, _ := strconv.Atoi(string(numberStr))
			adjacentNumbers = append(adjacentNumbers, number)
		}
	}

	// check right
	if charIndex+1 < len(stringSlice[lineIndex]) {
		if isDigit(stringSlice[lineIndex][charIndex+1]) {
			numberStr := string(stringSlice[lineIndex][charIndex+1])
			isNumber := true
			index := 1

			for isNumber {
				if charIndex+1+index < len(stringSlice[lineIndex]) && isDigit(stringSlice[lineIndex][charIndex+1+index]) {
					numberStr = numberStr + string(stringSlice[lineIndex][charIndex+1+index])
					index++
				} else {
					isNumber = false
				}
			}
			number, _ := strconv.Atoi(string(numberStr))
			adjacentNumbers = append(adjacentNumbers, number)
		}
	}

	//check top
	if lineIndex-1 >= 0 {
		middleEmpty := false
		if stringSlice[lineIndex-1][charIndex] == '.' {
			middleEmpty = true
		}

		if charIndex-1 > 0 {
			if middleEmpty {
				if isDigit(stringSlice[lineIndex-1][charIndex-1]) {
					numberStr := string(stringSlice[lineIndex-1][charIndex-1])
					isNumber := true
					index := 1

					for isNumber {
						if charIndex-1-index >= 0 && charIndex-1-index < len(stringSlice[lineIndex-1]) && isDigit(stringSlice[lineIndex-1][charIndex-1-index]) {
							numberStr = string(stringSlice[lineIndex-1][charIndex-1-index]) + numberStr
							index++
						} else {
							isNumber = false
						}
					}

					number, _ := strconv.Atoi(string(numberStr))
					adjacentNumbers = append(adjacentNumbers, number)
				}

				if isDigit(stringSlice[lineIndex-1][charIndex+1]) {
					numberStr := string(stringSlice[lineIndex-1][charIndex+1])
					isNumber := true
					index := 1

					for isNumber {
						if charIndex+1+index < len(stringSlice[lineIndex-1]) && isDigit(stringSlice[lineIndex-1][charIndex+1+index]) {
							numberStr = numberStr + string(stringSlice[lineIndex-1][charIndex+1+index])
							index++
						} else {
							isNumber = false
						}
					}

					number, _ := strconv.Atoi(string(numberStr))
					adjacentNumbers = append(adjacentNumbers, number)
				}
			} else {
				if isDigit(stringSlice[lineIndex-1][charIndex-1]) {
					numberStr := string(stringSlice[lineIndex-1][charIndex-1])
					rightIndex := 1
					leftIndex := 1
					isNumber := true

					for isNumber {
						if charIndex-1+rightIndex < len(stringSlice[lineIndex-1]) && isDigit(stringSlice[lineIndex-1][charIndex-1+rightIndex]) {
							numberStr = numberStr + string(stringSlice[lineIndex-1][charIndex-1+rightIndex])
							rightIndex++
						} else if charIndex-1-leftIndex >= 0 && isDigit(stringSlice[lineIndex-1][charIndex-1-leftIndex]) {
							numberStr = string(stringSlice[lineIndex-1][charIndex-1-leftIndex]) + numberStr
							leftIndex++
						} else {
							isNumber = false
						}
					}

					number, _ := strconv.Atoi(string(numberStr))
					adjacentNumbers = append(adjacentNumbers, number)
				} else if isDigit(stringSlice[lineIndex-1][charIndex+1]) {
					numberStr := string(stringSlice[lineIndex-1][charIndex+1])
					rightIndex := 1
					leftIndex := 1
					isNumber := true

					for isNumber {
						if charIndex+1+rightIndex < len(stringSlice[lineIndex-1]) && isDigit(stringSlice[lineIndex-1][charIndex+1+rightIndex]) {
							numberStr = numberStr + string(stringSlice[lineIndex-1][charIndex+1+rightIndex])
							rightIndex++
						} else if charIndex+1+leftIndex < len(stringSlice[lineIndex-1]) && isDigit(stringSlice[lineIndex-1][charIndex+1-leftIndex]) {
							numberStr = string(stringSlice[lineIndex-1][charIndex+1-leftIndex]) + numberStr
							leftIndex++
						} else {
							isNumber = false
						}
					}

					number, _ := strconv.Atoi(string(numberStr))
					adjacentNumbers = append(adjacentNumbers, number)
				} else if isDigit(stringSlice[lineIndex-1][charIndex]) {

					numberStr := string(stringSlice[lineIndex-1][charIndex])
					isNumber := true
					index := 1

					for isNumber {
						if charIndex-index >= 0 && isDigit(stringSlice[lineIndex-1][charIndex-index]) {
							numberStr = string(stringSlice[lineIndex-1][charIndex-index]) + numberStr
							index++
						} else {
							isNumber = false
						}
					}

					number, _ := strconv.Atoi(string(numberStr))
					adjacentNumbers = append(adjacentNumbers, number)

					if isDigit(stringSlice[lineIndex-1][charIndex+1]) {
						numberStr := string(stringSlice[lineIndex-1][charIndex+1])
						isNumber := true
						index := 1

						for isNumber {
							if charIndex+1+index < len(stringSlice[lineIndex-1]) && isDigit(stringSlice[lineIndex-1][charIndex+1+index]) {
								numberStr = numberStr + string(stringSlice[lineIndex-1][charIndex+1+index])
								index++
							} else {
								isNumber = false
							}
						}

						number, _ := strconv.Atoi(string(numberStr))
						adjacentNumbers = append(adjacentNumbers, number)
					}
				}
			}
		}
	}

	//check bottom
	if lineIndex+1 < len(stringSlice) {
		middleEmpty := false
		if stringSlice[lineIndex+1][charIndex] == '.' {
			middleEmpty = true
		}

		if charIndex-1 >= 0 {
			if middleEmpty {
				if isDigit(stringSlice[lineIndex+1][charIndex-1]) {
					numberStr := ""
					isNumber := true
					index := 1

					for isNumber {
						if charIndex-index >= 0 && isDigit(stringSlice[lineIndex+1][charIndex-index]) {
							numberStr = string(stringSlice[lineIndex+1][charIndex-index]) + numberStr
							index++
						} else {
							isNumber = false
						}
					}

					number, _ := strconv.Atoi(string(numberStr))
					adjacentNumbers = append(adjacentNumbers, number)
				}

				if isDigit(stringSlice[lineIndex+1][charIndex+1]) {
					numberStr := string(stringSlice[lineIndex+1][charIndex+1])
					isNumber := true
					index := 1

					for isNumber {
						if charIndex+1+index < len(stringSlice[lineIndex+1]) && isDigit(stringSlice[lineIndex+1][charIndex+1+index]) {
							numberStr = numberStr + string(stringSlice[lineIndex+1][charIndex+1+index])
							index++
						} else {
							isNumber = false
						}
					}

					number, _ := strconv.Atoi(string(numberStr))
					adjacentNumbers = append(adjacentNumbers, number)
				}
			} else {
				if isDigit(stringSlice[lineIndex+1][charIndex-1]) {
					numberStr := string(stringSlice[lineIndex+1][charIndex-1])
					rightIndex := 1
					leftIndex := 1
					isNumber := true

					for isNumber {
						if charIndex-1-rightIndex >= 0 && isDigit(stringSlice[lineIndex+1][charIndex-1+rightIndex]) {
							numberStr = numberStr + string(stringSlice[lineIndex+1][charIndex-1+rightIndex])
							rightIndex++
						} else if charIndex-1-leftIndex >= 0 && isDigit(stringSlice[lineIndex+1][charIndex-1-leftIndex]) {
							numberStr = string(stringSlice[lineIndex+1][charIndex-1-leftIndex]) + numberStr
							leftIndex++
						} else {
							isNumber = false
						}
					}

					number, _ := strconv.Atoi(string(numberStr))
					adjacentNumbers = append(adjacentNumbers, number)
				} else if isDigit(stringSlice[lineIndex+1][charIndex+1]) {
					numberStr := string(stringSlice[lineIndex+1][charIndex+1])
					rightIndex := 1
					leftIndex := 1
					isNumber := true

					for isNumber {
						if charIndex+1+rightIndex < len(stringSlice[lineIndex+1]) && isDigit(stringSlice[lineIndex+1][charIndex+1+rightIndex]) {
							numberStr = numberStr + string(stringSlice[lineIndex+1][charIndex+1+rightIndex])
							rightIndex++
						} else if charIndex+1+leftIndex < len(stringSlice[lineIndex+1]) && isDigit(stringSlice[lineIndex+1][charIndex+1-leftIndex]) {
							numberStr = string(stringSlice[lineIndex+1][charIndex+1-leftIndex]) + numberStr
							leftIndex++
						} else {
							isNumber = false
						}
					}

					number, _ := strconv.Atoi(string(numberStr))
					adjacentNumbers = append(adjacentNumbers, number)
				} else if isDigit(stringSlice[lineIndex+1][charIndex]) {

					numberStr := string(stringSlice[lineIndex+1][charIndex])
					isNumber := true
					index := 1

					for isNumber {
						if charIndex-index >= 0 && isDigit(stringSlice[lineIndex+1][charIndex-index]) {
							numberStr = string(stringSlice[lineIndex+1][charIndex-index]) + numberStr
							index++
						} else {
							isNumber = false
						}
					}

					number, _ := strconv.Atoi(string(numberStr))
					adjacentNumbers = append(adjacentNumbers, number)

					if isDigit(stringSlice[lineIndex+1][charIndex+1]) {
						numberStr := string(stringSlice[lineIndex+1][charIndex+1])
						isNumber := true
						index := 1

						for isNumber {
							if charIndex+1+index < len(stringSlice[lineIndex+1]) && isDigit(stringSlice[lineIndex+1][charIndex+1+index]) {
								numberStr = numberStr + string(stringSlice[lineIndex+1][charIndex+1+index])
								index++
							} else {
								isNumber = false
							}
						}

						number, _ := strconv.Atoi(string(numberStr))
						adjacentNumbers = append(adjacentNumbers, number)
					}
				}
			}
		}
	}

	return adjacentNumbers
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
