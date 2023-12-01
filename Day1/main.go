package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Unable to read input")
	}
	lines := strings.Split(string(input), "\n")

	total := 0
	for _, line := range lines {
		firstEncountered := false
		firstDigit := ""
		lastDigit := ""
		for _, character := range line {
			if unicode.IsDigit(character) {
				if firstEncountered == false {
					firstEncountered = true
					firstDigit = string(character)
					lastDigit = string(character)
				} else {
					lastDigit = string(character)
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
