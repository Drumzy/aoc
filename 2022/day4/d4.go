package main

import (
	"fmt"
	"strings"

	"github.com/Drumzy/aoc/2022/utils"
)

func solvePart1(input string) (TotalOverlap int) {
	assignments := utils.CleanseInput(input)
	for _, pairs := range assignments {
		elf1p, elf2p := utils.PiarsDecoupler(strings.Split(pairs, ","))
		if elf1p.Start >= elf2p.Start && elf1p.End <= elf2p.End {
			TotalOverlap++
		} else if elf2p.Start >= elf1p.Start && elf2p.End <= elf1p.End {
			TotalOverlap++
		}
	}
	return
}
func solvePart2(input string) (TotalOverlap int) {
	assignments := utils.CleanseInput(input)
	for _, pairs := range assignments {
		elf1p, elf2p := utils.PiarsDecoupler(strings.Split(pairs, ","))
		if (elf2p.End-elf1p.Start)*(elf1p.End-elf2p.Start) >= 0 {
			TotalOverlap++
		}
	}
	return
}
func main() {
	fmt.Println("Overlap sum Part 1:", solvePart1("d4_input.txt"))
	fmt.Println("Overlap sum Part 2:", solvePart2("d4_input.txt"))
}
