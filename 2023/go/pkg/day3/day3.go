package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type NumberCoordinate struct {
	xLeft  int
	xRight int
	y      int
	number int
}

func SumFromFile(filePath string) int {
	result := 0
	content, _ := os.ReadFile(filePath)
	array2D := create2DArray(string(content))
	directions := []Coordinate{
		{-1, 1}, {0, 1}, {1, 1},
		{-1, 0}, {1, 0},
		{-1, -1}, {0, -1}, {1, -1},
	}
	var seenCoordinates []NumberCoordinate

	for y := 0; y < len(array2D)-1; y++ {
		for x := 0; x < len(array2D[y])-1; x++ {

			if isSymbol(array2D[y][x]) {
				// fmt.Printf("%s", array2D[y][x])
				for _, dir := range directions {
					yCoor := y + dir.y
					xCoor := x + dir.x
					position := array2D[yCoor][xCoor]

					if isDigit(position) && !isSeen(seenCoordinates, xCoor, yCoor) {
						// fmt.Printf("%s [%d, %d]", position, xCoor, yCoor)
						num := position

						xLeft := xCoor - 1
						for ; xLeft >= 0 && isDigit(array2D[yCoor][xLeft]); xLeft-- {
							num = array2D[yCoor][xLeft] + num
						}
						xRight := xCoor + 1
						for ; xRight < len(array2D)-1 && isDigit(array2D[yCoor][xRight]); xRight++ {
							num = num + array2D[yCoor][xRight]
						}

						number, _ := strconv.Atoi(num)
						seenCoordinates = append(seenCoordinates, NumberCoordinate{xLeft: xLeft + 1, xRight: xRight - 1, y: yCoor, number: number})
						// fmt.Printf(">%s<", num)
					}
				}
			}
			fmt.Printf("%s", array2D[y][x])
		}
		fmt.Println()
	}

	// log.Printf("\nSeen: %v\n", seenCoordinates)
	for _, nc := range seenCoordinates {
		result += nc.number
	}

	return result
}

func isSeen(list []NumberCoordinate, x int, y int) bool {
	for _, nc := range list {
		if nc.y == y && nc.xLeft <= x && x <= nc.xRight {
			return true
		}
	}
	return false
}

func isSymbol(char string) bool {
	isPeriod := char == "."
	return !isDigit(char) && !isPeriod
}

func isDigit(char string) bool {
	return "0" <= char && char <= "9"
}

func create2DArray(fileContent string) [][]string {
	lines := strings.Split(fileContent, "\n")
	var array [][]string
	for _, line := range lines {
		if len(line) >= 0 {
			array = append(array, strings.Split(line, ""))
		}
	}
	return array
}
