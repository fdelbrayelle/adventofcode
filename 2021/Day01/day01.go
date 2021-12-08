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
	part1(nums)
	part2(nums)
}

func part1(nums []int) {
	var increasedCount int
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] < nums[i] {
			increasedCount++
		}
	}
	fmt.Println("Part 1 result is:", increasedCount)
}

func part2(nums []int) {
	var increasedCount int
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-3 {
			sumA := nums[i] + nums[i+1] + nums[i+2]
			sumB := nums[i+1] + nums[i+2] + nums[i+3]
			if sumB > sumA {
				increasedCount++
			}
		}
	}
	fmt.Println("Part 2 result is:", increasedCount)
}
