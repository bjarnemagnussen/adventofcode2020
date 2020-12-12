package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type operation struct {
	op    string
	steps int
}

type direction struct {
	x int
	y int
}

var degreeToDir = map[int]direction{
	0:   north,
	90:  east,
	180: south,
	270: west,
}

var dirToDegree = map[direction]int{
	north: 0,
	east:  90,
	south: 180,
	west:  270,
}

func (d direction) mul(a int) direction {
	return direction{a * d.x, a * d.y}
}

func (d direction) add(d2 direction) direction {
	return direction{d.x + d2.x, d.y + d2.y}
}

func (d direction) changeRight(degree int) direction {
	// New direction contributed by the x-axis.
	xNewDegree := (dirToDegree[direction{sign(d.x), 0}] + degree) % 360
	if xNewDegree < 0 {
		xNewDegree += 360
	}
	xNewDir := direction{abs(d.x) * degreeToDir[xNewDegree].x, abs(d.x) * degreeToDir[xNewDegree].y}

	// New direction contributed by the y-axis.
	yNewDegree := (dirToDegree[direction{0, sign(d.y)}] + degree) % 360
	if yNewDegree < 0 {
		yNewDegree += 360
	}
	yNewDir := direction{abs(d.y) * degreeToDir[yNewDegree].x, abs(d.y) * degreeToDir[yNewDegree].y}

	return xNewDir.add(yNewDir)
}

func (d direction) changeLeft(degree int) direction {
	return d.changeRight(-degree)
}

var (
	north = direction{-1, 0}
	east  = direction{0, 1}
	south = direction{1, 0}
	west  = direction{0, -1}
)

func main() {
	lines := readLines("input.txt")

	commands := make([]operation, len(lines))
	for i, line := range lines {
		cmd := string(line[0])
		steps, err := strconv.Atoi(line[1:])
		check(err)
		commands[i] = operation{cmd, steps}
	}

	{
		fmt.Println("--- Part One ---")

		heading := direction{0, 1}
		position := direction{0, 0}
		for _, c := range commands {
			switch c.op {
			case "N":
				position = position.add(north.mul(c.steps))
			case "E":
				position = position.add(east.mul(c.steps))
			case "S":
				position = position.add(south.mul(c.steps))
			case "W":
				position = position.add(west.mul(c.steps))
			case "F":
				position = position.add(heading.mul(c.steps))
			case "R":
				heading = heading.changeRight(c.steps)
			case "L":
				heading = heading.changeLeft(c.steps)
			}
		}

		fmt.Println(abs(position.y) + abs(position.x))
	}

	{
		fmt.Println("--- Part Two ---")

		waypoint := direction{-1, 10}
		position := direction{0, 0}
		for _, c := range commands {
			switch c.op {
			case "N":
				waypoint = waypoint.add(north.mul(c.steps))
			case "E":
				waypoint = waypoint.add(east.mul(c.steps))
			case "S":
				waypoint = waypoint.add(south.mul(c.steps))
			case "W":
				waypoint = waypoint.add(west.mul(c.steps))
			case "F":
				position = position.add(waypoint.mul(c.steps))
			case "R":
				waypoint = waypoint.changeRight(c.steps)
			case "L":
				waypoint = waypoint.changeLeft(c.steps)
			}
		}

		fmt.Println(abs(position.y) + abs(position.x))
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}
