package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type elfCalIntake struct {
	sum int
}

func main() {
	fmt.Printf("most carrying elf has : %d \n", solvePart1("./d1_input.txt"))
	fmt.Printf("Top three elfs has : %d", solvePart2("./d1_input.txt"))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func solvePart1(inputPath string) int {
	cleanElfs := cleanseElfsCalories(inputPath)
	var max int = 0
	for i := 0; i < len(cleanElfs); i++ {
		if max < cleanElfs[i].sum {
			max = cleanElfs[i].sum
		}
	}
	return max
}

func solvePart2(inputPath string) int {
	cleanElfs := cleanseElfsCalories(inputPath)
	firstMax := 0
	secondMax := 0
	thirdMax := 0

	for i := 0; i < len(cleanElfs); i++ {
		if firstMax < cleanElfs[i].sum {
			thirdMax = secondMax
			secondMax = firstMax
			firstMax = cleanElfs[i].sum
		} else if secondMax < cleanElfs[i].sum {
			thirdMax = secondMax
			secondMax = cleanElfs[i].sum
		} else if thirdMax < cleanElfs[i].sum {
			thirdMax = cleanElfs[i].sum
		}
	}

	return (firstMax + secondMax + thirdMax)
}

func cleanseElfsCalories(inputPath string) []elfCalIntake {
	rawInput, err := os.ReadFile(inputPath)
	checkError(err)
	elfs := strings.Split(string(rawInput), "\n")
	return splitByCalories(elfs)
}
func splitByCalories(arr []string) []elfCalIntake {
	var elfsCals []elfCalIntake
	var elfCalories elfCalIntake = elfCalIntake{sum: 0}
	for i := 0; i < len(arr); i++ {
		if arr[i] != "\r" {
			intake, err := strconv.Atoi(strings.TrimSpace(arr[i]))
			checkError(err)
			elfCalories.sum += intake
		} else {
			elfsCals = append(elfsCals, elfCalories)
			elfCalories.sum = 0
		}
	}
	return elfsCals
}
