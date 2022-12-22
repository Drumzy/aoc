package main

import (
	"fmt"
	"os"

	"github.com/Drumzy/aoc/2022/utils"
)

func ContainsDuplicates(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}
func part1(input string) int {
	for i := 0; i+4 <= len(input); i++ {
		if ContainsDuplicates(input[i : i+4]) {
			return i + 4
		}
	}

	return -1
}
func part2(input string) int {
	for i := 0; i+14 <= len(input); i++ {
		if ContainsDuplicates(input[i : i+14]) {
			return i + 14
		}
	}

	return -1
}
func main() {
	rawInput, err := os.ReadFile("d6_input.txt")
	utils.CheckError(err)
	input := string(rawInput)
	fmt.Println("Part 1 Answer : ", part1(input))
	fmt.Println("Part 2 Answer : ", part2(input))
}
