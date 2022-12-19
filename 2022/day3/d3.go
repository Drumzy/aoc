package main

import (
	"fmt"
	"strings"

	"github.com/Drumzy/aoc/2022/utils"
)

func main() {
	fmt.Println(solvePart1("d3_input.txt"))
	fmt.Println(solvePart2("d3_input.txt"))
}
func solvePart1(input string) int {
	Fscore := 0
	ruckSacks := utils.CleanseInput(input)
	for _, ruckSack := range ruckSacks {
		f_half := ruckSack[len(ruckSack)/2:]
		s_half := ruckSack[:len(ruckSack)/2]
		for _, char := range f_half {
			if strings.Contains(s_half, string(char)) {
				Fscore += utils.CharToInt(char)
				break
			}
		}
	}
	return Fscore
}

func solvePart2(input string) int {
	ruckSacks := utils.CleanseInput(input)
	m := make(map[rune][3]int)
	Sscore := 0
	groupPoints := 0
	for _, ruck := range ruckSacks {
		for _, char := range ruck {
			if val, ok := m[char]; ok {
				val[groupPoints] = 1
				m[char] = val
				if val[0] == 1 && val[1] == 1 && val[2] == 1 {
					Sscore += utils.CharToInt(char)
					break
				}
			} else {
				m[char] = [3]int{0, 0, 0}
				val := m[char]
				val[groupPoints] = 1
				m[char] = val
			}
		}
		groupPoints += 1
		if groupPoints == 3 {
			groupPoints = 0
			m = make(map[rune][3]int)
		}
	}

	return Sscore
}
