package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op  string
	val int
}

func main() {
	lines := readLines("input.txt")
	program := make([]instruction, len(lines))
	for i := 0; i < len(program); i++ {
		s := strings.Split(lines[i], " ")
		op, val := s[0], toInt(s[1])
		program[i] = instruction{op, val}
	}

	{
		fmt.Println("--- Part One ---")
		acc, _ := run(program)
		fmt.Println(acc)
	}

	{
		fmt.Println("--- Part Two ---")

		var acc int
		var fixed bool

		for i := range program {
			fix := make([]instruction, len(program))
			copy(fix, program)
			switch program[i].op {

			case "nop":
				fix[i].op = "jmp"

			case "jmp":
				fix[i].op = "nop"

			default:
				continue

			}

			acc, fixed = run(fix)
			if fixed {
				break
			}
		}

		fmt.Println(acc)
	}
}

func run(program []instruction) (int, bool) {
	visited := make([]bool, len(program))
	halts := true

	var acc, cur int
	for cur < len(program) {
		if visited[cur] {
			halts = false
			break
		}

		visited[cur] = true
		p := program[cur]

		switch p.op {
		case "nop":
			cur++
		case "acc":
			acc += p.val
			cur++
		case "jmp":
			cur += p.val
		}
	}

	return acc, halts
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
