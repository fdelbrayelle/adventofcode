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
	loopGame(boards, numbers)
}

func loopGame(boards []Board, numbers []int) bool {
	for i := 0; i < len(numbers); i = i + 5 {
		markedBoards := markBoards(boards, numbers)
	}
	return false
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

func markBoards(boards []Board, numbers []int) []Board {
	for i := 0; i < len(boards); i++ {
		for lineIdx, line := range boards[i].content {
			for rowIdx, row := range line {
				for _, num := range numbers {
					// fmt.Println("row:", row)
					// fmt.Println("num:", num)
					// fmt.Println("boards[i].content[lineIdx][rowIdx]:", boards[i].content[lineIdx][rowIdx])
					if row == num {
						boards[i].content[lineIdx][rowIdx] = -1
					}
					// fmt.Println("boards[i].content[lineIdx][rowIdx]:", boards[i].content[lineIdx][rowIdx])
				}
			}
		}
	}
	return boards
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
