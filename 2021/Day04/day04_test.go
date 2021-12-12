package main

import (
	"reflect"
	"testing"
)

func TestBuildBoards(t *testing.T) {
	// Given
	/*
		17,11,37,7,89,48,99,28,56,55,57,27,83,59,53,72,6,87,33,82,13,23,35,40,71,47,78,2,39,4,51,1,67,31,79,69,15,73,80,22,92,95,91,43,26,97,36,34,12,96,86,52,66,94,61,76,64,77,85,98,42,68,84,63,60,30,65,19,54,58,24,20,25,75,93,16,18,44,14,88,45,10,9,3,70,74,81,90,46,38,21,49,29,50,0,5,8,32,62,41

		57 80 91 40 12
		62 36 72  0 20
		55 60 25 92 96
		14  2 17 18 86
		1  4 90 66 38

		1 25 81 16 24
		33 40 86 28 96
		4 97 90 32 13
		50 38 35 14 56
		73 42  9 36 67
	*/
	lines := []string{
		"17,11,37,7,89,48,99,28,56,55,57,27,83,59,53,72,6,87,33,82,13,23,35,40,71,47,78,2,39,4,51,1,67,31,79,69,15,73,80,22,92,95,91,43,26,97,36,34,12,96,86,52,66,94,61,76,64,77,85,98,42,68,84,63,60,30,65,19,54,58,24,20,25,75,93,16,18,44,14,88,45,10,9,3,70,74,81,90,46,38,21,49,29,50,0,5,8,32,62,41",
		"57 80 91 40 12\n62 36 72  0 20\n55 60 25 92 96\n14  2 17 18 86\n 1  4 90 66 38",
		" 1 25 81 16 24\n33 40 86 28 96\n 4 97 90 32 13\n50 38 35 14 56\n73 42  9 36 67"}
	board1 := Board{content: make([][]int, 5)}
	board1.content[0] = []int{57, 80, 91, 40, 12}
	board1.content[1] = []int{62, 36, 72, 0, 20}
	board1.content[2] = []int{55, 60, 25, 92, 96}
	board1.content[3] = []int{14, 2, 17, 18, 86}
	board1.content[4] = []int{1, 4, 90, 66, 38}
	board2 := Board{content: make([][]int, 5)}
	board2.content[0] = []int{1, 25, 81, 16, 24}
	board2.content[1] = []int{33, 40, 86, 28, 96}
	board2.content[2] = []int{4, 97, 90, 32, 13}
	board2.content[3] = []int{50, 38, 35, 14, 56}
	board2.content[4] = []int{73, 42, 9, 36, 67}
	expectedBoards := []Board{board1, board2}

	// When
	var actualBoards []Board = buildBoards(lines)

	// Then
	if !reflect.DeepEqual(actualBoards, expectedBoards) {
		t.Errorf("Boards are different:\nactual  : %#v,\nexpected: %#v",
			actualBoards, expectedBoards)
	}
}

