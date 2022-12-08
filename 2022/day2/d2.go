package main

import "os"

func main() {

}

type Round struct {
	player1 string
	player2 string
}

func cleanseRounds(input string) []Round {
	rawInput, err := os.ReadFile(input)
	utils.checkError(err)
}
func splitByRound(arr []string) {

}
