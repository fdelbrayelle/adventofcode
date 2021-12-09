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
	horizontalPosition := 0
	depth := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		rows := strings.Split(string(line), " ")
		value, err := strconv.Atoi(rows[1])
		if err != nil {
			fmt.Println("strconv.Atoi Err")
		}
		switch rows[0] {
		case "forward":
			horizontalPosition += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}
	fmt.Println("Part 1 result is:", horizontalPosition*depth)
}

func part2(lines []string) {
	horizontalPosition := 0
	depth := 0
	aim := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		rows := strings.Split(string(line), " ")
		value, err := strconv.Atoi(rows[1])
		if err != nil {
			fmt.Println("strconv.Atoi Err")
		}
		switch rows[0] {
		case "forward":
			horizontalPosition += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}
	fmt.Println("Part 2 result is:", horizontalPosition*depth)
}
