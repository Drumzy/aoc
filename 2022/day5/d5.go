package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Drumzy/aoc/2022/utils"
)

type Stack []rune

func (s *Stack) Pop(n int) Stack {
	tmp := *s
	popedElements := tmp[len(tmp)-n:]
	*s = tmp[:len(tmp)-n]
	return popedElements
}

func (s *Stack) Append(n Stack) {
	tmp := *s
	tmp = append(tmp, n...)
	*s = tmp
}
func (s *Stack) Prepend(n Stack) {
	tmp := *s
	tmp = append(n, tmp...)
	*s = tmp
}
func (s *Stack) Reverse() {
	r := *s
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

type Instructions struct {
	Steps int
	From  int
	To    int
}

func transpose(slice []Stack) []Stack {
	xl := len(slice[0])
	yl := len(slice)
	result := make([]Stack, xl)
	for i := range result {
		result[i] = make(Stack, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[j][xl-1-i] = slice[i][j]
		}
	}
	var res2 []Stack
	var v_tmp Stack
	for _, v := range result {
		for i := 0; i < len(v); i++ {
			if v[i] != ' ' {
				v_tmp = append(v_tmp, v[i])
			}
		}
		res2 = append(res2, v_tmp)
		v_tmp = nil
	}
	return res2
}
func ProcessCrates(rawCrates string) []Stack {
	var tmp_stack Stack
	var tmp_stacks []Stack
	for _, v := range strings.Split(rawCrates, "\n") {
		for i := 1; i < len(v); i += 4 {
			tmp_stack = append(tmp_stack, rune(v[i]))
		}
		tmp_stacks = append(tmp_stacks, tmp_stack)
		tmp_stack = nil
	}
	stacks := transpose(tmp_stacks)
	return stacks
}
func ProcessInstuctions(rawInstructions string) ([]Instructions, error) {
	var rest string
	var steps, from, to int
	var CleanInstruction []Instructions
	for _, rawInstruction := range strings.Split(rawInstructions, "\n") {
		_, err := fmt.Sscanf(rawInstruction, "%s %d %s %d %s %d", &rest, &steps, &rest, &from, &rest, &to)
		if err != nil {
			return []Instructions{}, fmt.Errorf("wrong format")
		}
		CleanInstruction = append(CleanInstruction, Instructions{Steps: steps, From: from, To: to})
	}
	return CleanInstruction, nil
}

type Cargo struct {
	Crates       []Stack
	Instructions Instructions
}

func Rearrange(s, t *Stack, a int) {
	c := s.Pop(a)
	c.Reverse() // coment this line for part 2
	t.Append(c)
}

func parseInput(input string) string {
	rawInput, err := os.ReadFile(input)
	utils.CheckError(err)
	result := regexp.MustCompile(`\n\s*\n`).Split(string(rawInput), -1)
	drawing := result[0]
	instructions := result[1]

	crates := ProcessCrates(drawing)
	cleanInstructions, err := ProcessInstuctions(instructions)
	utils.CheckError(err)
	for _, v := range cleanInstructions {
		Rearrange(&crates[v.From-1], &crates[v.To-1], v.Steps)
	}
	var as strings.Builder
	for _, v := range crates {
		as.WriteRune(v[len(v)-1])
	}

	return as.String()
}
func main() {
	fmt.Println(parseInput("d5_input.txt"))
}
