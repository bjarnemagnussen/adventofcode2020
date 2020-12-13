package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input_test.txt")

	arrival := toInt(lines[0])

	{
		fmt.Println("--- Part One ---")

		var earliest string
		var departure int
		for departure = arrival; earliest == ""; departure++ {
			for _, id := range strings.Split(lines[1], ",") {
				if id == "x" {
					continue
				}

				if departure%toInt(id) == 0 {
					earliest = id
					break
				}
			}
		}

		fmt.Println((departure - 1 - arrival) * toInt(earliest))
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
