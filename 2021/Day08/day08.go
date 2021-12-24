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
	_, outputValues := getValues(inputLines)
	fmt.Println("Part 1 result is:", getDigitCount(outputValues, 1)+getDigitCount(outputValues, 4)+getDigitCount(outputValues, 7)+getDigitCount(outputValues, 8))
}

func getValues(inputLines []string) ([]string, []string) {
	inputValues := make([]string, len(inputLines)-1)
	outputValues := make([]string, len(inputLines)-1)
	for i := 0; i < len(inputLines); i++ {
		values := strings.Split(string(inputLines[i]), " | ")
		if i != len(inputLines)-1 {
			inputLines[i] = values[0]
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

/*
 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc
*/

func getDigitsLettersPart2() map[int][]string {
	digits := make(map[int][]string, 10)
	digits[0] = []string{"a", "b", "c", "d", "e", "g"}      // 6
	digits[1] = []string{"a", "b"}                          // 2 OK
	digits[2] = []string{"a", "c", "d", "f", "g"}           // 5
	digits[3] = []string{"a", "b", "c", "d", "f"}           // 5
	digits[4] = []string{"a", "b", "e", "f"}                // 4 OK
	digits[5] = []string{"b", "c", "d", "e", "f"}           // 5
	digits[6] = []string{"b", "c", "d", "e", "f", "g"}      // 6
	digits[7] = []string{"a", "b", "d"}                     // 3 OK
	digits[8] = []string{"a", "b", "c", "d", "e", "f", "g"} // 7 OK
	digits[9] = []string{"a", "b", "c", "d", "e", "f"}      // 6
	return digits
}

func decodeLine(outputLine string) (int, error) {
	result := ""
	numbers := strings.Split(string(outputLine), " ")
	for _, number := range numbers {
		for i := 0; i <= 9; i++ {
			digitLetters := getDigitsLettersPart1()[i]
			if len(number) == len(digitLetters) {
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
	return strconv.Atoi(result)
}

func decodeLines(outputLines []string) int {
	sum := 0
	for _, outputLine := range outputLines {
		result, _ := decodeLine(outputLine)
		sum += result
	}
	return sum
}
