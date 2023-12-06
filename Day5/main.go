package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	startIn  int64
	startOut int64
	steps    int64
}

var exampleInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func main() {
	file, _ := os.ReadFile("./input.txt")
	inputString := string(file)

	resChallenge1 := challenge1(inputString)
	resChallenge2 := challenge2(inputString)

	fmt.Printf("Challenge 1 %d\n", resChallenge1)
	fmt.Printf("Challenge 2 %d\n", resChallenge2)
}

func challenge1(input string) int64 {
	mapsAndSeeds := strings.Split(input, "\n\n")
	seeds := strings.Split(strings.Split(mapsAndSeeds[0], ": ")[1], " ")
	maps := createMapsSlice(mapsAndSeeds[1:])
	locations := []int64{}
	for i := 0; i < len(seeds); i++ {
		seed, _ := strconv.Atoi(seeds[i])
		seedCopy := int64(seed)
		for j := 0; j < 7; j++ {
			seedCopy = getValFromIntervalMaps(seedCopy, maps[j])
		}
		locations = append(locations, seedCopy)
	}
	sort.Slice(locations, func(i, j int) bool { return locations[i] < locations[j] })
	return locations[0]
}

func challenge2(input string) int64 {
	mapsAndSeeds := strings.Split(input, "\n\n")
	seeds := strings.Split(strings.Split(mapsAndSeeds[0], ": ")[1], " ")
	maps := createMapsSlice(mapsAndSeeds[1:])
	locations := []int64{}
	for i := 0; i < len(seeds)-1; i = i + 2 {
		seedStart, _ := strconv.Atoi(seeds[i])
		length, _ := strconv.Atoi(seeds[i+1])

		seedStartCopy := int64(seedStart)
		lengthCopy := int64(length)
		// fmt.Print64ln(seedStart, length)
		for seed := seedStartCopy; seed < seedStartCopy+lengthCopy; seed++ {
			seedCopy := seed
			for j := 0; j < 7; j++ {
				seedCopy = getValFromIntervalMaps(seedCopy, maps[j])
			}
			locations = append(locations, seedCopy)
		}

	}
	sort.Slice(locations, func(i, j int) bool { return locations[i] < locations[j] })
	return locations[0]
}

func getValFromIntervalMaps(seed int64, intervals []Interval) int64 {
	for _, int64erval := range intervals {
		if seed >= int64erval.startIn && seed < int64erval.startIn+int64erval.steps-1 {
			return int64erval.startOut + (seed - int64erval.startIn)
		}
	}
	return seed
}

func createMapsSlice(input []string) [][]Interval {
	res := make([][]Interval, 0)
	for _, mapRepresentation := range input {
		data := strings.Split(mapRepresentation, ":\n")[1:]
		intervals := make([]Interval, 0)
		for _, block := range data {
			lines := strings.Split(block, "\n")

			for _, line := range lines {
				numbers := strings.Split(line, " ")
				startIn, _ := strconv.Atoi(numbers[1])
				startOut, _ := strconv.Atoi(numbers[0])
				steps, _ := strconv.Atoi(numbers[2])

				startInCopy := int64(startIn)
				startOutCopy := int64(startOut)
				stepsCopy := int64(steps)
				intervals = append(intervals, Interval{startIn: startInCopy, startOut: startOutCopy, steps: stepsCopy})
			}
		}
		res = append(res, intervals)
	}
	return res
}
