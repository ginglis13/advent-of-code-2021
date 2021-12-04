package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const BINARY_LENGTH = 12

func binary2decimal(binary [BINARY_LENGTH]int) float64 {
	decimal := 0.0
	for i := 0; i < BINARY_LENGTH; i++ {
		if binary[i] == 1 {
			decimal += math.Pow(2, float64(11-i))
		}
	}
	return decimal
}

// Find most or least common bit at position
func findCommonality(lines []string, position int, most bool) string {
	zeroes := 0
	ones := 0
	for _, line := range lines {
		if line[position] == '0' {
			zeroes += 1
		} else {
			ones += 1
		}
	}

	if most == true {
		if ones >= zeroes {
			return "1"
		} else {
			return "0"
		}
	} else {
		if ones >= zeroes {
			return "0"
		} else {
			return "1"
		}
	}
}

func part1() {
	file, err := os.Open("input")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var gammaLines []string
	var epsilonLines []string
	for scanner.Scan() {
		gammaLines = append(gammaLines, scanner.Text())
		epsilonLines = append(epsilonLines, scanner.Text())
	}

	gamma := [BINARY_LENGTH]int{}
	for i := 0; i < BINARY_LENGTH; i++ {
		gamma[i], _ = strconv.Atoi(findCommonality(gammaLines, i, true))
	}

	epsilon := [BINARY_LENGTH]int{}
	for i := 0; i < BINARY_LENGTH; i++ {
		epsilon[i], _ = strconv.Atoi(findCommonality(epsilonLines, i, false))
	}

	// bin  to dec
	gammaDec := binary2decimal(gamma)
	epsilonDec := binary2decimal(epsilon)
	fmt.Println(gamma, gammaDec, epsilon, epsilonDec)

	fmt.Println(int(gammaDec * epsilonDec))
}

// Helper func for part to modify list such that it contains only the reading
func parseList(lines *[]string, most bool) {
	for i := 0; i < BINARY_LENGTH; i++ {
		keeps := []string{}
		keeper := findCommonality(*lines, i, most)
		for _, line := range *lines {
			if len(*lines) == 1 {
				return
			}
			if string(line[i]) == keeper {
				keeps = append(keeps, line)
			}
		}
		if len(keeps) > 0 {
			*lines = keeps
		}
	}
}

func part2() {
	// lifesupport = oxygen * co2
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var o2lines []string
	var co2lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		o2lines = append(o2lines, scanner.Text())
		co2lines = append(co2lines, scanner.Text())
	}

	parseList(&o2lines, true)
	parseList(&co2lines, false)

	// lines now arrays with 1 element, the resulting  binary digit reading
	// its a string - convert to an array of ints  for binary2decimal conversion
	o2 := [BINARY_LENGTH]int{}
	co2 := [BINARY_LENGTH]int{}
	for i, ch := range o2lines[0] {
		o2[i], _ = strconv.Atoi(string(ch))

	}
	for i, ch := range co2lines[0] {
		co2[i], _ = strconv.Atoi(string(ch))

	}

	fmt.Println(int(binary2decimal(co2) * binary2decimal(o2)))
}

func main() {
	part1()
	part2()
}
