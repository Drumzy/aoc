package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Drumzy/aoc/2022/utils"
)

func main() {
	fmt.Printf("Total score Part1 would be :%d \n", solvePart1("./d2_input.txt"))
	fmt.Printf("Total score Part2 would be :%d \n", solvePart2("./d2_input.txt"))
}

var resultsP1 = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}
var resultsP2 = map[string]int{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}

func solvePart1(input string) int {
	rounds := cleanseRounds(input)
	var sum = 0
	for _, round := range rounds {
		sum += resultsP1[string(string(round[0])+" "+string(round[2]))]
	}
	return sum
}
func solvePart2(input string) int {
	rounds := cleanseRounds(input)
	var sum = 0
	for _, round := range rounds {
		sum += resultsP2[string(string(round[0])+" "+string(round[2]))]
	}
	return sum
}
func cleanseRounds(input string) []string {
	rawInput, err := os.ReadFile(input)
	utils.CheckError(err)
	rounds := strings.Split(string(rawInput), "\n")
	return rounds
}
