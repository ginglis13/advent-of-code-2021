package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ginglis13/aocrunner"
)

func part1() string {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	sum := 0
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for j, line := range lines {
		for i, v := range line {
			val := int(v - '0')
			up := 0
			down := 0
			left := 0
			right := 0

			if i == 0 {
				right = int(line[i+1] - '0')
				left = 999999
				if j == 0 {
					down = int(lines[j+1][i] - '0')
					up = 999999
				} else if j == len(lines)-1 {
					up = int(lines[j-1][i] - '0')
					down = 999999
				} else {
					down = int(lines[j+1][i] - '0')
					up = int(lines[j-1][i] - '0')
				}
			} else if i == len(line)-1 {
				left = int(line[i-1] - '0')
				right = 999999
				if j == 0 {
					down = int(lines[j+1][i] - '0')
					up = 999999
				} else if j == len(lines)-1 {
					up = int(lines[j-1][i] - '0')
					down = 999999
				} else {
					down = int(lines[j+1][i] - '0')
					up = int(lines[j-1][i] - '0')
				}
			} else if j == 0 {
				down = int(lines[j+1][i] - '0')
				up = 999999
				left = int(line[i-1] - '0')
				right = int(line[i+1] - '0')
			} else if j == len(lines)-1 {
				up = int(lines[j-1][i] - '0')
				down = 999999
				left = int(line[i-1] - '0')
				right = int(line[i+1] - '0')

			} else {
				left = int(line[i-1] - '0')
				right = int(line[i+1] - '0')
				down = int(lines[j+1][i] - '0')
				up = int(lines[j-1][i] - '0')

			}

			if up > val && down > val && left > val && right > val {
				sum += val + 1
			}
		}
	}
	fmt.Println(sum)
	return fmt.Sprintf("%v", sum)
}

func part2() string {
	// Now just sum o :i
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var tubeMap [][]int

	// Create map
	for scanner.Scan() {
		line := scanner.Text()
		var tubeLine []int
		for _, v := range line {
			val := int(v - '0')
			tubeLine = append(tubeLine, val)
		}
		tubeMap = append(tubeMap, tubeLine)
	}

	fmt.Println(tubeMap)

	return fmt.Sprintf("%v", "")
}

func main() {

	c := aocrunner.Config{}
	c.Location = "/Users/ginglis/.config/aocrunner-go/config"
	solution := aocrunner.Solution{}

	solution.SetPart1(part1)
	solution.SetPart2(part2)
	solution.SetDay(9)
	solution.SetYear(2021)
	solution.SetCredentials(&c)

	// Submit calls print results
	part1()
	//part2()
	//fmt.Println(solution.SubmitPart1())
	//fmt.Println(solution.SubmitPart2())
}
