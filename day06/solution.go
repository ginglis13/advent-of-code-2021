// aoc 2021 day 6
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ginglis13/aocrunner"
)

func simulateSpawn(fishies []int, days int) int {

	// Trying to do a more bb math way
	// numSpawned = math.Floor(days - val) / 6 for OG fishies
	// those numSpawned should somehow track what the parent's count was
	// The rest of them... trickier
	// numSpawned = math.Floor(days - val) / 8 for OG fishies
	// but have to account for parents count ..
	// numSpawned = math.Floor((days-parentCount) - val) / 8 for OG fishies
	for i := 0; i < days; i++ {
		fmt.Println("DAY: ", i)
		// Run a day - subtract them jawns
		dayLen := len(fishies)
		for j := 0; j < dayLen; j++ {
			fishies[j] -= 1
			if fishies[j] == -1 { // 0 is valid number -1 indicates switch that jawn
				// that fish's days = 6
				// new fish's day = 8
				fishies[j] = 6
				fishies = append(fishies, 8)
			}
		}

	}

	return len(fishies)
}

func part1() string {
	// each lanternfish creates another every 7 days
	// model each lanternfish as a number representing # days until creates new fish
	// each new fish has 2 more days from its parent fish
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fishies []int
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, ",")
		for _, v := range l {
			val, _ := strconv.Atoi(v)
			fishies = append(fishies, val)
		}
	}

	// now simulate spawning
	return fmt.Sprintf("%v", simulateSpawn(fishies, 80))
}

func part2() string {
	// each lanternfish creates another every 7 days
	// model each lanternfish as a number representing # days until creates new fish
	// each new fish has 2 more days from its parent fish
	file, _ := os.Open("input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fishies []int
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, ",")
		for _, v := range l {
			val, _ := strconv.Atoi(v)
			fishies = append(fishies, val)
		}
	}

	// Create fishmap, count of fish w X days left
	fishMap := make(map[int]int) // daysLeft, count
	for _, daysLeft := range fishies {
		count, keyExists := fishMap[daysLeft]
		if keyExists {
			fishMap[daysLeft] = count + 1
		} else {
			fishMap[daysLeft] = 1
		}
	}

	for i := 0; i < 256; i++ {
		fmt.Println(fishMap)
		copy := make(map[int]int) // daysLeft, count
		for key, val := range fishMap {
			if key == 0 {
				copy[6] += val
				copy[8] += val
			} else {
				copy[key-1] += val
			}

		}
		fishMap = copy
	}

	sum := 0
	for _, val := range fishMap {
		sum += val
	}

	return fmt.Sprintf("%v", sum)
}

func main() {

	c := aocrunner.Config{}
	c.Location = "/Users/ginglis/.config/aocrunner-go/config"
	solution := aocrunner.Solution{}

	solution.SetPart1(part1)
	solution.SetPart2(part2)
	solution.SetDay(6)     // defaults to today's day
	solution.SetYear(2021) // defaults to this year
	solution.SetCredentials(&c)

	// Submit calls print results
	fmt.Println(solution.SubmitPart1())
	fmt.Println(solution.SubmitPart2())
}
