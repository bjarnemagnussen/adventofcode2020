package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(countTrees(lines, [2]int{3, 1}))
	}

	{
		fmt.Println("--- Part Two ---")
		res := 1
		directions := [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
		for _, d := range directions {
			res *= countTrees(lines, d)
		}
		fmt.Println(res)
	}
}

func countTrees(lines []string, direction [2]int) (count int) {
	right, down := direction[0], direction[1]
	offset := 0
	for i := 0; i < len(lines); i += down {
		if lines[i][offset] == '#' {
			count++
		}
		offset = (offset + right) % len(lines[i])
	}
	return
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
