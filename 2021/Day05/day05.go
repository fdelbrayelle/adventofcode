package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	lines := extractLines(inputLines)
	diagramMaxSize := findDiagramMaxSize(lines)
	diagram := buildDiagram(lines, diagramMaxSize)
	numberMostDangerousZones := findNumberMostDangerousZones(diagram)
	fmt.Println("Part 1 result is:", numberMostDangerousZones)
	diagramWithDiagonalLines := buildDiagramWithDiagonalLines(lines, diagramMaxSize)
	numberMostDangerousZonesWithDiagonalLines := findNumberMostDangerousZones(diagramWithDiagonalLines)
	fmt.Println("Part 2 result is:", numberMostDangerousZonesWithDiagonalLines)
}

func extractLines(inputLines []string) []Line {
	lines := make([]Line, len(inputLines))
	for i, inputLine := range inputLines {
		coords := strings.Split(string(inputLine), " -> ")
		if len(coords) == 2 {
			srcCoords := strings.Split(string(coords[0]), ",")
			dstCoords := strings.Split(string(coords[1]), ",")
			x1, _ := strconv.Atoi(srcCoords[0])
			y1, _ := strconv.Atoi(srcCoords[1])
			x2, _ := strconv.Atoi(dstCoords[0])
			y2, _ := strconv.Atoi(dstCoords[1])
			lines[i] = Line{src: Coord{x: x1, y: y1}, dst: Coord{x: x2, y: y2}}
		}
	}
	return lines
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
					for i := line.src.y; i <= line.dst.y; i++ {
						diagram[i][line.src.x]++
					}
				} else if yDst < 0 { // xDst = 0 ; yDst < 0 ; 1,3 -> 1,2
					for i := line.src.y; i >= line.dst.y; i-- {
						diagram[i][line.src.x]++
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
					for i := line.src.x; i >= line.dst.x; i-- {
						diagram[line.src.y][i]++
					}
				}
			}
		}
	}
	return diagram
}

func buildDiagramWithDiagonalLines(lines []Line, diagramMaxSize int) [][]int {
	diagram := make([][]int, diagramMaxSize)
	for i := 0; i < diagramMaxSize; i++ {
		diagram[i] = make([]int, diagramMaxSize)
		for j := 0; j < diagramMaxSize; j++ {
			diagram[i][j] = 0
		}
	}
	for _, line := range lines {
		xDst := line.dst.x - line.src.x
		yDst := line.dst.y - line.src.y
		if line.src.x == line.dst.x || line.src.y == line.dst.y {
			if xDst == 0 {
				if yDst > 0 { // xDst = 0 ; yDst > 0 ; 1,2 -> 1,3
					for i := line.src.y; i <= line.dst.y; i++ {
						diagram[i][line.src.x]++
					}
				} else if yDst < 0 { // xDst = 0 ; yDst < 0 ; 1,3 -> 1,2
					for i := line.src.y; i >= line.dst.y; i-- {
						diagram[i][line.src.x]++
					}
				} else { // xDst = 0 ; yDst = 0 ; 1,1 -> 1,1
					diagram[line.src.y][line.src.x]++
				}
			} else if xDst > 0 {
				if yDst == 0 { // xDst > 0 ; yDst = 0 ; 1,1 -> 3,1
					for i := line.src.x; i <= line.src.x+xDst; i++ {
						diagram[line.src.y][i]++
					}
				}
			} else if xDst < 0 {
				if yDst == 0 { // xDst < 0 ; yDst = 0 ; 3,2 -> 1,2
					for i := line.src.x; i >= line.dst.x; i-- {
						diagram[line.src.y][i]++
					}
				}
			}
		} else if math.Abs(float64(xDst)) == math.Abs(float64(yDst)) {
			if xDst < 0 {
				if yDst < 0 { // 6,4 -> 2,0 ; yDst = -4 ; xDst = -4 ; xDst < 0 ; yDst < 0
					x := line.src.x
					for y := line.src.y; y >= line.dst.y; y-- {
						if x != line.dst.x-1 {
							diagram[y][x]++
							x--
						}
					}
				} else { // 8,0 -> 0,8 ; yDst = 8 ; xDst = -8 ; xDst < 0 ; yDst > 0
					y := line.src.y
					for x := line.src.x; x >= line.dst.x; x-- {
						if y != line.dst.y+1 {
							diagram[y][x]++
							y++
						}
					}
				}
			} else { // 5,5 -> 8,2 ; yDst = -3 ; xDst = 3 ; xDst > 0 ; yDst < 0
				if yDst < 0 {
					y := line.src.y
					for x := line.src.x; x <= line.dst.x; x++ {
						if y != line.dst.y-1 {
							diagram[y][x]++
							y--
						}
					}
				} else if yDst > 0 { // 0,0 -> 8,8 ; yDst = 8 ; xDst = 8 - xDst > 0 ; yDst < 0
					y := line.src.y
					for x := line.src.x; x <= line.dst.x; x++ {
						if y != line.dst.y+1 {
							diagram[y][x]++
							y++
						}
					}
				}
			}
		}
	}
	return diagram
}

func findNumberMostDangerousZones(diagram [][]int) int {
	numberMostDangerousZones := 0
	for lineIdx := range diagram {
		for rowIdx := range diagram[lineIdx] {
			if diagram[lineIdx][rowIdx] > 1 {
				numberMostDangerousZones++
			}
		}
	}
	return numberMostDangerousZones
}