func TestMarkBoardsFirstWinner(t *testing.T) {
	// Given
	numbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	initialBoard1 := Board{content: make([][]int, 5)}
	initialBoard1.content[0] = []int{22, 13, 17, 11, 0}
	initialBoard1.content[1] = []int{8, 2, 23, 4, 24}
	initialBoard1.content[2] = []int{21, 9, 14, 16, 7}
	initialBoard1.content[3] = []int{6, 10, 3, 18, 5}
	initialBoard1.content[4] = []int{1, 12, 20, 15, 19}
	initialBoard2 := Board{content: make([][]int, 5)}
	initialBoard2.content[0] = []int{3, 15, 0, 2, 22}
	initialBoard2.content[1] = []int{9, 18, 13, 17, 5}
	initialBoard2.content[2] = []int{19, 8, 7, 25, 23}
	initialBoard2.content[3] = []int{20, 11, 10, 24, 4}
	initialBoard2.content[4] = []int{14, 21, 16, 12, 6}
	initialBoard3 := Board{content: make([][]int, 5)}
	initialBoard3.content[0] = []int{14, 21, 17, 24, 4}
	initialBoard3.content[1] = []int{10, 16, 15, 9, 19}
	initialBoard3.content[2] = []int{18, 8, 23, 26, 20}
	initialBoard3.content[3] = []int{22, 11, 13, 6, 5}
	initialBoard3.content[4] = []int{2, 0, 12, 3, 7}
	initialBoards := []Board{initialBoard1, initialBoard2, initialBoard3}
	board1 := Board{content: make([][]int, 5)}
	board1.content[0] = []int{22, 13, -1, -1, -1}
	board1.content[1] = []int{8, -1, -1, -1, -1}
	board1.content[2] = []int{-1, -1, -1, 16, -1}
	board1.content[3] = []int{6, 10, 3, 18, -1}
	board1.content[4] = []int{1, 12, 20, 15, 19}
	board2 := Board{content: make([][]int, 5)}
	board2.content[0] = []int{3, 15, -1, -1, 22}
	board2.content[1] = []int{-1, 18, 13, -1, -1}
	board2.content[2] = []int{19, 8, -1, 25, -1}
	board2.content[3] = []int{20, -1, 10, -1, -1}
	board2.content[4] = []int{-1, -1, 16, 12, 6}
	board3 := Board{content: make([][]int, 5)}
	board3.content[0] = []int{-1, -1, -1, -1, -1}
	board3.content[1] = []int{10, 16, 15, -1, 19}
	board3.content[2] = []int{18, 8, -1, 26, 20}
	board3.content[3] = []int{22, -1, 13, 6, -1}
	board3.content[4] = []int{-1, -1, 12, 3, -1}
	board3.hasWon = true
	expectedBoards := []Board{board1, board2, board3}
	expectedBoardId := 2
	expectedCurrentNumber := 24

	// When
	actualBoards, actualBoardId, actualCurrentNumber := markBoards(initialBoards, numbers, true)

	// Then
	if !reflect.DeepEqual(actualBoards, expectedBoards) {
		t.Errorf("Boards are different:\nactual  : %#v,\nexpected: %#v",
			actualBoards, expectedBoards)
	}

	if actualBoardId != expectedBoardId {
		t.Errorf("Winner board id is wrong:\nactual  : %d,\nexpected: %d",
			actualBoardId, expectedBoardId)
	}

	if actualCurrentNumber != expectedCurrentNumber {
		t.Errorf("Winning current number is wrong:\nactual  : %d,\nexpected: %d",
			actualCurrentNumber, expectedCurrentNumber)
	}
}

func TestMarkBoardsLastWinner(t *testing.T) {
	// Given
	numbers := []int{0, 1, 2, 9, 10, 11, 18, 19, 20, 21, 22, 23}
	initialBoard1 := Board{content: make([][]int, 3)}
	initialBoard1.content[0] = []int{0, 1, 2}
	initialBoard1.content[1] = []int{3, 4, 5}
	initialBoard1.content[2] = []int{6, 7, 8}
	initialBoard2 := Board{content: make([][]int, 3)}
	initialBoard2.content[0] = []int{9, 10, 11}
	initialBoard2.content[1] = []int{12, 13, 14}
	initialBoard2.content[2] = []int{15, 16, 17}
	initialBoard3 := Board{content: make([][]int, 3)}
	initialBoard3.content[0] = []int{18, 19, 20}
	initialBoard3.content[1] = []int{21, 22, 23}
	initialBoard3.content[2] = []int{24, 25, 26}
	initialBoards := []Board{initialBoard1, initialBoard2, initialBoard3}
	board1 := Board{content: make([][]int, 5)}
	board1.content[0] = []int{-1, -1, -1}
	board1.content[1] = []int{3, 4, 5}
	board1.content[2] = []int{6, 7, 8}
	board1.hasWon = true
	board2 := Board{content: make([][]int, 3)}
	board2.content[0] = []int{-1, -1, -1}
	board2.content[1] = []int{12, 13, 14}
	board2.content[2] = []int{15, 16, 17}
	board2.hasWon = true
	board3 := Board{content: make([][]int, 3)}
	board3.content[0] = []int{-1, -1, -1}
	board3.content[1] = []int{21, 22, 23}
	board3.content[2] = []int{24, 25, 26}
	board3.hasWon = true
	expectedBoards := []Board{board1, board2, board3}
	expectedBoardId := 2

	// When
	actualBoards, actualBoardId, _ := markBoards(initialBoards, numbers, false)

	// Then
	if !reflect.DeepEqual(actualBoards, expectedBoards) {
		t.Errorf("Boards are different:\nactual  : %#v,\nexpected: %#v",
			actualBoards, expectedBoards)
	}

	if actualBoardId != expectedBoardId {
		t.Errorf("Winner board id is wrong:\nactual  : %d,\nexpected: %d",
			actualBoardId, expectedBoardId)
	}
}

