package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")

		var mem = make(map[int]int, 0)
		var mask = make([]int, 2)
		for i := 0; i < len(lines); i++ {
			line := strings.Split(lines[i], " = ")

			if lines[i][:4] == "mask" {
				// Make two bit-masks:
				// "0-mask": convert x-to-1, then use AND operator
				// "1-mask": convert x-to-0, then use OR operator
				mask0, err := strconv.ParseInt(strings.Replace(line[1], "X", "1", -1), 2, 0)
				check(err)
				mask1, err := strconv.ParseInt(strings.Replace(line[1], "X", "0", -1), 2, 0)
				check(err)
				mask[0] = int(mask0)
				mask[1] = int(mask1)
				continue
			}

			address := toInt(strings.TrimSuffix(strings.TrimPrefix(line[0], "mem["), "]"))
			mem[address] = (toInt(line[1]) & mask[0]) | mask[1]
		}

		var sum int
		for _, v := range mem {
			if v != 0 {
				sum += v
			}
		}

		fmt.Println(sum)
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
