// Advent of Code Day 1
// https://adventofcode.com/2021/day/1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BOARD_SIZE = 5

// Starting position used to find which board is requested in y position
func isWinnerHorizontal(boards [][]string, startingPosition int) bool {
	for i := 0; i < BOARD_SIZE; i++ {
		pos1 := boards[startingPosition+i][0]
		pos2 := boards[startingPosition+i][1]
		pos3 := boards[startingPosition+i][2]
		pos4 := boards[startingPosition+i][3]
		pos5 := boards[startingPosition+i][4]
		if pos1 == "X" && pos2 == "X" && pos3 == "X" && pos4 == "X" && pos5 == "X" {
			return true
		}
	}

	return false
}

// Starting position used to find which board is requested in y position
func isWinnerVertical(boards [][]string, startingPosition int) bool {
	for i := 0; i < BOARD_SIZE; i++ {
		pos1 := boards[startingPosition][i]
		pos2 := boards[startingPosition+1][i]
		pos3 := boards[startingPosition+2][i]
		pos4 := boards[startingPosition+3][i]
		pos5 := boards[startingPosition+4][i]
		if pos1 == "X" && pos2 == "X" && pos3 == "X" && pos4 == "X" && pos5 == "X" {
			return true
		}
	}

	return false
}

func markBoards(number string, boards [][]string) {
	Y := len(boards)

	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < Y; j++ {
			if boards[j][i] == number {
				boards[j][i] = "X"
			}
		}
	}
}

func calculateScore(boards [][]string, startingPosition int, drawing string) int {
	drawingNumber, _ := strconv.Atoi(drawing)
	sum := 0
	for i := 0; i < BOARD_SIZE; i++ {
		for j := startingPosition; j < startingPosition+BOARD_SIZE; j++ {
			if boards[j][i] != "X" {
				val, _ := strconv.Atoi(boards[j][i])
				sum += val
			}
		}
	}

	return sum * drawingNumber
}

func part1(numberDrawings []string, boards [][]string) int {
	lenBoards := len(boards)
	for _, number := range numberDrawings {
		// 1. Mark all the boards
		markBoards(number, boards)

		// 2. For each board in  boards, check if horizontal or vertical winner
		// 3. if winner, calculate score of winning board (sum of all unmarked * winning number)
		for i := 0; i < lenBoards; i += 5 {
			if isWinnerVertical(boards, i) || isWinnerHorizontal(boards, i) {
				return calculateScore(boards, i, number)
			}
		}
	}

	return 0
}

func part2(numberDrawings []string, boards [][]string) int {
	lenBoards := len(boards)
	numBoards := len(boards) / 5
	var prevWinners []int // store position of previous winning boards
	for _, number := range numberDrawings {
		// 1. Mark all the boards
		markBoards(number, boards)

		// 2. For each board in  boards, check if horizontal or vertical winner
		// 3. if winner, calculate score of winning board (sum of all unmarked * winning number)
		for i := 0; i < lenBoards; i += 5 {
			if isWinnerVertical(boards, i) || isWinnerHorizontal(boards, i) {
				inthere := false // see if we've already marked a board as a winner. If not, reduce the number of remaining boards
				for _, winner := range prevWinners {
					if i == winner {
						inthere = true
					}

				}
				if !inthere {
					numBoards -= 1
				}
				if numBoards == 0 { // we got there
					return calculateScore(boards, i, number)
				}
				prevWinners = append(prevWinners, i)
			}
		}
	}

	return 0
}

func main() {
	var numberDrawings []string
	var boards [][]string
	file, err := os.Open("input")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	i := 0
	for scanner.Scan() {
		if i == 0 {
			numberDrawings = strings.Split(scanner.Text(), ",")
			i++
			continue
		}
		if scanner.Text() == "" {
			continue
		}

		// Build Boards
		boards = append(boards, strings.Fields(scanner.Text()))

	}

	fmt.Println(part1(numberDrawings, boards))
	fmt.Println(part2(numberDrawings, boards))
}
