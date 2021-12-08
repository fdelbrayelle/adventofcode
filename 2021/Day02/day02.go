package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("ioutil.ReadFile Err")
	}
	lines := strings.Split(string(content), "\n")
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
			horizontalPosition = horizontalPosition + value
		case "up":
			depth = depth - value
		case "down":
			depth = depth + value
		}
	}
	fmt.Println("Part 1 result is:", horizontalPosition*depth)
}
