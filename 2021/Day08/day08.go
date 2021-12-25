package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("ioutil.ReadFile Err")
	}
	inputLines := strings.Split(string(content), "\n")
	inputValues, outputValues := getValues(inputLines)
	fmt.Println("Part 1 result is:", getDigitCount(outputValues, 1)+getDigitCount(outputValues, 4)+getDigitCount(outputValues, 7)+getDigitCount(outputValues, 8))
	fmt.Println("Part 2 result is:", decodeLines(inputValues, outputValues))
}

func getValues(inputLines []string) ([]string, []string) {
	inputValues := make([]string, len(inputLines)-1)
	outputValues := make([]string, len(inputLines)-1)
	for i := 0; i < len(inputLines); i++ {
		values := strings.Split(string(inputLines[i]), " | ")
		if len(values) == 2 {
			inputValues[i] = values[0]
			outputValues[i] = values[1]
		}
	}
	return inputValues, outputValues
}

/*
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
*/

func getDigitsLettersPart1() map[int][]string {
	digits := make(map[int][]string, 10)
	digits[0] = []string{"a", "b", "c", "e", "f", "g"}      // 6
	digits[1] = []string{"c", "f"}                          // 2 OK
	digits[2] = []string{"a", "c", "d", "e", "g"}           // 5
	digits[3] = []string{"a", "c", "d", "f", "g"}           // 5
	digits[4] = []string{"b", "c", "d", "f"}                // 4 OK
	digits[5] = []string{"a", "b", "d", "f", "g"}           // 5
	digits[6] = []string{"a", "b", "d", "e", "f", "g"}      // 6
	digits[7] = []string{"a", "c", "f"}                     // 3 OK
	digits[8] = []string{"a", "b", "c", "d", "e", "f", "g"} // 7 OK
	digits[9] = []string{"a", "b", "c", "d", "f", "g"}      // 6
	return digits
}

func getDigitCount(outputLines []string, digit int) int {
	count := 0
	segmentsNumber := len(getDigitsLettersPart1()[digit])
	for _, outputLine := range outputLines {
		numbers := strings.Split(string(outputLine), " ")
		for _, number := range numbers {
			if len(number) == segmentsNumber {
				count++
			}
		}
	}
	return count
}

func findDigitsLetters(inputLine string) map[int][]string {
	digits := make(map[int][]string, 10)
	numbers := strings.Split(string(inputLine), " ")
	sort.Slice(numbers, func(i, j int) bool {
		return len(numbers[i]) < len(numbers[j])
	})
	for _, number := range numbers {
		numberLetters := strings.Split(string(number), "")
		switch len(numberLetters) {
		case 2:
			digits[1] = []string{string(numberLetters[0]), string(numberLetters[1])}
		case 3:
			digits[7] = []string{string(numberLetters[0]), string(numberLetters[1]), string(numberLetters[2])}
		case 4:
			digits[4] = []string{string(numberLetters[0]), string(numberLetters[1]), string(numberLetters[2]), string(numberLetters[3])}
		case 5:
			digitsContent := []string{string(numberLetters[0]), string(numberLetters[1]), string(numberLetters[2]), string(numberLetters[3]), string(numberLetters[4])}
			if containsNTimesDigitLetter(numberLetters, digits[7], 3) {
				digits[3] = digitsContent
			} else if containsNTimesDigitLetter(numberLetters, digits[4], 3) {
				digits[5] = digitsContent
			} else if containsNTimesDigitLetter(numberLetters, digits[7], 2) {
				digits[2] = digitsContent
			}
		case 6:
			digitsContent := []string{string(numberLetters[0]), string(numberLetters[1]), string(numberLetters[2]), string(numberLetters[3]), string(numberLetters[4]), string(numberLetters[5])}
			if containsNTimesDigitLetter(numberLetters, digits[3], 5) {
				digits[9] = digitsContent
			} else if containsNTimesDigitLetter(numberLetters, digits[5], 5) {
				digits[6] = digitsContent
			} else {
				digits[0] = digitsContent
			}
		case 7:
			digits[8] = []string{string(numberLetters[0]), string(numberLetters[1]), string(numberLetters[2]), string(numberLetters[3]), string(numberLetters[4]), string(numberLetters[5]), string(numberLetters[6])}
		}
	}
	for _, digit := range digits {
		sort.Slice(digit, func(i, j int) bool {
			return digit[i] < digit[j]
		})
	}
	return digits
}

func containsNTimesDigitLetter(numberLetters []string, digits []string, expectedCount int) bool {
	count := 0
	for _, digit := range digits {
		for _, numberLetter := range numberLetters {
			if digit == numberLetter {
				count++
			}
		}
	}
	return count == expectedCount
}

func decodeLine(inputLine string, outputLine string) (int, error) {
	result := ""
	numbers := strings.Split(string(outputLine), " ")
	digitsLetters := findDigitsLetters(inputLine)
	for _, number := range numbers {
		for i := 0; i <= 9; i++ {
			digitLetters := digitsLetters[i]
			if len(number) == len(digitLetters) {
				if i == 1 || i == 4 || i == 7 || i == 8 {
					result += strconv.Itoa(i)
				} else {
					numberLetters := strings.Split(string(number), "")
					sort.Slice(numberLetters, func(i, j int) bool {
						return numberLetters[i] < numberLetters[j]
					})
					sort.Slice(digitLetters, func(i, j int) bool {
						return digitLetters[i] < digitLetters[j]
					})
					if reflect.DeepEqual(numberLetters, digitLetters) {
						result += strconv.Itoa(i)
					}
				}
			}
		}
	}
	return strconv.Atoi(result)
}

func decodeLines(inputLines []string, outputLines []string) int {
	sum := 0
	for idx := range inputLines {
		result, _ := decodeLine(inputLines[idx], outputLines[idx])
		sum += result
	}
	return sum
}