func TestCheckBoardsInLine(t *testing.T) {
	// Given
	board1 := Board{content: make([][]int, 5)}
	board1.content[0] = []int{22, 13, -1, -1, -1}
	board1.content[1] = []int{8, 2, -1, -1, -1}
	board1.content[2] = []int{-1, -1, -1, 16, -1}
	board1.content[3] = []int{6, 10, 3, 18, -1}
	board1.content[4] = []int{1, 12, 20, 15, 19}
	board2 := Board{content: make([][]int, 5)}
	board2.content[0] = []int{3, 15, -1, -1, 22}
	board2.content[1] = []int{-1, 18, 13, -1, -1}
	board2.content[2] = []int{19, 8, -1, 25, -1}
	board2.content[3] = []int{20, -1, 10, -1, -1}
	board2.content[4] = []int{14, -1, 16, 12, 6}
	board3 := Board{content: make([][]int, 5)}
	board3.content[0] = []int{-1, -1, -1, -1, -1}
	board3.content[1] = []int{10, 16, 15, -1, 19}
	board3.content[2] = []int{18, 8, -1, 26, 20}
	board3.content[3] = []int{22, -1, 13, 6, -1}
	board3.content[4] = []int{-1, -1, 12, 3, -1}
	markedBoards := []Board{board1, board2, board3}

	// When
	var actualBoards []Board = checkBoards(markedBoards)

	// Then
	if !actualBoards[2].hasWon {
		t.Errorf("There should have a winner!")
	}
}

func TestCheckBoardsInRow(t *testing.T) {
	// Given
	board1 := Board{content: make([][]int, 5)}
	board1.content[0] = []int{22, -1, -1, -1, -1}
	board1.content[1] = []int{8, -1, -1, -1, -1}
	board1.content[2] = []int{-1, -1, -1, -1, -1}
	board1.content[3] = []int{6, -1, 3, 18, -1}
	board1.content[4] = []int{1, 12, 20, 15, 19}
	board2 := Board{content: make([][]int, 5)}
	board2.content[0] = []int{3, 15, -1, -1, 22}
	board2.content[1] = []int{-1, 18, -1, -1, -1}
	board2.content[2] = []int{19, 8, -1, 25, -1}
	board2.content[3] = []int{20, -1, -1, -1, -1}
	board2.content[4] = []int{-1, -1, -1, 12, 6}
	board3 := Board{content: make([][]int, 5)}
	board3.content[0] = []int{-1, -1, -1, -1, -1}
	board3.content[1] = []int{-1, -1, 15, -1, 19}
	board3.content[2] = []int{18, 8, -1, 26, 20}
	board3.content[3] = []int{22, -1, -1, 6, -1}
	board3.content[4] = []int{-1, -1, 12, 3, -1}
	markedBoards := []Board{board1, board2, board3}

	// When
	var actualBoards []Board = checkBoards(markedBoards)

	// Then
	if !actualBoards[1].hasWon {
		t.Errorf("There should have a winner!")
	}
}

func TestGetUnmarkedNumbersSum(t *testing.T) {
	// Given
	board1 := Board{content: make([][]int, 5)}
	board1.content[0] = []int{22, 13, -1, -1, -1}
	board1.content[1] = []int{8, 2, -1, -1, -1}
	board1.content[2] = []int{-1, -1, -1, 16, -1}
	board1.content[3] = []int{6, 10, 3, 18, -1}
	board1.content[4] = []int{1, 12, 20, 15, 19}
	board2 := Board{content: make([][]int, 5)}
	board2.content[0] = []int{3, 15, -1, -1, 22}
	board2.content[1] = []int{-1, 18, 13, -1, -1}
	board2.content[2] = []int{19, 8, -1, 25, -1}
	board2.content[3] = []int{20, -1, 10, -1, -1}
	board2.content[4] = []int{14, -1, 16, 12, 6}
	board3 := Board{content: make([][]int, 5)}
	board3.content[0] = []int{-1, -1, -1, -1, -1}
	board3.content[1] = []int{10, 16, 15, -1, 19}
	board3.content[2] = []int{18, 8, -1, 26, 20}
	board3.content[3] = []int{22, -1, 13, 6, -1}
	board3.content[4] = []int{-1, -1, 12, 3, -1}
	boards := []Board{board1, board2, board3}
	expectedBoardId := 2
	expectedUnmarkedNumbersSum := 188

	// When
	var actualUnmarkedNumbersSum int = getUnmarkedNumbersSum(boards[expectedBoardId])

	// Then
	if actualUnmarkedNumbersSum != expectedUnmarkedNumbersSum {
		t.Errorf("Sum should be %d, not %d!", expectedUnmarkedNumbersSum, actualUnmarkedNumbersSum)
	}
}

func TestAllHasWon(t *testing.T) {
	// Given
	board1 := Board{hasWon: true}
	board2 := Board{hasWon: true}
	board3 := Board{hasWon: true}
	boards := []Board{board1, board2, board3}

	// When
	var actualAllHasWon bool = allHasWon(boards)

	// Then
	if !actualAllHasWon {
		t.Error("All boards should have won!")
	}
}
