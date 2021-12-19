package main

import (
	"reflect"
	"testing"
)

func TestPassOneDay(t *testing.T) {
	// Given
	fishes := []int{3, 4, 3, 1, 2}
	expectedFishes := []int{2, 3, 2, 0, 1}

	// When
	var actualFishes = passDaysPart1(fishes, 1)

	// Then
	if !reflect.DeepEqual(actualFishes, expectedFishes) {
		t.Errorf("Fishes are different:\nactual  : %v,\nexpected: %v",
			actualFishes, expectedFishes)
	}
}

func TestPass2Days(t *testing.T) {
	// Given
	fishes := []int{3, 4, 3, 1, 2}
	expectedFishes := []int{1, 2, 1, 6, 0, 8}

	// When
	var actualFishes = passDaysPart1(fishes, 2)

	// Then
	if !reflect.DeepEqual(actualFishes, expectedFishes) {
		t.Errorf("Fishes are different:\nactual  : %v,\nexpected: %v",
			actualFishes, expectedFishes)
	}
}

func TestPass18Days(t *testing.T) {
	// Given
	fishes := []int{3, 4, 3, 1, 2}
	expectedFishes := []int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}

	// When
	var actualFishes = passDaysPart1(fishes, 18)

	// Then
	if !reflect.DeepEqual(actualFishes, expectedFishes) {
		t.Errorf("Fishes are different:\nactual  : %v,\nexpected: %v",
			actualFishes, expectedFishes)
	}

	if len(actualFishes) != len(expectedFishes) {
		t.Errorf("Fishes total was incorrect:\nactual  : %v,\nexpected: %v",
			len(actualFishes), len(expectedFishes))
	}
}

func TestPass80Days(t *testing.T) {
	// Given
	fishes := []int{3, 4, 3, 1, 2}
	expectedLenFishes := 5934

	// When
	var actualFishes = passDaysPart1(fishes, 80)

	// Then
	if len(actualFishes) != expectedLenFishes {
		t.Errorf("Fishes total was incorrect:\nactual  : %v,\nexpected: %v",
			len(actualFishes), expectedLenFishes)
	}
}

func TestPass256Days(t *testing.T) {
	// Given
	fishes := []int{3, 4, 3, 1, 2}
	expectedLenFishes := 26984457539

	// When
	var actualLenFishes = passDaysPart2(fishes, 256)

	// Then
	if actualLenFishes != expectedLenFishes {
		t.Errorf("Fishes total was incorrect:\nactual  : %v,\nexpected: %v",
			actualLenFishes, expectedLenFishes)
	}
}
