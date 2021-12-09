package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ginglis13/aocrunner"
)

/*
const zero = "abcdeg"
const one = "ab"
const two = "acdfg"
const three = "abcdf"
const four = "abef"
const five = "bcdef"
const six = "bcdefg"
const seven = "abd"
const eight = "abcdefg"
const nine = "abcdef"
*/

// how we gotta sort a string
// https://golangbyexample.com/sort-string-golang/
func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func check1478(reading string) int {

	outputs := strings.Fields(reading)
	sum := 0

	for _, o := range outputs {
		output := sortString(o)
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
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, "|")

		inputs := strings.Fields(l[0])
		outputs := strings.Fields(l[1])
		zero := ""
		one := ""
		two := ""
		three := ""
		four := ""
		five := ""
		six := ""
		seven := ""
		eight := ""
		nine := ""

		//var mappings = map[string]string{"a": "", "b": "", "c": "", "d": "", "e": "", "f": "", "g": ""}

		for _, i := range inputs {
			input := sortString(i)
			fmt.Println(input)
			if len(input) == 2 {
				fmt.Println(input, " => 1")
				one = input
			} else if len(input) == 4 {
				fmt.Println(input, " => 4")
				four = input
				fmt.Println(four)
			} else if len(input) == 3 {
				fmt.Println(input, " => 7")
				seven = input
			} else if len(input) == 7 {
				fmt.Println(input, " => 8")
				eight = input
				fmt.Println(eight)
			} else if len(input) == 6 { // 0 or 6 or 9
				fmt.Println(input, " => 8")
				fmt.Println(zero, six, nine)
			} else if len(input) == 5 { // 2 or 3 or 5
				fmt.Println(input, " => 8")
				fmt.Println(two, three, five)
			} else {
				continue
			}
		}

		// 1. Find top most segment since we know 1 and 7
		// 2. Label two righthand side segments since we know 1 and 7
		topmost := ""
		righthand := ""
		for _, ch := range seven {
			// check for topmost
			if byte(ch) != one[0] && byte(ch) != one[1] {
				topmost = string(ch)
			} else { // determine righthand
				righthand = fmt.Sprintf("%v%v", righthand, string(ch))

			}
		}
		fmt.Println("TOP", topmost, "RIGHT", righthand)

		// 3. If we know 4, we know 9 (only other besides 8 that shares all segments)
		for _, input := range inputs {
			if len(input) == 6 {
				fmt.Println("len 6: ", input)
				fmt.Println("4: ", four)
				// See if it shares all of 4's chars --> 9
				shares := 0
				for _, ch := range input {
					if byte(ch) == four[0] || byte(ch) == four[1] || byte(ch) == four[2] || byte(ch) == four[3] {
						shares += 1
					}
				}
				if shares == 4 {
					nine = input
				}
			}
		}
		fmt.Println("NINE ==> ", nine)

		// Now we know 9. Can determine between 6 and 0 -> 0 will have both rightmost, 6 will only have 1
		for _, input := range inputs {
			if len(input) == 6 && input != nine {
				shares := 0
				for _, ch := range input {
					if byte(ch) == righthand[0] || byte(ch) == righthand[1] {
						shares += 1
					}
				}
				if shares == 2 {
					zero = input
				} else {
					six = input
				}
			}
		}

		fmt.Println("ZERO ==> ", zero)
		fmt.Println("SIX ==> ", six)

		// Now needa figure out 2 3 5

		// 3 is gonna have both righthand
		for _, input := range inputs {
			if len(input) == 5 {
				shares := 0
				for _, ch := range input {
					if byte(ch) == righthand[0] || byte(ch) == righthand[1] {
						shares += 1
					}
				}
				if shares == 2 {
					three = input
				}
			}
		}

		fmt.Println("THREE ==> ", three)

		// Now 2 5
		// can check w 9 -> 5 shares everything w 9
		for _, input := range inputs {
			if len(input) == 5 && input != three {
				shares := 0
				for _, ch := range input {
					if byte(ch) == nine[0] || byte(ch) == nine[1] || byte(ch) == nine[2] || byte(ch) == nine[3] || byte(ch) == nine[4] || byte(ch) == nine[5] {
						shares += 1
					}
				}
				if shares == 5 {
					five = input
				} else {
					two = input
				}
			}
		}

		fmt.Println("FIVE ==> ", five)
		fmt.Println("TWO ==> ", two)

		digitStr := ""
		for _, o := range outputs {
			output := sortString(o)
			if len(output) == 2 {
				digitStr = fmt.Sprintf("%v%v", digitStr, 1)
			} else if len(output) == 4 {
				digitStr = fmt.Sprintf("%v%v", digitStr, 4)
			} else if len(output) == 3 {
				digitStr = fmt.Sprintf("%v%v", digitStr, 7)
			} else if len(output) == 7 {
				digitStr = fmt.Sprintf("%v%v", digitStr, 8)
			} else if output == sortString(zero) {
				digitStr = fmt.Sprintf("%v%v", digitStr, 0)
			} else if output == sortString(two) {
				digitStr = fmt.Sprintf("%v%v", digitStr, 2)
			} else if output == sortString(three) {
				digitStr = fmt.Sprintf("%v%v", digitStr, 3)
			} else if output == sortString(five) {
				digitStr = fmt.Sprintf("%v%v", digitStr, 5)
			} else if output == sortString(six) {
				digitStr = fmt.Sprintf("%v%v", digitStr, 6)
			} else if output == sortString(nine) {
				digitStr = fmt.Sprintf("%v%v", digitStr, 9)
			}
		}
		fmt.Println("DIGIT STR: ", digitStr)
		val, _ := strconv.Atoi(digitStr)
		sum += val
	}

	fmt.Println("PART2 ", sum)
	return ""
}

func main() {

	part1()
	part2()
	c := aocrunner.Config{}
	c.Location = "/Users/ginglis/.config/aocrunner-go/config"
	solution := aocrunner.Solution{}

	solution.SetPart1(part1)
	solution.SetPart2(part2)
	solution.SetDay(7)          // defaults to today's day
	solution.SetYear(2021)      // defaults to this year
	solution.SetCredentials(&c) // defaults to getting credentials from $HOME/.config/aocrunner-go/config

	// Submit calls print results
	//fmt.Println(solution.SubmitPart2())
}
