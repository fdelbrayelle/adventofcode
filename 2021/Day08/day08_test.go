package main

import (
	"testing"
)

func TestGetDigit1Count(t *testing.T) {
	// Given
	outputValues := []string{"fdgacbe cefdb cefbgd gcbe", "fcgedb cgb dgebacf gc", "cg cg fdcagb cbg", "efabcd cedba gadfec cb", "gecf egdcabf bgf bfgea", "gebdcfa ecba ca fadegcb", "cefg dcbef fcge gbcadfe", "ed bcgafe cdgba cbgef", "gbdfcae bgc cg cgb", "fgae cfgab fg bagce"}
	expectedDigitsCount := 26

	// When
	var actualDigitsCount = getDigitCount(outputValues, 1) + getDigitCount(outputValues, 4) + getDigitCount(outputValues, 7) + getDigitCount(outputValues, 8)

	// Then
	if actualDigitsCount != expectedDigitsCount {
		t.Errorf("Digit 1 count was incorrect:\nactual  : %v,\nexpected: %v",
			actualDigitsCount, expectedDigitsCount)
	}
}
