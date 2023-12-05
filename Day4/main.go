package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

var exampleInput string = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53  5 44 | 69 82  6 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func main() {

	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal("Unable to parse input file")
	}
	dataArray := strings.Split(string(file), "\n")
	resChallenge1 := 0
	for _, line := range dataArray {
		winningNums, acquiredNums := parseInput(line)
		resChallenge1 += calculatePoints(winningNums, acquiredNums)
	}
	cards := make(map[int]int)
	i := 0
	resChallenge2 := 0
	for i < len(dataArray) {
		winningNums, acquiredNums := parseInput(dataArray[i])
		numMatches := getNumberOfMatches(winningNums, acquiredNums)
		resChallenge2++
		card := i + 1
		cards[card]++
		count := cards[card]
		for x := 1; x <= numMatches; x++ {
			cards[card+x] += count
			resChallenge2 += count
		}
		i++
	}

	fmt.Printf("Challenge 1 result: %d\n", resChallenge1)
	fmt.Printf("Challenge 2 result: %d\n", resChallenge2)

}

func parseInput(line string) ([]string, []string) {

	data := strings.Split(line, ":")
	gameData := strings.Split(data[1], "|")
	winningNumsStr, acquiredNumsStr := gameData[0], gameData[1]
	winningNumbers := strings.Split(strings.Trim(winningNumsStr, " "), " ")
	acquiredNumbers := strings.Split(strings.Trim(acquiredNumsStr, " "), " ")

	return winningNumbers, acquiredNumbers
}

func calculatePoints(winningNums, acquiredNums []string) int {
	matches := 0
	for _, val := range winningNums {
		if val == "" {
			continue
		}
		if slices.Contains(acquiredNums, val) {
			matches += 1
		}
	}
	return int(math.Pow(2, float64(matches-1)))
}

func getNumberOfMatches(winningNums, acquiredNums []string) int {
	matches := 0
	for _, val := range winningNums {
		if val == "" {
			continue
		}
		if slices.Contains(acquiredNums, val) {
			matches += 1
		}
	}
	return matches
}

func insert(a []string, index int, value string) []string {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
