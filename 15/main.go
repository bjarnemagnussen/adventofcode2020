package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := make(map[int]int, 0)
	var last int
	for scanner.Scan() {
		for i, s := range strings.Split(scanner.Text(), ",") {
			input[toInt(s)] = i + 1
			last = toInt(s)
		}
	}

	{
		fmt.Println("--- Part One ---")

		last := last
		numbers := make(map[int]int)
		for k, v := range input {
			numbers[k] = v
		}
		var out int
		for turn := len(numbers); turn < 2020; turn++ {
			out = 0
			if v, ok := numbers[last]; ok {
				out = turn - v
			}
			numbers[last] = turn
			last = out
		}

		fmt.Println(out)
	}

	{
		fmt.Println("--- Part Two ---")

		last := last
		numbers := make(map[int]int)
		for k, v := range input {
			numbers[k] = v
		}
		var out int
		for turn := len(numbers); turn < 30000000; turn++ {
			out = 0
			if v, ok := numbers[last]; ok {
				out = turn - v
			}
			numbers[last] = turn
			last = out
		}

		fmt.Println(out)
	}
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
