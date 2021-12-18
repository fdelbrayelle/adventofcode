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
	inputLines := strings.Split(string(content), ",")
	fishes := extractFishes(inputLines)
	fmt.Println("Part 1 result is:", len(passDays(fishes, 80)))
}

func extractFishes(inputLines []string) []int {
	fishes := make([]int, len(inputLines))
	for i := 0; i < len(inputLines); i++ {
		fishes[i], _ = strconv.Atoi(strings.Trim(inputLines[i], "\n"))
	}
	return fishes
}

func passDays(fishes []int, days int) []int {
	for d := 0; d < days; d++ {
		babies := 0
		for f := 0; f < len(fishes); f++ {
			if fishes[f] > 0 {
				fishes[f]--
			} else {
				fishes[f] = 6
				babies++
			}
		}
		for i := 0; i < babies; i++ {
			fishes = append(fishes, 8)
		}
	}
	return fishes
}
