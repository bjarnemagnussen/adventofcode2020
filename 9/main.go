package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	// size denotes the number of elements in the preamble.
	size int = 25
	// size = 5 // For test data-set
)

func main() {
	numbers := readNumbers("input.txt")

	var target, idx int

	{
		fmt.Println("--- Part One ---")

		preamble := make([]int, size)
		for i := len(preamble); i < len(numbers); i++ {
			copy(preamble, numbers[i-len(preamble):i])
			sort.Ints(preamble)

			if ok := twoPointer(preamble, numbers[i]); !ok {
				target, idx = numbers[i], i
				break
			}
		}

		fmt.Println(target)
	}

	{
		fmt.Println("--- Part Two ---")

		start, end, sum := 0, 1, numbers[0]+numbers[1]
		for start < idx-1 {
			if sum == target {
				break
			} else if sum > target {
				start++
				sum -= numbers[start-1]
			} else {
				end++
				sum += numbers[end]
			}
		}

		mini, maxi := numbers[start], numbers[start]
		for i := start + 1; i <= end; i++ {
			mini, maxi = min(mini, numbers[i]), max(maxi, numbers[i])
		}

		fmt.Println(mini + maxi)
	}

}

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
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

// twoPointer finds a touple in nums that matches target using the two-pointer
// technique.
func twoPointer(nums []int, target int) bool {
	i, j := 0, len(nums)-1
	for i < j {
		switch {
		case nums[i]+nums[j] > target:
			j--
		case nums[i]+nums[j] < target:
			i++
		default:
			return true
		}
	}
	return false
}
