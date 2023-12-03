package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	start  int
	finish int
	value  int
	row    int
}

type SpecialChar struct {
	row   int
	col   int
	value string
}

var adjacents [][]int = [][]int{
	{1, 0},
	{-1, 0},
	{1, 1},
	{1, -1},
	{0, 1},
	{0, -1},
	{-1, 1},
	{-1, -1},
}
var testInput string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func main() {

	file, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("Unabel to parse the input file")
	}
	challenge1Res := challenge1(string(file))
	fmt.Printf("Challenge 1 result: %d\n", challenge1Res)

	challenge2Res := challenge2(string(file))
	fmt.Printf("Challenge 2 result: %d\n", challenge2Res)

}

func challenge1(input string) int {
	lines := strings.Split(input, "\n")
	numPositions, specialCharsPositions := getPositions(lines)
	visited := make([]Number, 0)
	res := 0
	for _, char := range specialCharsPositions {
		for _, pos := range adjacents {
			dx, dy := pos[0], pos[1]
			adjX, adjY := char.row+dx, char.col+dy

			for _, number := range numPositions {
				seen := slices.ContainsFunc(visited, func(s Number) bool { return s.row == number.row && s.finish == number.finish })
				if number.row == adjX && adjY >= number.start && adjY <= number.finish && !seen {
					visited = append(visited, number)
					res += number.value
				}
			}
		}
	}

	return res
}

func challenge2(input string) int {
	lines := strings.Split(input, "\n")
	numPositions, specialCharsPositions := getPositions(lines)

	res := 0
	for _, char := range specialCharsPositions {
		if char.value != "*" {
			continue
		}
		visited := make([]Number, 0)
		gearRatio := 1
		numAdjacents := 0
		for _, pos := range adjacents {
			dx, dy := pos[0], pos[1]
			adjX, adjY := char.row+dx, char.col+dy

			for _, number := range numPositions {
				seen := slices.ContainsFunc(visited, func(s Number) bool { return s.row == number.row && s.finish == number.finish })
				if number.row == adjX && adjY >= number.start && adjY <= number.finish && !seen {
					numAdjacents += 1
					visited = append(visited, number)
					gearRatio *= number.value
				}
			}
		}
		if numAdjacents == 2 {
			res += gearRatio
		}
	}

	return res
}

func getPositions(lines []string) ([]Number, []SpecialChar) {
	numbers := make([]Number, 0)
	chars := make([]SpecialChar, 0)

	for i, line := range lines {
		col := 0
		for col < len(line) {
			if unicode.IsDigit(rune(line[col])) {
				num := ""
				start := col
				for col < len(line) && unicode.IsDigit(rune(line[col])) {
					num += string(line[col])
					col++
				}
				val, _ := strconv.Atoi(num)
				number := Number{start: start, finish: col - 1, value: val, row: i}
				numbers = append(numbers, number)
			} else if string(line[col]) != "." {
				specialChar := SpecialChar{row: i, col: col, value: string(line[col])}
				chars = append(chars, specialChar)
				col++
			} else {
				col++
			}
		}
	}
	return numbers, chars
}
