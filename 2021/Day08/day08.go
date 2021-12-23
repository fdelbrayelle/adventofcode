package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("ioutil.ReadFile Err")
	}
	inputLines := strings.Split(string(content), "\n")
	outputValues := getOutputValues(inputLines)
	fmt.Println("Part 1 result is:", getDigitCount(outputValues, 1)+getDigitCount(outputValues, 4)+getDigitCount(outputValues, 7)+getDigitCount(outputValues, 8))
}

func getOutputValues(inputLines []string) []string {
	outputValues := make([]string, len(inputLines)-1)
	for i := 0; i < len(inputLines); i++ {
		values := strings.Split(string(inputLines[i]), " | ")
		if i != len(inputLines)-1 {
			outputValues[i] = values[1]
		}
	}
	return outputValues
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

func getDigits() map[int][]string {
	digits := make(map[int][]string, 10)
	digits[0] = []string{"a", "b", "c", "e", "f", "g"}      // 6
	digits[1] = []string{"c", "f"}                          // 2 OK
	digits[2] = []string{"a", "c", "d", "e", "g"}           // 5
	digits[3] = []string{"a", "c", "d", "f", "g"}           // 5
	digits[4] = []string{"b", "c", "d", "f"}                // 4 OK
	digits[5] = []string{"a", "b", "d", "f", "g"}           // 5
	digits[6] = []string{"a", "b", "d", "e", "f", "g"}      // 6
	digits[7] = []string{"a", "c", "f"}                     // 3 OK
	digits[8] = []string{"a", "b", "c", "d", "e", "f", "g"} // 8 OK
	digits[9] = []string{"a", "b", "c", "d", "f", "g"}      // 6
	return digits
}

func getDigitCount(outputLines []string, digit int) int {
	count := 0
	segmentsNumber := len(getDigits()[digit])
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
