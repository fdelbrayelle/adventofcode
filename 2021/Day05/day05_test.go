package main

import (
	"reflect"
	"testing"
)

func TestExtractLines(t *testing.T) {
	// Given
	inputLines := []string{"0,9 -> 5,9", "8,0 -> 0,8", "9,4 -> 3,4", "2,2 -> 2,1", "7,0 -> 7,4", "6,4 -> 2,0", "0,9 -> 2,9", "3,4 -> 1,4", "0,0 -> 8,8", "5,5 -> 8,2"}
	expectedLines := make([]Line, 10)
	expectedLines[0] = Line{src: Coord{0, 9}, dst: Coord{5, 9}}
	expectedLines[1] = Line{src: Coord{8, 0}, dst: Coord{0, 8}}
	expectedLines[2] = Line{src: Coord{9, 4}, dst: Coord{3, 4}}
	expectedLines[3] = Line{src: Coord{2, 2}, dst: Coord{2, 1}}
	expectedLines[4] = Line{src: Coord{7, 0}, dst: Coord{7, 4}}
	expectedLines[5] = Line{src: Coord{6, 4}, dst: Coord{2, 0}}
	expectedLines[6] = Line{src: Coord{0, 9}, dst: Coord{2, 9}}
	expectedLines[7] = Line{src: Coord{3, 4}, dst: Coord{1, 4}}
	expectedLines[8] = Line{src: Coord{0, 0}, dst: Coord{8, 8}}
	expectedLines[9] = Line{src: Coord{5, 5}, dst: Coord{8, 2}}

	// When
	var actualLines = extractLines(inputLines)

	// Then
	if !reflect.DeepEqual(actualLines, expectedLines) {
		t.Errorf("Lines are different:\nactual  : %v,\nexpected: %v",
			actualLines, extractLines)
	}
}

func TestBuildDiagram(t *testing.T) {
	// Given
	/*
		0,9 -> 5,9
		8,0 -> 0,8
		9,4 -> 3,4
		2,2 -> 2,1
		7,0 -> 7,4
		6,4 -> 2,0
		0,9 -> 2,9
		3,4 -> 1,4
		0,0 -> 8,8
		5,5 -> 8,2
	*/
	var lines []Line
	lines = append(lines, Line{src: Coord{0, 9}, dst: Coord{5, 9}})
	lines = append(lines, Line{src: Coord{8, 0}, dst: Coord{0, 8}})
	lines = append(lines, Line{src: Coord{9, 4}, dst: Coord{3, 4}})
	lines = append(lines, Line{src: Coord{2, 2}, dst: Coord{2, 1}})
	lines = append(lines, Line{src: Coord{7, 0}, dst: Coord{7, 4}})
	lines = append(lines, Line{src: Coord{6, 4}, dst: Coord{2, 0}})
	lines = append(lines, Line{src: Coord{0, 9}, dst: Coord{2, 9}})
	lines = append(lines, Line{src: Coord{3, 4}, dst: Coord{1, 4}})
	lines = append(lines, Line{src: Coord{0, 0}, dst: Coord{8, 8}})
	lines = append(lines, Line{src: Coord{5, 5}, dst: Coord{8, 2}})
	expectedDiagram := make([][]int, len(lines))
	expectedDiagram[0] = []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}
	expectedDiagram[1] = []int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}
	expectedDiagram[2] = []int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}
	expectedDiagram[3] = []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}
	expectedDiagram[4] = []int{0, 1, 1, 2, 1, 1, 1, 2, 1, 1}
	expectedDiagram[5] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expectedDiagram[6] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expectedDiagram[7] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expectedDiagram[8] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expectedDiagram[9] = []int{2, 2, 2, 1, 1, 1, 0, 0, 0, 0}
	expectedDiagramMaxSize := 10

	// When
	var actualDiagram = buildDiagram(lines, expectedDiagramMaxSize)

	// Then
	if !reflect.DeepEqual(actualDiagram, expectedDiagram) {
		t.Errorf("Diagrams are different:\nactual  : %v,\nexpected: %v",
			actualDiagram, expectedDiagram)
	}
}

func TestFindDiagramMaxSize(t *testing.T) {
	// Given
	/*
		0,9 -> 5,9
		8,0 -> 0,8
		9,4 -> 3,4
		2,2 -> 2,1
		7,0 -> 7,4
		6,4 -> 2,0
		0,9 -> 2,9
		3,4 -> 1,4
		0,0 -> 8,8
		5,5 -> 8,2
	*/
	var lines []Line
	lines = append(lines, Line{src: Coord{0, 9}, dst: Coord{5, 9}})
	lines = append(lines, Line{src: Coord{8, 0}, dst: Coord{0, 8}})
	lines = append(lines, Line{src: Coord{9, 4}, dst: Coord{3, 4}})
	lines = append(lines, Line{src: Coord{2, 2}, dst: Coord{2, 1}})
	lines = append(lines, Line{src: Coord{7, 0}, dst: Coord{7, 4}})
	lines = append(lines, Line{src: Coord{6, 4}, dst: Coord{2, 0}})
	lines = append(lines, Line{src: Coord{0, 9}, dst: Coord{2, 9}})
	lines = append(lines, Line{src: Coord{3, 4}, dst: Coord{1, 4}})
	lines = append(lines, Line{src: Coord{0, 0}, dst: Coord{8, 8}})
	lines = append(lines, Line{src: Coord{5, 5}, dst: Coord{8, 2}})
	expectedDiagramMaxSize := 10

	// When
	var actualDiagramMaxSize = findDiagramMaxSize(lines)

	// Then
	if actualDiagramMaxSize != expectedDiagramMaxSize {
		t.Errorf("Diagram max size was incorrect:\nactual  : %d,\nexpected: %d",
			actualDiagramMaxSize, expectedDiagramMaxSize)
	}
}
