package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ginglis13/aocrunner"
)

type Insertion struct {
	element  string
	position int
}

func part1() string {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	template := ""
	rules := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if len(strings.Fields(line)) == 1 {
			template = line
			continue
		}
		spl := strings.Split(line, "->")
		rules[strings.TrimSpace(spl[0])] = strings.TrimSpace(spl[1])
	}

	fmt.Println(template)
	for times := 0; times < 40; times++ {
		templatePairs := make(map[string][]int)

		fast := 1
		for i := 0; i < len(template) && fast < len(template); i++ {
			templatePairs[template[i:fast+1]] = append(templatePairs[template[i:fast+1]], i)
			fast++
		}

		// Create insertion rules
		allToInsert := []Insertion{}
		for pair, positions := range templatePairs {
			if target, exists := rules[pair]; exists {
				for _, position := range positions {
					toInsert := Insertion{target, position}
					allToInsert = append(allToInsert, toInsert)
				}
			}
		}

		// Execute insertion rules
		newString := ""
		for i, c := range template {
			newString += string(c)
			for _, rule := range allToInsert {
				if rule.position == i { // insert between i and i + 1
					newString += rule.element
				}

			}
		}
		template = newString
	}

	// now get freqs
	freqs := make(map[rune]int)
	for _, r := range template {
		if _, exists := freqs[r]; exists {
			freqs[r] += 1
		} else {
			freqs[r] = 1
		}
	}

	min := 99_999
	max := 0
	for _, val := range freqs {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return fmt.Sprintf("%v", max-min)
}

func part2() string {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	return fmt.Sprintf("%v", "")
}

func main() {
	c := aocrunner.Config{}
	c.Location = "/Users/ginglis/.config/aocrunner-go/config"
	solution := aocrunner.Solution{}

	solution.SetPart1(part1)
	solution.SetPart2(part2)
	solution.SetDay(14)
	solution.SetYear(2021)
	solution.SetCredentials(&c)

	// Submit calls print results
	fmt.Println(part1())
	//fmt.Println(part2())
	//fmt.Println(solution.SubmitPart1())
	//fmt.Println(solution.SubmitPart2())
}
