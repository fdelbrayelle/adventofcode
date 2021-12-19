package main

import (
	"testing"
)

func TestGetCheapestTotalFuelCost(t *testing.T) {
	// Given
	horizontalPositions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	expectedCheapestTotalFuelCost := float64(37)

	// When
	var actualCheapestTotalFuelCost = getCheapestTotalFuelCost(horizontalPositions)

	// Then
	if actualCheapestTotalFuelCost != expectedCheapestTotalFuelCost {
		t.Errorf("Cheapest total fuel cost was incorrect:\nactual  : %v,\nexpected: %v",
			actualCheapestTotalFuelCost, expectedCheapestTotalFuelCost)
	}
}

func TestGetCheapestTotalFuelCostPart2(t *testing.T) {
	// Given
	horizontalPositions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	expectedCheapestTotalFuelCost := 168
	expectedBestPosition := 5

	// When
	actualCheapestTotalFuelCost, actualBestPosition := getCheapestTotalFuelCostPart2(horizontalPositions)

	// Then
	if actualCheapestTotalFuelCost != expectedCheapestTotalFuelCost {
		t.Errorf("Cheapest total fuel cost was incorrect:\nactual  : %v,\nexpected: %v",
			actualCheapestTotalFuelCost, expectedCheapestTotalFuelCost)
	}

	if actualBestPosition != expectedBestPosition {
		t.Errorf("Best position was incorrect:\nactual  : %v,\nexpected: %v",
			actualBestPosition, expectedBestPosition)
	}
}

func TestGetTotalFuelCostFromPositionToAnother(t *testing.T) {
	// Given
	src, dst := 5, 16
	expectedFuelCost := 66

	// When
	var actualFuelCost = getTotalFuelCostFromPositionToAnother(src, dst)

	// Then
	if actualFuelCost != expectedFuelCost {
		t.Errorf("Total fuel cost was incorrect:\nactual  : %v,\nexpected: %v",
			actualFuelCost, expectedFuelCost)
	}
}
