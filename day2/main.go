package main

import (
	"aoc/scanner"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	RED_CUBES   = 12
	GREEN_CUBES = 13
	BLUE_CUBES  = 14
)

type Set struct {
	redCubes   int
	greenCubes int
	blueCubes  int
}

func main() {
	gameIdsSum := 0
	power := 0
	stringSlice, err := scanner.ScanFile("day2/input.txt")

	if err != nil {
		panic(err)
	}

	for gameId, game := range stringSlice {
		setsStrings := strings.Split(strings.Split(game, ": ")[1], ";")
		sets := make([]Set, 0)
		set := Set{}

		for _, setString := range setsStrings {
			setString = strings.TrimSpace(setString)
			setArray := strings.Split(setString, ", ")

			for _, setString := range setArray {
				number, color := strings.Split(setString, " ")[0], strings.Split(setString, " ")[1]
				value, _ := strconv.Atoi(number)

				switch color {
				case "red":
					set.redCubes = value
				case "green":
					set.greenCubes = value
				case "blue":
					set.blueCubes = value
				}
				fmt.Println(number, color)
			}

			sets = append(sets, set)
		}

		isValid := true
		minRed, minGreen, minBlue := 0, 0, 0
		for _, set := range sets {
			if set.redCubes > RED_CUBES || set.greenCubes > GREEN_CUBES || set.blueCubes > BLUE_CUBES {
				isValid = false
			}

			minRed = int(math.Max(float64(set.redCubes), float64(minRed)))
			minGreen = int(math.Max(float64(set.greenCubes), float64(minGreen)))
			minBlue = int(math.Max(float64(set.blueCubes), float64(minBlue)))

		}
		fmt.Printf("game %d: %d red, %d green, %d blue\n", gameId+1, minRed, minGreen, minBlue)

		power += minRed * minGreen * minBlue

		if isValid {
			gameIdsSum += gameId + 1
		}
		fmt.Println("---")
	}

	fmt.Println(gameIdsSum)
	fmt.Println(power)
}
