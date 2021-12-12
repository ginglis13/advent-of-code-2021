package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/ginglis13/aocrunner"
)

var brackets = map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
var revBrackets = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}

func part1() string {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scores := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var stack []rune
		for _, ch := range line {
			_, isExists := scores[ch]
			if !isExists {
				stack = append(stack, ch)
			} else { // pop stack and check
				check := stack[len(stack)-1]
				if brackets[ch] != check {
					sum += scores[ch]
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	fmt.Println(sum)

	return fmt.Sprintf("%v", sum)
}

func part2() string {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scores := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

	//var incompleteLines []string
	var allIncompleteScores []int
	for scanner.Scan() {
		line := scanner.Text()
		var stack []rune
		keepLine := true
		for _, ch := range line {
			_, isExists := scores[ch]
			if !isExists {
				stack = append(stack, ch)
			} else { // pop stack and check
				check := stack[len(stack)-1]
				if brackets[ch] != check {
					// discard this line
					keepLine = false
					break
				}
				stack = stack[:len(stack)-1]
			}
		}
		if keepLine {
			score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				score = score * 5
				char := revBrackets[stack[i]]
				score += scores[char]

			}
			allIncompleteScores = append(allIncompleteScores, score)
		}
	}

	// sort allIncompleteScores
	sort.Ints(allIncompleteScores)
	mid := len(allIncompleteScores) / 2

	return fmt.Sprintf("%v", allIncompleteScores[mid])
}

func main() {
	c := aocrunner.Config{}
	c.Location = "/Users/ginglis/.config/aocrunner-go/config"
	solution := aocrunner.Solution{}

	solution.SetPart1(part1)
	solution.SetPart2(part2)
	solution.SetDay(10)
	solution.SetYear(2021)
	solution.SetCredentials(&c)

	// Submit calls print results
	fmt.Println(part2())
	//fmt.Println(solution.SubmitPart1())
	fmt.Println(solution.SubmitPart2())
}
