package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Board struct {
	content [][]int
	hasWon  bool
}

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("ioutil.ReadFile Err")
	}
	lines := strings.Split(string(content), "\n\n")
	trimNumbers := strings.Split(lines[0], ",")
	numbers := make([]int, len(trimNumbers))
	for i, s := range trimNumbers {
		numbers[i], _ = strconv.Atoi(s)
	}
	part1(lines, numbers)
	part2(lines, numbers)
}

func part1(lines []string, numbers []int) {
	boards := buildBoards(lines)
	_, boardId, firstWinningNumber := markBoards(boards, numbers, true)
	unmarkedNumbersSum := getUnmarkedNumbersSum(boards[boardId])
	fmt.Println("Part 1 result is:", unmarkedNumbersSum*firstWinningNumber)
}

func part2(lines []string, numbers []int) {
	boards := buildBoards(lines)
	_, lastWonBoardId, lastWinningNumber := markBoards(boards, numbers, false)
	unmarkedNumbersSumFromLastBoard := getUnmarkedNumbersSum(boards[lastWonBoardId])
	fmt.Println("Part 2 result is:", unmarkedNumbersSumFromLastBoard*lastWinningNumber)
}

func buildBoards(lines []string) []Board {
	lines = lines[1:]
	var boards []Board
	board := Board{}
	board.content = make([][]int, 5)
	for lineIdx := 0; lineIdx < len(lines); lineIdx++ {
		if len(lines[lineIdx]) == 0 {
			continue
		}
		boardLine := strings.Split(lines[lineIdx], "\n")
		for i := 0; i < len(boardLine); i++ {
			rows := strings.Fields(boardLine[i])
			numRows := len(rows)
			if numRows != 0 {
				board.content[i%numRows] = make([]int, numRows)
				for j := 0; j < numRows; j++ {
					board.content[i%numRows][j], _ = strconv.Atoi(rows[j])
				}
				if i%numRows == numRows-1 {
					boards = append(boards, board)
					board = Board{}
					board.content = make([][]int, numRows)
				}
			}
		}
	}
	return boards
}

func markBoards(boards []Board, numbers []int, firstWinner bool) ([]Board, int, int) {
	boardIdx := 0
	winningNumber := 0
	for n := 0; n < len(numbers); n++ {
		for i := 0; i < len(boards); i++ {
			for lineIdx := 0; lineIdx < len(boards[i].content); lineIdx++ {
				for rowIdx := 0; rowIdx < len(boards[i].content[lineIdx]); rowIdx++ {
					if boards[i].content[lineIdx][rowIdx] == numbers[n] {
						boards[i].content[lineIdx][rowIdx] = -1
						boardIdx = i
						winningNumber = numbers[n]
						if checkBoard(boards[i]) {
							boards[i].hasWon = true
							if firstWinner || allHasWon(boards) {
								return boards, i, numbers[n]
							}
						}
					}
				}
			}
		}
	}
	return boards, boardIdx, winningNumber
}

func allHasWon(boards []Board) bool {
	for _, board := range boards {
		if !board.hasWon {
			return false
		}
	}
	return true
}

func checkBoard(board Board) bool {
	marksInLine := make(map[int]int)
	for _, line := range board.content {
		marksInRow := 0
		for rowIdx, row := range line {
			if row == -1 {
				marksInRow++
				marksInLine[rowIdx]++
			}
		}
		if marksInRow == len(line) {
			return true
		}
	}
	for _, m := range marksInLine {
		if m == len(board.content) {
			return true
		}
	}
	return false
}

func getUnmarkedNumbersSum(board Board) int {
	sum := 0
	for _, line := range board.content {
		for _, row := range line {
			if row != -1 {
				sum += row
			}
		}
	}
	return sum
}
