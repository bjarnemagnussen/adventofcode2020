package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	jolts := readNumbers("input.txt")
	jolts = append(jolts, 0)
	sort.Ints(jolts)

	{
		fmt.Println("--- Part One ---")

		var jolt1, jolt3 int
		for i := 0; i < len(jolts)-1; i++ {
			diff := jolts[i+1] - jolts[i]
			switch diff {

			case 1:
				jolt1++

			case 3:
				jolt3++

			default:
				panic(fmt.Errorf("cannot chain all joltages"))

			}
		}

		fmt.Println(jolt1 * (jolt3 + 1))
	}
	{
		fmt.Println("--- Part Two ---")

		// combinations := make([]int, len(jolts))
		// combinations[len(jolts)-1] = 1

		// for i := len(jolts) - 2; i >= 0; i-- {
		// 	combinations[i] = combinations[i+1]
		// 	if i+2 < len(jolts) && jolts[i+2]-jolts[i] <= 3 {
		// 		combinations[i] += combinations[i+2]
		// 	}
		// 	if i+3 < len(jolts) && jolts[i+3]-jolts[i] <= 3 {
		// 		combinations[i] += combinations[i+3]
		// 	}
		// }

		combinations := make([]int, 4)
		// initial values.
		combinations[2], combinations[3] = 1, 1
		combinations[1] = 1
		if jolts[len(jolts)-1]-jolts[len(jolts)-3] <= 3 {
			combinations[1] = 2
		}

		for i := len(jolts) - 4; i >= 0; i-- {
			combinations[0] = combinations[1]
			if jolts[i+2]-jolts[i] <= 3 {
				combinations[0] += combinations[2]
			}
			if jolts[i+3]-jolts[i] <= 3 {
				combinations[0] += combinations[3]
			}
			// move up by one step.
			copy(combinations[1:], combinations[:3])
		}

		fmt.Println(combinations[0])
	}

}

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, toInt(scanner.Text()))
	}
	return numbers
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
