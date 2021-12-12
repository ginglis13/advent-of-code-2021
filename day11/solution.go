package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ginglis13/aocrunner"
)

func boundsCheck(width int, height int, x int, y int) bool {
	if x < width && x >= 0 && y < height && y >= 0 {
		return true
	}

	return false
}

var flashes int

func flash(octopi [][]int, x int, y int) {
	if x < 0 || x >= len(octopi[0]) || y < 0 || y >= len(octopi) {
		return
	}

	if octopi[y][x] > 9 {
		octopi[y][x] = -1_000_000_000
		flashes += 1

		// set all neighbors +1
		width := len(octopi[0])
		height := len(octopi)
		if boundsCheck(width, height, x+1, y) {
			octopi[y][x+1] += 1
		}
		if boundsCheck(width, height, x-1, y) {
			octopi[y][x-1] += 1
		}
		if boundsCheck(width, height, x, y+1) {
			octopi[y+1][x] += 1
		}
		if boundsCheck(width, height, x, y-1) {
			octopi[y-1][x] += 1
		}
		if boundsCheck(width, height, x+1, y+1) {
			octopi[y+1][x+1] += 1
		}
		if boundsCheck(width, height, x+1, y-1) {
			octopi[y-1][x+1] += 1
		}
		if boundsCheck(width, height, x-1, y+1) {
			octopi[y+1][x-1] += 1
		}
		if boundsCheck(width, height, x-1, y-1) {
			octopi[y-1][x-1] += 1
		}

		// call flash on neighbors
		flash(octopi, x+1, y)
		flash(octopi, x-1, y)
		flash(octopi, x, y+1)
		flash(octopi, x, y-1)
		flash(octopi, x+1, y+1)
		flash(octopi, x+1, y-1)
		flash(octopi, x-1, y+1)
		flash(octopi, x-1, y-1)
	}

}

func part1() string {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var octopi [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var o []int
		for _, ch := range line {
			o = append(o, int(ch-'0'))
		}
		octopi = append(octopi, o)
	}

	for sims := 0; sims < 100; sims++ {
		// Increase all vals by 1
		x := len(octopi[0])
		y := len(octopi)

		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				octopi[j][i] += 1
			}
		}
		// Sim - all > 9 flash -> set all adjacent +1
		// Octopus can flash only once per step
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				// set all neighbors, call flash on neighbors!
				flash(octopi, i, j)
			}
		}
		// Set all vals = -1 to 0
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				// set all neighbors, call flash on neighbors!
				if octopi[j][i] < 0 {
					octopi[j][i] = 0
				}
			}
		}

	}

	return fmt.Sprintf("%v", flashes)
}

func part2() string {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var octopi [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var o []int
		for _, ch := range line {
			o = append(o, int(ch-'0'))
		}
		octopi = append(octopi, o)
	}

	for sims := 0; sims < 100_000_000; sims++ {
		// Increase all vals by 1
		x := len(octopi[0])
		y := len(octopi)

		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				octopi[j][i] += 1
			}
		}
		// Sim - all > 9 flash -> set all adjacent +1
		// Octopus can flash only once per step
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				// set all neighbors, call flash on neighbors!
				flash(octopi, i, j)
			}
		}
		// Set all vals < -1 to 0
		numFlashedThisStep := 0
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				// set all neighbors, call flash on neighbors!
				if octopi[j][i] <= 0 {
					octopi[j][i] = 0
					numFlashedThisStep += 1
				}
			}
		}

		// Check if board has flashed simultaneously
		if numFlashedThisStep == 100 {
			fmt.Println(octopi)
			return fmt.Sprintf("%v", sims+1)
		}

	}
	return fmt.Sprintf("%v", "")
}

func main() {
	c := aocrunner.Config{}
	c.Location = "/Users/ginglis/.config/aocrunner-go/config"
	solution := aocrunner.Solution{}

	solution.SetPart1(part1)
	solution.SetPart2(part2)
	solution.SetDay(11)
	solution.SetYear(2021)
	solution.SetCredentials(&c)

	// Submit calls print results
	//fmt.Println(part2())
	//fmt.Println(solution.SubmitPart1())
	fmt.Println(solution.SubmitPart2())
}
