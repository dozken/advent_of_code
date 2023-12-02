package day1

import (
	"fmt"
	"strconv"
	"strings"
)

type parseLine func(string) (int, error)

func ParseLinePart1(line string) (int, error) {
	chars := strings.Split(line, "")
	firstDigit, secondDigit := -1, -1

	for i, j := 0, len(chars)-1; j > -1; i, j = i+1, j-1 {
		if firstDigit == -1 {
			if l, err := strconv.Atoi(chars[i]); err == nil {
				firstDigit = l
			}
		}

		if secondDigit == -1 {
			if r, err := strconv.Atoi(chars[j]); err == nil {
				secondDigit = r
			}
		}

		if firstDigit != -1 && secondDigit != -1 {
			// log.Printf("First: %d, Second: %d", firstDigit, secondDigit)
			return firstDigit*10 + secondDigit, nil
		}

	}

	return -1, fmt.Errorf("Could not Max in Neovim")
}

func ParseLinePart2(line string) (int, error) {
	digits := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for k, v := range digits {
		line = strings.ReplaceAll(line, k, v)
	}

	chars := strings.Split(line, "")
	firstDigit, secondDigit := -1, -1
	for i, j := 0, len(chars)-1; j > -1; i, j = i+1, j-1 {

		if firstDigit == -1 {
			if l, err := strconv.Atoi(chars[i]); err == nil {
				firstDigit = l
			}
		}

		if secondDigit == -1 {
			if r, err := strconv.Atoi(chars[j]); err == nil {
				secondDigit = r
			}
		}

		if firstDigit != -1 && secondDigit != -1 {
			return firstDigit*10 + secondDigit, nil
		}

	}

	return -1, fmt.Errorf("Could not find two numbers")
}
