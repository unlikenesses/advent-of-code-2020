package main

import (
	"fmt"
	"sort"
	"strconv"
	"github.com/unlikenesses/utils"
)

func main() {
	lines := utils.ReadInput()
	sorted := sortAdapters(lines)
	partOneValue := partOne(sorted)
	fmt.Println(partOneValue)
	// Part two defeated me... But at least I'm not alone:
	// https://www.reddit.com/r/adventofcode/comments/kalrxl/2020_day_10_a_new_record/
}

func sortAdapters(lines []string) []int {
	var adapters []int
	var num int
	for _, line := range lines {
		num, _ = strconv.Atoi(line)
		adapters = append(adapters, num)
	}
	sort.Ints(adapters)

	return adapters
}

func partOne(adapters []int) int {
	var jolt int
	var oneDiffs int
	var twoDiffs int
	var threeDiffs int
	for _, adapter := range adapters {
		if adapter - 1 == jolt {
			oneDiffs++
			jolt = adapter
		} else if adapter - 2 == jolt {
			twoDiffs++
			jolt = adapter
		} else if adapter - 3 == jolt {
			threeDiffs++
			jolt = adapter
		} 
	}
	threeDiffs += 1

	return oneDiffs * threeDiffs
}