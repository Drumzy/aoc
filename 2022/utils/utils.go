package utils

import (
	"os"
	"strconv"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CharToInt(c rune) int {
	if int(c) > 90 {
		return int(c) - 96
	} else {
		return int(c) - 38
	}
}

func CleanseInput(input string) []string {
	rawInput, err := os.ReadFile(input)
	CheckError(err)
	return strings.Split(strings.TrimSpace(string(rawInput)), "\n")
}

func IntParser(pair []string) (start int, end int) {
	start, _ = strconv.Atoi(strings.TrimSpace(pair[0]))
	end, _ = strconv.Atoi(strings.TrimSpace(pair[1]))

	return
}

// Day4
type ElfPair struct {
	Start int
	End   int
}

func PiarsDecoupler(pair []string) (ElfPair, ElfPair) {
	fElfS, fElfE := IntParser(strings.Split(pair[0], "-"))
	SElfS, SElfE := IntParser(strings.Split(pair[1], "-"))
	return ElfPair{Start: fElfS, End: fElfE}, ElfPair{Start: SElfS, End: SElfE}
}
