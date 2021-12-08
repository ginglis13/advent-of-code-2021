package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ginglis13/aocrunner"
)

func median(crabs []int) int {
	sort.Ints(crabs)
	mid := len(crabs) / 2

	return crabs[mid]
}

func mean(crabs []int) int {
	sum := 0
	for _, val := range crabs {
		sum += val
	}

	return sum / len(crabs)
}

func makeFreqMap(crabs []int) map[int]int {
	freqs := make(map[int]int)
	for _, pos := range crabs {
		count, keyExists := freqs[pos]
		if keyExists {
			freqs[pos] = count + 1
		} else {
			freqs[pos] = 1
		}
	}

	return freqs
}

func part1() string {
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var crabs []int
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, ",")
		for _, v := range l {
			val, _ := strconv.Atoi(v)
			crabs = append(crabs, val)
		}
	}

	freqs := makeFreqMap(crabs)
	median := median(crabs)
	fuelUse := 0.0
	for key, occurrences := range freqs {
		for i := 0; i < occurrences; i++ {
			fuelUse += math.Abs(float64(median) - float64(key))
		}
	}

	return fmt.Sprintf("%v", fuelUse)
}

func calcFuelUsage(diff int) float64 {
	sum := 0.0
	for i := 1; i <= diff; i++ {
		sum += float64(i)
	}

	return sum
}

func part2() string {
	// Now just sum o :i
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var crabs []int
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, ",")
		for _, v := range l {
			val, _ := strconv.Atoi(v)
			crabs = append(crabs, val)
		}
	}

	//  have to find position based on exponential increase of fuel use
	// -> use mean
	freqs := makeFreqMap(crabs)
	mean := mean(crabs)
	fuelUse := 0.0
	for key, occurrences := range freqs {
		for i := 0; i < occurrences; i++ {
			diff := math.Abs(float64(mean) - float64(key))
			fuelUse += calcFuelUsage(int(diff))
		}
	}

	return fmt.Sprintf("%v", int(fuelUse))
}

func main() {

	c := aocrunner.Config{}
	c.Location = "/Users/ginglis/.config/aocrunner-go/config"
	solution := aocrunner.Solution{}

	solution.SetPart1(part1)
	solution.SetPart2(part2)
	solution.SetDay(7)
	solution.SetYear(2021)
	solution.SetCredentials(&c)

	// Submit calls print results
	fmt.Println(solution.SubmitPart1())
	fmt.Println(solution.SubmitPart2())
}
