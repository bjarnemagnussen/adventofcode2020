package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ross-oreto/go-tree"

	"bufio"
	"os"
)

func main() {
	lines := readLines("input.txt")

	btree := tree.New()

	tickets := make([][]int, 0)
	notes := make(map[string][]string, 0)

	section := 0
	for _, line := range lines {
		if line == "" {
			section++
			continue
		} else if line == "your ticket:" || line == "nearby tickets:" {
			continue
		}

		switch section {

		case 0:
			col := strings.Split(line, ": ")
			vals := strings.Split(col[1], " or ")
			notes[col[0]] = vals
			for _, v := range vals {
				mini, maxi := toInt(strings.Split(v, "-")[0]), toInt(strings.Split(v, "-")[1])
				for i := mini; i <= maxi; i++ {
					btree.Insert(tree.IntVal(i))
				}
			}

		default:
			vals := strings.Split(line, ",")
			t := make([]int, len(vals))
			for i, val := range vals {
				t[i] = toInt(val)
			}
			tickets = append(tickets, t)
		}

	}

	{
		fmt.Println("--- Part One ---")

		var sum int
		n := 1
		for _, ticket := range tickets[1:] {
			var delete bool
			for j := range ticket {
				if !btree.Contains(tree.IntVal(ticket[j])) {
					sum += ticket[j]
					delete = true
				}
			}
			if !delete {
				tickets[n] = ticket
				n++
			}
		}
		tickets = tickets[:n]

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
