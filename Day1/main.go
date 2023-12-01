package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var integersWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func nextInt(s string, firstFound bool) string {
	for i, number := range integersWords {
		if strings.Index(s, number) == 0 {
			return strconv.Itoa(i + 1)
		}
	}
	return "-1"
}

func main() {
	input, err := os.ReadFile("./input2.txt")
	if err != nil {
		log.Fatal("Unable to read input")
	}
	lines := strings.Split(string(input), "\n")

	total := 0
	for _, line := range lines {
		firstEncountered := false
		firstDigit := ""
		lastDigit := ""
		for i, character := range line {
			if unicode.IsDigit(character) {
				if firstEncountered == false {
					firstEncountered = true
					firstDigit = string(character)
					lastDigit = string(character)
				} else {
					lastDigit = string(character)
				}
			} else {
				dig := nextInt(line[i:], firstEncountered)
				if dig != "-1" {
					if firstEncountered == false {
						firstEncountered = true
						firstDigit = dig
						lastDigit = dig
					} else {
						lastDigit = dig
					}
				}
			}
		}
		currentNumber, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			log.Fatal("Unable to parse number from line")
		}
		total += currentNumber
	}
	fmt.Println(total)
}
