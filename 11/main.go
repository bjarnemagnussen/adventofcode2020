package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	floor    = '.'
	empty    = 'L'
	occupied = '#'
)

func main() {
	lines := readLines("input.txt")
	seats := make([][]rune, len(lines))
	for i, line := range lines {
		seats[i] = make([]rune, len(line))
		for j := range line {
			seats[i][j] = rune(line[j])
		}
	}

	{
		fmt.Println("--- Part One ---")

		prev := make([][]rune, len(seats))
		deepcopy(prev, seats)
		next := make([][]rune, len(prev))
		deepcopy(next, prev)

		changed := true
		used := []int{0, 0}
		for changed {
			changed = false
			used[1], used[0] = used[0], 0
			for i := range prev {
				for j := range prev[i] {
					if prev[i][j] == floor {
						continue
					}

					numSeated := inUse(prev, i, j)
					if prev[i][j] == occupied && numSeated >= 4 {
						changed = true
						next[i][j] = empty
					} else if prev[i][j] == empty && numSeated == 0 {
						changed = true
						next[i][j] = occupied
						used[0]++
					} else if prev[i][j] == occupied {
						used[0]++
					}
				}
			}
			deepcopy(prev, next)
		}

		fmt.Println(used[1])
	}
	{
		fmt.Println("--- Part Two ---")

		prev := make([][]rune, len(seats))
		deepcopy(prev, seats)
		next := make([][]rune, len(prev))
		deepcopy(next, prev)

		changed := true
		used := []int{0, 0}
		for changed {
			changed = false
			used[1], used[0] = used[0], 0
			for i := range prev {
				for j := range prev[i] {
					if prev[i][j] == floor {
						continue
					}

					numSeated := inSight(prev, i, j)
					if prev[i][j] == occupied && numSeated >= 5 {
						changed = true
						next[i][j] = empty
					} else if prev[i][j] == empty && numSeated == 0 {
						changed = true
						next[i][j] = occupied
						used[0]++
					} else if prev[i][j] == occupied {
						used[0]++
					}
				}
			}
			deepcopy(prev, next)
		}

		fmt.Println(used[1])

	}
}

func inUse(seats [][]rune, x, y int) (result int) {
	a, b, c, d := max(0, x-1), min(x+1, len(seats)-1), max(0, y-1), min(y+1, len(seats[x])-1)
	for i := a; i <= b; i++ {
		for j := c; j <= d; j++ {
			if i == x && j == y {
				continue
			}
			if seats[i][j] == occupied {
				result++
			}
		}
	}
	return result
}

func inSight(seats [][]rune, x, y int) (result int) {
NORTH:
	for i := x - 1; i >= 0; i-- {
		switch seats[i][y] {
		case occupied:
			result++
			break NORTH
		case empty:
			break NORTH
		}
	}

NORTHEAST:
	for i, j := x-1, y+1; i >= 0 && j < len(seats[x]); i, j = i-1, j+1 {
		switch seats[i][j] {
		case occupied:
			result++
			break NORTHEAST
		case empty:
			break NORTHEAST
		}
	}

EAST:
	for j := y + 1; j < len(seats[x]); j++ {
		switch seats[x][j] {
		case occupied:
			result++
			break EAST
		case empty:
			break EAST
		}
	}

SOUTHEAST:
	for i, j := x+1, y+1; i < len(seats) && j < len(seats[x]); i, j = i+1, j+1 {
		switch seats[i][j] {
		case occupied:
			result++
			break SOUTHEAST
		case empty:
			break SOUTHEAST
		}
	}

SOUTH:
	for i := x + 1; i < len(seats); i++ {
		switch seats[i][y] {
		case occupied:
			result++
			break SOUTH
		case empty:
			break SOUTH
		}
	}

SOUTHWEST:
	for i, j := x+1, y-1; i < len(seats) && j >= 0; i, j = i+1, j-1 {
		switch seats[i][j] {
		case occupied:
			result++
			break SOUTHWEST
		case empty:
			break SOUTHWEST
		}
	}

WEST:
	for j := y - 1; j >= 0; j-- {
		switch seats[x][j] {
		case occupied:
			result++
			break WEST
		case empty:
			break WEST
		}
	}

NORTHWEST:
	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		switch seats[i][j] {
		case occupied:
			result++
			break NORTHWEST
		case empty:
			break NORTHWEST
		}
	}
	return result
}

func deepcopy(dst, src [][]rune) {
	for i := range src {
		dst[i] = make([]rune, len(src[i]))
		copy(dst[i], src[i])
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

// func neighbours(i, j int, layout [][]rune) (count int) {
// 	if i-1 >= 0 && layout[i-1][j] == '#' {
// 		if j+1<=len(layout) && layout[i-1][j+1] == '#'
// 		count++
// 	}
// 	if i+1<=len(layout[0]) && layout[i+1][j] == '#' {
// 		count++
// 	}
// 	if
// }

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
