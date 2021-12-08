package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ginglis13/aocrunner"
)

func check1478(reading string) int {
	outputs := strings.Fields(reading)
	sum := 0

	for _, output := range outputs {
		if len(output) == 2 || len(output) == 4 || len(output) == 3 || len(output) == 7 {
			sum += 1
		}
	}

	return sum
}

func part1() string {
	// Now just sum o :i
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, "|")
		sum += check1478(l[1])
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
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, "|")
		fmt.Println(l)
	}

	return ""
}

func main() {

	c := aocrunner.Config{}
	c.Location = "/Users/ginglis/.config/aocrunner-go/config"
	solution := aocrunner.Solution{}

	solution.SetPart1(part1)
	solution.SetPart2(part2)
	solution.SetDay(8)          // defaults to today's day
	solution.SetYear(2021)      // defaults to this year
	solution.SetCredentials(&c) // defaults to getting credentials from $HOME/.config/aocrunner-go/config

	// Submit calls print results
	fmt.Println(solution.SubmitPart2())
}
