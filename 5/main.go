package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")
	r := strings.NewReplacer("B", "1", "R", "1", "F", "0", "L", "0")

	ids := make([]int, len(lines))
	for i, line := range lines {
		binary := r.Replace(line)
		row, err := strconv.ParseInt(binary[:7], 2, 64)
		check(err)
		col, err := strconv.ParseInt(binary[7:], 2, 64)
		check(err)
		ids[i] = int(8*row + col)
	}
	sort.Ints(ids)

	{
		fmt.Println("--- Part One ---")
		fmt.Println(ids[len(ids)-1])
	}

	{
		fmt.Println("--- Part Two ---")
		var missing int
		for i := 0; i < len(ids)-2; i++ {
			if ids[i+1]-ids[i] == 2 {
				missing = ids[i] + 1
			}
		}
		fmt.Println(missing)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
