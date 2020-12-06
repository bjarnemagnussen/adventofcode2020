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

		var countAny int
		var answers [26]int
		for _, line := range lines {
			if line == "" {
				// reset
				answers = [26]int{}
				continue
			}

			for _, ch := range line {
				countAny += (1 - answers[ch-'a'])
				answers[ch-'a'] = 1
			}
		}

		fmt.Println(countAny)
	}

	{
		fmt.Println("--- Part Two ---")

		var countAll, people int
		var answers [26]int
		for i, line := range lines {
			people++
			if line == "" {
				// reset
				answers, people = [26]int{}, 0
				continue
			}

			for _, ch := range line {
				answers[ch-'a']++
				if (i == len(lines)-1 || lines[i+1] == "") && answers[ch-'a'] == people {
					countAll++
				}
			}
		}

		fmt.Println(countAll)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
