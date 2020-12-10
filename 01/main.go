package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	nums := readNumbers("input.txt")
	sort.Ints(nums)

	{
		fmt.Println("--- Part One ---")
		fmt.Println(twoPointer(nums, 2020))
	}

	{
		fmt.Println("--- Part Two ---")
		var solution int

		// Fix k'th integer and use two-pointer to find a touple that matches.
		for k := range nums {
			if ret := twoPointer(nums[k+1:], 2020-nums[k]); ret > 0 {
				solution = nums[k] * ret
				break
			}
		}

		fmt.Println(solution)
	}

}

// twoPointer finds a touple in nums that mathces target using the two-pointer
// technique.
func twoPointer(nums []int, target int) (sol int) {
	i, j := 0, len(nums)-1
	for i < j {
		switch {
		case nums[i]+nums[j] > target:
			j--
		case nums[i]+nums[j] < target:
			i++
		default:
			return nums[i] * nums[j]
		}
	}
	return
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
