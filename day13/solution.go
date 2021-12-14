package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

// For updating dimensions after a fold
func getDimensions(coords map[Point]bool) Point {
	xMax := 0
	yMax := 0
	for pt := range coords {
		if pt.x > xMax {
			xMax = pt.x
		}
		if pt.y > yMax {
			yMax = pt.y
		}

	}
	return Point{xMax, yMax}
}

func followInstruction(coords map[Point]bool, instruction string, dimensions Point) (int, Point) {
	// get x vs y
	spl := strings.Split(instruction, "=")
	foldPoint, _ := strconv.Atoi(spl[1])
	if spl[0] == "x" {
		for pt := range coords {
			// fold along a vertical line
			// all points right of line now have inverse coords i.e. the endpt now pt 0
			// new X will be dimensions - pt.x
			if pt.x <= foldPoint {
				continue
			}

			newX := dimensions.x - pt.x
			newPt := Point{newX, pt.y}
			coords[newPt] = true

			// Complete fold by removing the OG point from the coords set
			_, ok := coords[pt]
			if ok {
				delete(coords, pt)

			}

		}

	} else if spl[0] == "y" {
		for pt := range coords {
			// fold along a horizontal line
			// all points below line now have inverse coords i.e. the endpt now pt 0
			// new Y will be dimensions - pt.y
			if pt.y <= foldPoint {
				continue
			}

			newY := dimensions.y - pt.y
			newPt := Point{pt.x, newY}
			coords[newPt] = true

			// Complete fold by removing the OG point from the coords set
			_, ok := coords[pt]
			if ok {
				delete(coords, pt)

			}
		}
	}

	// update dimensions - needed for multiple steps
	dimensions = getDimensions(coords)

	return len(coords), dimensions
}

// Part 2 helper
func followInstructions(coords map[Point]bool, instructions []string, dimensions Point) Point {
	for _, instruction := range instructions {
		_, dimension := followInstruction(coords, instruction, dimensions)
		dimensions = dimension
	}

	return dimensions
}

func part1() string {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	coords := make(map[Point]bool)
	var instructions []string
	xMax := 0
	yMax := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "fold") {
			// get fields
			fields := strings.Fields(line)
			instructions = append(instructions, fields[2])
			continue
		}
		// create Point from coords
		xy := strings.Split(line, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
		coords[Point{x, y}] = true
	}

	res, _ := followInstruction(coords, instructions[0], Point{xMax, yMax})
	return fmt.Sprintf("%v", res)
}

func part2() string {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	coords := make(map[Point]bool)
	var instructions []string
	xMax := 0
	yMax := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "fold") {
			// get fields
			fields := strings.Fields(line)
			instructions = append(instructions, fields[2])
			continue
		}
		// create Point from coords
		xy := strings.Split(line, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
		coords[Point{x, y}] = true
	}

	finalDim := followInstructions(coords, instructions, Point{xMax, yMax})
	// now print the letter grids
	for j := 0; j <= finalDim.y; j++ {
		for i := 0; i <= finalDim.x; i++ {
			if _, exists := coords[Point{i, j}]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return fmt.Sprintf("%v", "")
}

func main() {
	fmt.Println("P1: ", part1())
	fmt.Println("P2: ")
	fmt.Println(part2())
}
