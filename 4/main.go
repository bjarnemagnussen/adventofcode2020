package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type passport map[string]string

func main() {
	lines := readLines("input.txt")
	passports := make([]passport, len(lines))
	for i, l := range lines {
		passport := make(passport, 0)
		for _, els := range strings.Split(l, " ") {
			el := strings.Split(els, ":")
			passport[el[0]] = el[1]
		}
		passports[i] = passport
	}

	{
		fmt.Println("--- Part One ---")

		var count int
		for _, p := range passports {
			_, err := isValidPart1(p)
			if err == nil {
				count++
			}
		}
		fmt.Println(count)
	}

	{
		fmt.Println("--- Part Two ---")

		var count int
		for _, p := range passports {
			_, err := isValidPart2(p)
			if err == nil {
				count++
			}
		}
		fmt.Println(count)
	}
}

func isValidPart1(p passport) (bool, error) {
	var keys int
	var country bool
	for k := range p {
		switch k {
		case "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid":
			keys++
		case "cid":
			country = true
		}
	}

	var err error
	if keys < 7 {
		err = fmt.Errorf("not a valid credential")
	}

	return keys == 7 && country, err
}

func isValidPart2(p passport) (bool, error) {
	var keys int
	var country bool
	for k, v := range p {
		switch k {

		case "byr":
			year, err := strconv.Atoi(v)
			if err == nil && year >= 1920 && year <= 2002 {
				keys++
			}

		case "iyr":
			year, err := strconv.Atoi(v)
			if err == nil && year >= 2010 && year <= 2020 {
				keys++
			}

		case "eyr":
			year, err := strconv.Atoi(v)
			if err == nil && year >= 2020 && year <= 2030 {
				keys++
			}

		case "hgt":
			if strings.HasSuffix(v, "cm") {
				m, err := strconv.Atoi(strings.TrimSuffix(v, "cm"))
				if err == nil && m >= 150 && m <= 193 {
					keys++
				}
			} else if strings.HasSuffix(v, "in") {
				m, err := strconv.Atoi(strings.TrimSuffix(v, "in"))
				if err == nil && m >= 59 && m <= 76 {
					keys++
				}
			}

		case "hcl":
			if len(v) == 7 && v[0] == '#' {
				if _, err := hex.DecodeString(v[1:]); err == nil {
					keys++
				}
			}

		case "ecl":
			switch v {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
				keys++
			}

		case "pid":
			_, err := strconv.Atoi(v)
			if err == nil && len(v) == 9 {
				keys++
			}

		case "cid":
			country = true
		}
	}

	var err error
	if keys < 7 {
		err = fmt.Errorf("not a valid credential")
	}

	return keys == 7 && country, err
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	var line strings.Builder
	proceed := true
	for proceed {
		proceed = scanner.Scan()
		switch {
		case proceed == false, scanner.Text() == "":
			lines = append(lines, strings.TrimSpace(line.String()))
			line.Reset()
		default:
			line.WriteString(scanner.Text())
			line.WriteRune(' ')
		}
	}
	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
