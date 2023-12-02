package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maximums map[string]int = map[string]int{"blue": 14, "green": 13, "red": 12}

func possibleLine(line string) int {
	lineSplit := strings.Split(line, ":")
	gameId, err := strconv.Atoi(strings.Split(lineSplit[0], " ")[1])

	if err != nil {
		log.Fatal("Unable to parse game Id")
	}

	sets := strings.Split(lineSplit[1], ";")

	for _, set := range sets {
		ballColors := strings.Split(set, ",")
		for _, color := range ballColors {
			colorAndNumber := strings.Split(color, " ")
			color := colorAndNumber[2]
			number, err := strconv.Atoi(colorAndNumber[1])
			if err != nil {
				log.Fatal("Error while parsing file")
			}
			if number > maximums[color] {
				return 0
			}
		}
	}
	return gameId
}

func minimumSet(line string) int {

	lineSplit := strings.Split(line, ":")

	sets := strings.Split(lineSplit[1], ";")
	blue, green, red := 0, 0, 0
	for _, set := range sets {
		ballColors := strings.Split(set, ",")
		for _, color := range ballColors {
			colorAndNumber := strings.Split(color, " ")
			color := colorAndNumber[2]
			number, err := strconv.Atoi(colorAndNumber[1])
			if err != nil {
				log.Fatal("Error while parsing file")
			}
			switch color {
			case "blue":
				blue = max(blue, number)
			case "red":
				red = max(red, number)
			case "green":
				green = max(green, number)
			}
		}
	}
	return red * green * blue
}

func main() {

	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal("Unable to read the input file")
	}

	lines := strings.Split(string(file), "\n")

	challenge1Res := 0
	challenge2Res := 0
	for _, line := range lines {
		challenge1Res += possibleLine(line)
		challenge2Res += minimumSet(line)
	}
	fmt.Println(challenge1Res)
	fmt.Println(challenge2Res)
}
