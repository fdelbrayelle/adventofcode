package main

import (
	"reflect"
	"testing"
)

func TestMakeBoard(t *testing.T) {
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
	var actualBoards []Board = makeBoard(lines)

	// Then
	if !reflect.DeepEqual(actualBoards, expectedBoards) {
		t.Errorf("Boards are different:\nactual  : %d,\nexpected: %d",
			actualBoards, expectedBoards)
	}
}
