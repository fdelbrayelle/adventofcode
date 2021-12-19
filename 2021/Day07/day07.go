package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("ioutil.ReadFile Err")
	}
	inputLines := strings.Split(string(content), ",")
	horizontalPositions := extractHorizontalPositions(inputLines)
	fmt.Println("Part 1 result is:", getCheapestTotalFuelCost(horizontalPositions))
	horizontalPositions = extractHorizontalPositions(inputLines)
	cheapestTotalFuelCostPart2, _ := getCheapestTotalFuelCostPart2(horizontalPositions)
	fmt.Println("Part 2 result is:", cheapestTotalFuelCostPart2)
}

func extractHorizontalPositions(inputLines []string) []int {
	horizontalPositions := make([]int, len(inputLines))
	for i := 0; i < len(inputLines); i++ {
		horizontalPositions[i], _ = strconv.Atoi(strings.Trim(inputLines[i], "\n"))
	}
	return horizontalPositions
}

func getCheapestTotalFuelCost(horizontalPositions []int) float64 {
	cheapestTotalFuelCost := float64(0)
	for i := 0; i < len(horizontalPositions); i++ {
		totalFuelCost := float64(0)
		for j := 0; j < len(horizontalPositions); j++ {
			totalFuelCost += math.Abs(float64(horizontalPositions[j] - horizontalPositions[i]))
		}
		if i == 0 || totalFuelCost < cheapestTotalFuelCost {
			cheapestTotalFuelCost = totalFuelCost
		}
	}
	return cheapestTotalFuelCost
}

func getCheapestTotalFuelCostPart2(horizontalPositions []int) (int, int) {
	cheapestTotalFuelCost := 0
	bestPosition := 0
	for i := 0; i < len(horizontalPositions); i++ {
		totalFuelCost := 0
		for j := 0; j < len(horizontalPositions); j++ {
			totalFuelCost += getTotalFuelCostFromPositionToAnother(i, horizontalPositions[j])
		}
		if i == 0 || totalFuelCost < cheapestTotalFuelCost {
			cheapestTotalFuelCost = totalFuelCost
			bestPosition = i
		}
	}
	return cheapestTotalFuelCost, bestPosition
}

func getTotalFuelCostFromPositionToAnother(src int, dst int) int {
	totalFuelCost := 0
	if src < dst {
		moveCost := 1
		for i := src; i < dst; i++ {
			totalFuelCost += moveCost
			moveCost++
		}
	} else if src > dst {
		moveCost := 1
		for i := src; i > dst; i-- {
			totalFuelCost += moveCost
			moveCost++
		}
	}
	return totalFuelCost
}
