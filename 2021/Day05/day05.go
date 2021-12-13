package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Line struct {
	src, dst Coord
}

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("ioutil.ReadFile Err")
	}
	inputLines := strings.Split(string(content), "\n")
	lines := make([]Line, len(inputLines))
	for i, inputLine := range inputLines {
		coords := strings.Split(string(inputLine), " -> ")
		srcCoords := strings.Split(string(coords[0]), ",")
		dstCoords := strings.Split(string(coords[0]), ",")
		x1, _ := strconv.Atoi(srcCoords[0])
		y1, _ := strconv.Atoi(srcCoords[1])
		x2, _ := strconv.Atoi(dstCoords[0])
		y2, _ := strconv.Atoi(dstCoords[1])
		lines[i] = Line{src: Coord{x: x1, y: y1}, dst: Coord{x: x2, y: y2}}
		diagramMaxSize := findDiagramMaxSize(lines)
		buildDiagram(lines, diagramMaxSize)
	}
}

func findDiagramMaxSize(lines []Line) int {
	highestNumber := 0
	for _, line := range lines {
		if line.src.x > highestNumber {
			highestNumber = line.src.x
		} else if line.src.y > highestNumber {
			highestNumber = line.src.y
		} else if line.dst.x > highestNumber {
			highestNumber = line.dst.x
		} else if line.dst.y > highestNumber {
			highestNumber = line.dst.y
		}
	}
	return highestNumber + 1
}

func buildDiagram(lines []Line, diagramMaxSize int) [][]int {
	diagram := make([][]int, diagramMaxSize)
	for i := 0; i < diagramMaxSize; i++ {
		diagram[i] = make([]int, diagramMaxSize)
		for j := 0; j < diagramMaxSize; j++ {
			diagram[i][j] = 0
		}
	}
	for _, line := range lines {
		if line.src.x == line.dst.x || line.src.y == line.dst.y {
			xDst := line.dst.x - line.src.x
			yDst := line.dst.y - line.src.y
			if xDst == 0 {
				if yDst > 0 { // xDst = 0 ; yDst > 0 ; 1,2 -> 1,3
					for i := line.src.y; i <= line.src.y+yDst; i++ {
						diagram[i][line.src.y]++
					}
				} else if yDst < 0 { // xDst = 0 ; yDst < 0 ; 1,3 -> 1,2
					for i := line.src.y + yDst; i >= line.src.y; i-- {
						diagram[i][line.src.y]++
					}
				} else { // xDst = 0 ; yDst = 0 ; 1,1 -> 1,1
					diagram[line.src.y][line.src.x]++
				}
			} else if xDst > 0 {
				// xDst > 0 ; yDst > 0 ; 1,1 -> 2,2 (not managed for the moment)
				// xDst > 0 ; yDst < 0 ; 1,3 -> 2,2 (not managed for the moment)
				if yDst == 0 { // xDst > 0 ; yDst = 0 ; 1,1 -> 3,1
					for i := line.src.x; i <= line.src.x+xDst; i++ {
						diagram[line.src.y][i]++
					}
				}
			} else if xDst < 0 {
				// xDst < 0 ; yDst > 0 ; 2,1 -> 1,3 (not managed for the moment)
				// xDst < 0 ; yDst < 0 ; 2,4 -> 1,3 (not managed for the moment)
				if yDst == 0 { // xDst < 0 ; yDst = 0 ; 3,2 -> 1,2
					for i := line.src.x + xDst; i >= line.src.x; i-- {
						diagram[line.src.y][i]++
					}
				}
			}
		}
	}
	return diagram
}
