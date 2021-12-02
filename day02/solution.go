// Advent of Code Day 2
// https://adventofcode.com/2021/day/2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("input")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	horizontal := 0
	vertical := 0
	for scanner.Scan() {
		line := scanner.Text()

		command := strings.Fields(line)

		value, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "up":
			vertical -= value
			break
		case "down":
			vertical += value
			break
		case "forward":
			horizontal += value
			break
		default:
			break

		}
	}

	fmt.Println(horizontal * vertical)
}

func part2() {
	file, err := os.Open("input")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	horizontal := 0
	vertical := 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()

		command := strings.Fields(line)

		value, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "up":
			aim -= value
			break
		case "down":
			aim += value
			break
		case "forward":
			horizontal += value
			vertical += aim * value
			break
		default:
			break

		}
	}

	fmt.Println(horizontal * vertical)
}

func main() {
	part1()
	part2()
}
