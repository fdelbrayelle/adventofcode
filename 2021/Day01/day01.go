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
	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("strconv.Atoi Err")
		}
		nums = append(nums, n)
	}
	var increasedCount int
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] < nums[i] {
			increasedCount++
		}
	}
	fmt.Println("Part 1 result is:", increasedCount)
}

func part2() {
	// TODO
}
