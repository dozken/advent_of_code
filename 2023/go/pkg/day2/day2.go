package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type parseLine func(string) (int, error)

type Game struct {
	id int
	// cubes []Cube
	red   int
	blue  int
	green int
}

// type Cube struct {
//     red int
//     blue int
//     green int
// }

// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
func ParseLinePart1(line string) (int, error) {
	parts := strings.Split(line, ":")

	id, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
	if err != nil {
		return 0, err
	}

	game := &Game{
		id:    id,
		red:   0,
		blue:  0,
		green: 0,
	}

	cubeSubsets := strings.Split(parts[1], ";")
	for _, cubeSubset := range cubeSubsets {
		cubes := strings.Split(cubeSubset, ",")

		for _, cube := range cubes {
			cubeParts := strings.Split(strings.Trim(cube, " "), " ")

			digit, _ := strconv.Atoi(cubeParts[0])
			color := cubeParts[1]

			if color == "red" {
				game.red = digit
			} else if color == "green" {
				game.green = digit
			} else if color == "blue" {
				game.blue = digit
			}
		}

		if game.red > 12 || game.green > 13 || game.blue > 14 {
			return 0, nil
		}

		fmt.Printf("%d => %d,%d,%d\n", game.id, game.red, game.green, game.blue)
	}

	return game.id, nil
}

// The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?
