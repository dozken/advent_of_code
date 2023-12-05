package day4

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1(line string) (int, error) {
	card := line[strings.Index(line, ": ")+2:]
	re := regexp.MustCompile(`\s+`)
	card = re.ReplaceAllString(card, " ")

	parts := strings.Split(card, " | ")

	winningNumbers := strings.Split(parts[0], " ")
	cardNumbers := strings.Split(parts[1], " ")

	winningSet := map[string]bool{}
	for _, v := range winningNumbers {
		winningSet[strings.Trim(v, " ")] = true
	}

	var myNumbers []int
	for _, v := range cardNumbers {
		_, ok := winningSet[strings.Trim(v, " ")]
		// fmt.Printf("%v %v \n", v, ok)
		if ok {
			num, _ := strconv.Atoi(v)
			myNumbers = append(myNumbers, num)
		}
	}
	// fmt.Printf("%d \n", myNumbers)

	result := 0
	for i := range myNumbers {
		if i > 0 {
			result *= (2)
		} else {
			result = 1
		}
	}

	// fmt.Println(result)

	return result, nil
}

func Part2(line string) int {
	// id := line[strings.Index(line, " "): strings.Index(line, ":")]
	// fmt.Println(id)
	card := line[strings.Index(line, ": ")+2:]
	re := regexp.MustCompile(`\s+`)
	card = re.ReplaceAllString(card, " ")

	parts := strings.Split(card, " | ")

	winningNumbers := strings.Split(parts[0], " ")
	cardNumbers := strings.Split(parts[1], " ")

	winningSet := map[string]bool{}
	for _, v := range winningNumbers {
		winningSet[strings.Trim(v, " ")] = true
	}

	var myNumbers []int
	for _, v := range cardNumbers {
		_, ok := winningSet[strings.Trim(v, " ")]
		// fmt.Printf("%v %v \n", v, ok)
		if ok {
			num, _ := strconv.Atoi(v)
			myNumbers = append(myNumbers, num)
		}
	}
	// fmt.Printf("%d \n", myNumbers)

	result := 0
	for range myNumbers {
		result++
	}

	// fmt.Println(result)

	return result
}

type parseLine func(string) int

func SumFromFile(pathToFile string) int {
	f, _ := os.Open(pathToFile)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	cardMap := make(map[int]int)
	sum := 0
	cardId := 1
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("%d: ", cardId)
		count := Part2(line)

		for i := cardId; i <= cardId+count; i++ {
			if i != cardId {
				cardMap[i] += cardMap[cardId]
			} else {
				cardMap[i] += 1
			}
		}

		// fmt.Printf("->%v\n", cardMap)
		cardId++
	}

	for _, v := range cardMap {
		sum += v
	}

	// fmt.Printf("%v", cardMap)

	return sum
}
