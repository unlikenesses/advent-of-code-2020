package main

import (
	"fmt"
	"sort"
	"strconv"
	"github.com/unlikenesses/utils"
)

func main() {
	lines := utils.ReadInput()
	invalidNumber := partOne(lines)
	fmt.Println(invalidNumber)
	encryptionWeakness := partTwo(invalidNumber, lines)
	fmt.Println(encryptionWeakness)
}

func partOne(lines []string) int {
	var preamble []int
	var num int
	for i, line := range lines {
		num, _ = strconv.Atoi(line)
		if i < 25 {
			preamble = append(preamble, num)
			continue
		}
		if ! isValid(num, preamble) {
			break
		}
		// remove first element from preamble
		preamble = preamble[1:]
		// add number to end of preamble
		preamble = append(preamble, num)
	}

	return num
}

func isValid(num int, preamble []int) bool {
	for i, first := range preamble {
		for j, second := range preamble {
			if i != j {
				if first + second == num {
					return true
				}
			}
		}
	}

	return false
}

func partTwo(invalidNumber int, lines []string) int {
	numbers := getRange(invalidNumber, lines)
	sort.Ints(numbers)

	return numbers[0] + numbers[len(numbers) - 1]
}

func getRange(invalidNumber int, lines []string) []int {
	var numbers []int
	var sum int
	var num int
	for len(lines) > 0 {
		for _, line := range lines {
			num, _ = strconv.Atoi(line)
			numbers = append(numbers, num)
			sum += num
			if sum == invalidNumber {
				return numbers
			}
			if sum > invalidNumber {
				break
			}
		}
		lines = lines[1:]
		sum = 0
		numbers = nil
	}

	return numbers
}