package main

import (
	"fmt"
	"github.com/unlikenesses/utils"
)

func main() {
	lines := utils.ReadInput()
	seatIds := parseBoardingPasses(lines)
	highest := getHighestId(seatIds)
	fmt.Println(highest)
	mySeat := getMySeat(seatIds)
	fmt.Println(mySeat) 
}

func parseBoardingPasses(lines []string) []int {
	var seatIds []int
	for _, line := range lines {
		id := parseBoardingPass(line);
		seatIds = append(seatIds, id)
	}

	return seatIds
}

func parseBoardingPass(pass string) int {
	min := 0
	max := 127
	// Get row
	for i := 0; i < len(pass) - 3; i++ {
		char := string(pass[i])
		spread := (max - min) + 1
		half := min + (spread / 2)
		if char == "F" {
			max = half - 1
		} else {
			min = half
		}
	}
	row := min
	// Get col
	min = 0
	max = 8
	for i := 7; i < len(pass); i++ {
		char := string(pass[i])
		spread := (max - min) + 1
		half := min + (spread / 2)
		if char == "L" {
			max = half - 1
		} else {
			min = half
		}
	}
	col := min
	// Get id
	id := calculateSeatId(row, col)

	return id
}

func calculateSeatId(row, col int) int {
	return (row * 8) + col
}

func getHighestId(seatIds []int) int {
	highest := 0
	for _, seatId := range seatIds {
		if seatId > highest {
			highest = seatId
		}
	}

	return highest
}

func getMySeat(seatIds []int) int {
	minRow := 0
	maxRow := 127
	minCol := 0
	maxCol := 7
	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			seatId := calculateSeatId(r, c)
			if (! inSeatIds(seatIds, seatId)) && inSeatIds(seatIds, seatId + 1) && inSeatIds(seatIds, seatId - 1) {
				return seatId
			}
		}
	}

	return 0
}

func inSeatIds(ids []int, id int) bool {
	for _, i := range ids {
		if i == id {
			return true
		}
	}
	return false
}