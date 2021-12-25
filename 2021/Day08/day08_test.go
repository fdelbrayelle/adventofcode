package main

import (
	"reflect"
	"testing"
)

func TestGetDigitCounts(t *testing.T) {
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

func TestDecodeLine(t *testing.T) {
	// Given
	inputLine := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"
	outputLine := "cdfeb fcadb cdfeb cdbaf"
	expectedResult := 5353

	// When
	var actualResult, _ = decodeLine(inputLine, outputLine)

	// Then
	if actualResult != expectedResult {
		t.Errorf("Result was incorrect:\nactual  : %v,\nexpected: %v",
			actualResult, expectedResult)
	}
}

func TestDecodeLinePart2(t *testing.T) {
	// Given
	inputLine := "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb"
	outputLine := "fdgacbe cefdb cefbgd gcbe"
	expectedResult := 8394

	// When
	var actualResult, _ = decodeLine(inputLine, outputLine)

	// Then
	if actualResult != expectedResult {
		t.Errorf("Result was incorrect:\nactual  : %v,\nexpected: %v",
			actualResult, expectedResult)
	}
}

func TestDecodeLines(t *testing.T) {
	// Given
	inputLines := []string{"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb", "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec", "fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef", "fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega", "aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga", "fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf", "dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf", "bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd", "egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg", "gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc"}
	outputLines := []string{"fdgacbe cefdb cefbgd gcbe", "fcgedb cgb dgebacf gc", "cg cg fdcagb cbg", "efabcd cedba gadfec cb", "gecf egdcabf bgf bfgea", "gebdcfa ecba ca fadegcb", "cefg dcbef fcge gbcadfe", "ed bcgafe cdgba cbgef", "gbdfcae bgc cg cgb", "fgae cfgab fg bagce"}
	expectedResult := 61229

	// When
	var actualResult = decodeLines(inputLines, outputLines)

	// Then
	if actualResult != expectedResult {
		t.Errorf("Result was incorrect:\nactual  : %v,\nexpected: %v",
			actualResult, expectedResult)
	}
}

func TestFindDigitsLetters(t *testing.T) {
	// Given
	inputLine := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"
	expectedDigitsLetters := make(map[int][]string, 10)
	expectedDigitsLetters[0] = []string{"a", "b", "c", "d", "e", "g"}      // 6
	expectedDigitsLetters[1] = []string{"a", "b"}                          // 2 OK
	expectedDigitsLetters[2] = []string{"a", "c", "d", "f", "g"}           // 5
	expectedDigitsLetters[3] = []string{"a", "b", "c", "d", "f"}           // 5
	expectedDigitsLetters[4] = []string{"a", "b", "e", "f"}                // 4 OK
	expectedDigitsLetters[5] = []string{"b", "c", "d", "e", "f"}           // 5
	expectedDigitsLetters[6] = []string{"b", "c", "d", "e", "f", "g"}      // 6
	expectedDigitsLetters[7] = []string{"a", "b", "d"}                     // 3 OK
	expectedDigitsLetters[8] = []string{"a", "b", "c", "d", "e", "f", "g"} // 7 OK
	expectedDigitsLetters[9] = []string{"a", "b", "c", "d", "e", "f"}      // 6

	// When
	var actualDigitsLetters = findDigitsLetters(inputLine)

	// Then
	if !reflect.DeepEqual(actualDigitsLetters, expectedDigitsLetters) {
		t.Errorf("Digits letters were incorrect:\nactual  : %v,\nexpected: %v",
			actualDigitsLetters, expectedDigitsLetters)
	}
}

func TestContainsNTimesDigitLetter(t *testing.T) {
	// Given
	numberLetters := []string{"f", "b", "c", "a", "d"}
	digits := []string{"d", "a", "b"}
	expectedCount := 3

	// When
	var actualResult = containsNTimesDigitLetter(numberLetters, digits, expectedCount)

	// Then
	if !actualResult {
		t.Errorf("Number letters did not contain expected count from digits")
	}
}
