package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("ioutil.ReadFile Err")
	}
	lines := strings.Split(string(content), "\n")
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	bit0Counts, bit1Counts := getBitCounts(lines)
	var gammaRate []string
	for i := 0; i < len(bit0Counts); i++ {
		if bit0Counts[i] >= bit1Counts[i] {
			gammaRate = append(gammaRate, "0")
		} else {
			gammaRate = append(gammaRate, "1")
		}
	}
	epsilonRate := revertGammaRate(gammaRate)
	gammaRateInt, _ := strconv.ParseInt(strings.Join(gammaRate, ""), 2, 64)
	epsilonRateInt, _ := strconv.ParseInt(strings.Join(epsilonRate, ""), 2, 64)
	fmt.Println("Part 1 result is:", gammaRateInt*epsilonRateInt)
}

func part2(lines []string) int64 {
	ogr, _ := strconv.ParseInt(getOxygenGeneratorRating(lines, 0), 2, 64)
	csr, _ := strconv.ParseInt(getCO2ScrubberRating(lines, 0), 2, 64)
	result := ogr * csr
	fmt.Println("Part 2 result is:", result)
	return result
}

func getOxygenGeneratorRating(lines []string, rowIdx int) string {
	bit0Counts, bit1Counts := getBitCounts(lines)
	if len(lines) == 1 {
		return lines[0]
	}
	var remainingLines = make([]string, 0)
	mostCommonBitInRow := mostCommonBit(bit0Counts, bit1Counts, rowIdx)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.Split(line, "")[rowIdx] == mostCommonBitInRow {
			remainingLines = append(remainingLines, line)
		}
	}
	return getOxygenGeneratorRating(remainingLines, rowIdx+1)
}

func getCO2ScrubberRating(lines []string, rowIdx int) string {
	bit0Counts, bit1Counts := getBitCounts(lines)
	if len(lines) == 1 {
		return lines[0]
	}
	var remainingLines = make([]string, 0)
	leastCommonBitInRow := leastCommonBit(bit0Counts, bit1Counts, rowIdx)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.Split(line, "")[rowIdx] == leastCommonBitInRow {
			remainingLines = append(remainingLines, line)
		}
	}
	return getCO2ScrubberRating(remainingLines, rowIdx+1)
}

func getBitCounts(lines []string) (map[int]int, map[int]int) {
	var bit0Counts = make(map[int]int)
	var bit1Counts = make(map[int]int)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		for row := 0; row < len(line); row++ {
			switch line[row] {
			case '0':
				bit0Counts[row]++
			case '1':
				bit1Counts[row]++
			}
		}
	}
	return bit0Counts, bit1Counts
}

func mostCommonBit(bit0Counts map[int]int, bit1Counts map[int]int, row int) string {
	if bit0Counts[row] <= bit1Counts[row] {
		return "1"
	}
	return "0"
}

func leastCommonBit(bit0Counts map[int]int, bit1Counts map[int]int, row int) string {
	if bit0Counts[row] <= bit1Counts[row] {
		return "0"
	}
	return "1"
}

func revertGammaRate(gammaRate []string) []string {
	var epsilonRate []string
	for i := 0; i < len(gammaRate); i++ {
		val := "0"
		if gammaRate[i] == "0" {
			val = "1"
		}
		epsilonRate = append(epsilonRate, val)
	}
	return epsilonRate
}
