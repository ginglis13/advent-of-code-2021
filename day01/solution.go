// Advent of Code Day 1
// https://adventofcode.com/2021/day/1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func collectReadings() []int {
	file, err := os.Open("input")
	defer file.Close()

	if err != nil {
		panic(err) // oh boi
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var readings []int

	for scanner.Scan() {
		reading, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		readings = append(readings, reading)
	}

	return readings
}

func sum(a []int) int {
	res := 0
	for _, val := range a {
		res += val
	}

	return res
}

func part1() {
	readings := collectReadings()

	count := 0
	for i, ln := range readings {
		if i+1 >= len(readings) {
			break
		}
		if ln < readings[i+1] {
			count += 1
		}
	}

	fmt.Println(count)
}

func part2() {
	readings := collectReadings()
	count := 0

	// Compare slices of 3
	for i, _ := range readings {
		if i+3 >= len(readings) {
			break
		}
		if sum(readings[i:(i+3)]) < sum(readings[(i+1):(i+4)]) {
			count += 1
		}
	}

	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
