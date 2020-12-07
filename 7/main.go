package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type suitcase struct {
	num  int
	name string
}

func main() {
	var regexBag = regexp.MustCompile(`([0-9]+) ([a-z ]+) bag`)
	containsBags := make(map[string][]suitcase, 0)
	outerBags := make(map[string][]string, 0)

	for _, line := range readLines("input.txt") {
		splittedLine := strings.Split(line, "contain")
		outer := strings.TrimSpace(strings.Replace(splittedLine[0], "bags", "", 1))
		contains := regexBag.FindAllStringSubmatch(splittedLine[1], -1)

		for _, c := range contains {
			d, err := strconv.Atoi(c[1])
			check(err)
			s := suitcase{d, c[2]}

			containsBags[outer] = append(containsBags[outer], s)
			outerBags[s.name] = append(outerBags[s.name], outer)
		}
	}

	{
		fmt.Println("--- Part One ---")

		var count int
		used := make(map[string]bool)

		bagsToCheck := []string{"shiny gold"}
		for len(bagsToCheck) > 0 {
			var bag string
			bag, bagsToCheck = bagsToCheck[0], bagsToCheck[1:]
			bagsToCheck = append(bagsToCheck, outerBags[bag]...)

			for _, b := range outerBags[bag] {
				if !used[b] {
					count++
					used[b] = true
				}
			}
		}

		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")

		var count int
		bagsToAdd := containsBags["shiny gold"]

		for len(bagsToAdd) > 0 {
			bag := bagsToAdd[0]
			bagsToAdd = bagsToAdd[1:]
			count += bag.num

			for _, b := range containsBags[bag.name] {
				bagsToAdd = append(bagsToAdd, suitcase{bag.num * b.num, b.name})
			}
		}

		fmt.Println(count)
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
