package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Board struct {
	content [][]int
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
	boards := buildBoards(lines)
	_, boardId, winningCurrentNumber := markBoards(boards, numbers)
	unmarkedNumbersSum := getUnmarkedNumbersSum(boards[boardId])
	fmt.Println("Part 1 result is:", unmarkedNumbersSum*winningCurrentNumber)
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

func markBoards(boards []Board, numbers []int) ([]Board, int, int) {
	for n := 0; n < len(numbers); n++ {
		for boardIdx := 0; boardIdx < len(boards); boardIdx++ {
			for lineIdx := 0; lineIdx < len(boards[boardIdx].content); lineIdx++ {
				for rowIdx := 0; rowIdx < len(boards[boardIdx].content[lineIdx]); rowIdx++ {
					if boards[boardIdx].content[lineIdx][rowIdx] == numbers[n] {
						boards[boardIdx].content[lineIdx][rowIdx] = -1
						if checkBoards(boards) {
							return boards, boardIdx, numbers[n]
						}
					}
				}
			}
		}
	}
	return boards, -1, -1
}

func checkBoards(boards []Board) bool {
	for i := 0; i < len(boards); i++ {
		marksInLine := make(map[int]int)
		for _, line := range boards[i].content {
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
			if m == len(boards[i].content) {
				return true
			}
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