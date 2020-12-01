package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	part1_num1, part1_num2 := partOneNumbers(lines)
	fmt.Println(part1_num1 * part1_num2)

	part2_num1, part2_num2, part2_num3 := partTwoNumbers(lines)
	fmt.Println(part2_num1 * part2_num2 * part2_num3)
}

func partOneNumbers(lines []string) (int, int) {
	for _, line := range lines {
		num1, _ := strconv.Atoi(line)
		for _, line := range lines {
			num2, _ := strconv.Atoi(line)
			if num1 != num2 {
				if num1+num2 == 2020 {
					return num1, num2
				}
			}
		}
	}
	return 0, 0
}

func partTwoNumbers(lines []string) (int, int, int) {
	for _, line := range lines {
		num1, _ := strconv.Atoi(line)
		for _, line := range lines {
			num2, _ := strconv.Atoi(line)
			if num1 != num2 {
				for _, line := range lines {
					num3, _ := strconv.Atoi(line)
					if num3 != num1 && num3 != num2 {
						if num1+num2+num3 == 2020 {
							return num1, num2, num3
						}
					}
				}
			}
		}
	}
	return 0, 0, 0
}
