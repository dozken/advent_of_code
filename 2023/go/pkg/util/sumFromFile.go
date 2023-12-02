package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type parseLine func(string) (int, error)

func SumFromFile(pathToFile string, fn parseLine) int {
	f, err := os.Open(pathToFile)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		count, err := fn(line)
		if err != nil {
			log.Fatal("Could not parse line: ", err)
		}
		sum += count
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func SumFromText(data string, fn parseLine) int {
	lines := strings.Split(data, "\n")

	sum := 0
	for _, v := range lines {
		count, err := fn(v)
		if err != nil {
			log.Fatal("Could not parse line: ", err)
		}
		sum += count
	}

	return sum
}
