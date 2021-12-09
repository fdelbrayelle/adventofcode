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
