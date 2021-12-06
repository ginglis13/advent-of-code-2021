// Advent of Code Day 1
// https://adventofcode.com/2021/day/1

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const FLOOR_SIZE = 1000

func countPoints(grid [FLOOR_SIZE][FLOOR_SIZE]int) int {
	count := 0
	for i := 0; i < FLOOR_SIZE; i++ {
		for j := 0; j < FLOOR_SIZE; j++ {
			if grid[j][i] >= 2 {
				count += 1
			}
		}
	}

	return count
}

// consider only horizontal
func part1(ventLines [][]int) {
	var grid [FLOOR_SIZE][FLOOR_SIZE]int
	for _, line := range ventLines {
		x1 := line[0]
		y1 := line[1]
		x2 := line[2]
		y2 := line[3]
		if x1 == x2 {
			// determine which y is greater
			var ylow, yhigh int
			if y1 > y2 {
				yhigh = y1
				ylow = y2
			} else {
				yhigh = y2
				ylow = y1
			}
			for y := ylow; y <= yhigh; y++ {
				grid[y][x1] += 1 // already initialized to 0
			}
		} else if y1 == y2 {
			// determine which x is greater
			var xlow, xhigh int
			if x1 > x2 {
				xhigh = x1
				xlow = x2
			} else {
				xhigh = x2
				xlow = x1
			}
			for x := xlow; x <= xhigh; x++ {
				grid[y1][x] += 1 // already initialized to 0
			}
		}
	}
	fmt.Println(countPoints(grid))
}

// Now consider 45deg diag
func part2(ventLines [][]int) {
	var grid [FLOOR_SIZE][FLOOR_SIZE]int
	for _, line := range ventLines {
		x1 := line[0]
		y1 := line[1]
		x2 := line[2]
		y2 := line[3]
		if x1 == x2 {
			// determine which y is greater
			var ylow, yhigh int
			if y1 > y2 {
				yhigh = y1
				ylow = y2
			} else {
				yhigh = y2
				ylow = y1
			}
			for y := ylow; y <= yhigh; y++ {
				grid[y][x1] += 1 // already initialized to 0
			}
		} else if y1 == y2 {
			// determine which x is greater
			var xlow, xhigh int
			if x1 > x2 {
				xhigh = x1
				xlow = x2
			} else {
				xhigh = x2
				xlow = x1
			}
			for x := xlow; x <= xhigh; x++ {
				grid[y1][x] += 1 // already initialized to 0
			}
		} else if math.Abs(float64(x1)-float64(x2)) == math.Abs(float64(y1)-float64(y2)) { // diag
			var dx, dy int
			if x1 < x2 {
				dx = 1
			} else {
				dx = -1
			}

			if y1 < y2 {
				dy = 1
			} else {
				dy = -1
			}

			x := x1
			y := y1

			for {
				grid[y][x] += 1
				if x == x2 || y == y2 {
					break
				}
				x += dx
				y += dy
			}

		}
	}
	fmt.Println(countPoints(grid))
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	// small brain parsing of this jawn
	var ventLines [][]int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.Replace(line, " -> ", ",", len(line))
		l := strings.Split(line, ",")
		var ventLine []int
		for _, v := range l {
			val, _ := strconv.Atoi(v)
			ventLine = append(ventLine, val)
		}
		ventLines = append(ventLines, ventLine)
	}

	part1(ventLines)
	part2(ventLines)
}
