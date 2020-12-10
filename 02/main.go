package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	limits   [2]int
	char     string
	password string
}

func main() {
	input := make([]entry, 0)

	lines := readLines("input.txt")
	for _, l := range lines {
		entries := strings.Split(strings.Replace(l, ":", "", -1), " ")

		start := toInt(strings.Split(entries[0], "-")[0])
		end := toInt(strings.Split(entries[0], "-")[1])

		input = append(input, entry{[2]int{start, end}, entries[1], entries[2]})
	}

	{
		fmt.Println("--- Part One ---")

		var valid int
		for _, in := range input {
			count := strings.Count(in.password, in.char)
			if count >= in.limits[0] && count <= in.limits[1] {
				valid++
			}
		}

		fmt.Println(valid)
	}

	{
		fmt.Println("--- Part Two ---")

		var valid int
		for _, in := range input {
			first, second := in.limits[0]-1, in.limits[1]-1
			if (string(in.password[first]) != in.char) != (string(in.password[second]) != in.char) {
				valid++
			}
		}

		fmt.Println(valid)
	}
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
