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
