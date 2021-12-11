package main

import "testing"

func TestPart2(t *testing.T) {
	// Given
	lines := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010"}
	var expectedLifeSupportRating int64 = 230

	// When
	var actualLifeSupportRating int64 = part2(lines)

	// Then
	if actualLifeSupportRating != expectedLifeSupportRating {
		t.Errorf("Life support rating was incorrect, actual: %d, expected: %d",
			actualLifeSupportRating, expectedLifeSupportRating)
	}
}

func TestGetOxygenGeneratorRating(t *testing.T) {
	// Given
	lines := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010"}
	expectedOxygenGeneratorRating := "10111"

	// When
	actualOxygenGeneratorRating := getOxygenGeneratorRating(lines, 0)

	// Then
	if actualOxygenGeneratorRating != expectedOxygenGeneratorRating {
		t.Errorf("Oxygen generator rating was incorrect, actual: %s, expected: %s",
			actualOxygenGeneratorRating, expectedOxygenGeneratorRating)
	}
}

func TestGetCO2ScrubberRating(t *testing.T) {
	// Given
	lines := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010"}
	expectedCO2ScrubberRating := "01010"

	// When
	actualCO2ScrubberRating := getCO2ScrubberRating(lines, 0)

	// Then
	if actualCO2ScrubberRating != expectedCO2ScrubberRating {
		t.Errorf("CO2 scrubber rating was incorrect, actual: %s, expected: %s",
			actualCO2ScrubberRating, expectedCO2ScrubberRating)
	}
}
