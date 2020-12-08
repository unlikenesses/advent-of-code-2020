package main

import (
	"fmt"
	"sort"
	"github.com/unlikenesses/utils"
)

type byArrayLength [][]string
func (a byArrayLength) Len() int { return len(a)}
func (a byArrayLength) Swap(i, j int) { a[i], a[j] = a[j], a[i]}
func (a byArrayLength) Less(i, j int) bool {
	li, lj := len(a[i]), len(a[j])
	return li < lj
}

func main() {
	lines := utils.ReadInput()
	// counts := getCountsPartOne(lines)
	counts := getCountsPartTwo(lines)
	sum := arraySum(counts)
	fmt.Println(sum)
}

func getCountsPartOne(lines []string) []int {
	var counts []int
	var count int
	var answers []string
	for _, line := range lines {
		if line == "" {
			count = len(answers)
			counts = append(counts, count)
			count = 0
			answers = nil
		}
		for i := 0; i < len(line); i++ {
			if (! inArray(string(line[i]), answers)) {
				answers = append(answers, string(line[i]))
			}
		}
	}
	count = len(answers)
	counts = append(counts, count)

	return counts
}

func getCountsPartTwo(lines []string) []int {
	var counts []int
	var count int
	var groupAnswers [][]string
	var singlePersonAnswers []string
	for _, line := range lines {
		if line == "" {
			intersection := arrayIntersection(groupAnswers)
			count = len(intersection)
			counts = append(counts, count)
			count = 0
			groupAnswers = nil
		}
		for i := 0; i < len(line); i++ {
			if (! inArray(string(line[i]), singlePersonAnswers)) {
				singlePersonAnswers = append(singlePersonAnswers, string(line[i]))
			}
		}
		if (len(singlePersonAnswers) > 0) {
			groupAnswers = append(groupAnswers, singlePersonAnswers)
		}
		singlePersonAnswers = nil
	}
	intersection := arrayIntersection(groupAnswers)
	count = len(intersection)
	counts = append(counts, count)

	return counts
}

func inArray(needle string, haystack []string) bool {
	for _, el := range haystack {
		if (el == needle) {
			return true
		}
	}

	return false
}

func arraySum(input []int) int {
	sum := 0
	for _, number := range input {
		sum += number
	}

	return sum
}

func arrayIntersection(arrays [][]string) []string {
	sort.Sort(byArrayLength(arrays))
	var intersection []string
	if len(arrays) == 1 {
		return arrays[0]
	}
	for _, el := range arrays[0] {
		if isElementInOtherArrays(el, arrays) {
			intersection = append(intersection, el)
		}
	}

	return intersection
}

func isElementInOtherArrays(el string, arrays [][]string) bool {
	var inOthers bool
	for i := 1; i < len(arrays); i++ {
		if inArray(el, arrays[i]) {
			inOthers = true
		} else {
			return false
		}
	}

	return inOthers
}