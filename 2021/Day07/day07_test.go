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
		t.Errorf("Cheapest totalfuel cost was incorrect:\nactual  : %v,\nexpected: %v",
			actualCheapestTotalFuelCost, expectedCheapestTotalFuelCost)
	}
}
