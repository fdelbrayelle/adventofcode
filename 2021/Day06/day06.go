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
	fmt.Println("Part 1 result is:", len(passDaysPart1(fishes, 80)))
	fishes = extractFishes(inputLines)
	fmt.Println("Part 2 result is:", passDaysPart2(fishes, 256))
}

func extractFishes(inputLines []string) []int {
	fishes := make([]int, len(inputLines))
	for i := 0; i < len(inputLines); i++ {
		fishes[i], _ = strconv.Atoi(strings.Trim(inputLines[i], "\n"))
	}
	return fishes
}

func passDaysPart1(fishes []int, days int) []int {
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

func passDaysPart2(fishes []int, days int) int {
	calendar := make([]int, 9)
	for i := 0; i < len(calendar); i++ {
		calendar[i] = 0
	}
	for f := 0; f < len(fishes); f++ {
		calendar[fishes[f]]++
	}

	for d := 0; d < days; d++ {
		add6number, add8number := 0, 0
		for i := 0; i < len(calendar); i++ {
			if calendar[i] > 0 {
				if i == 0 {
					add6number += calendar[i]
					add8number += calendar[i]
				} else if i > 0 {
					calendar[i-1] += calendar[i]
				}
				calendar[i] = 0
			}
		}
		calendar[6] += add6number
		calendar[8] += add8number
	}

	sum := 0
	for i := 0; i < len(calendar); i++ {
		sum += calendar[i]
	}
	return sum
}
