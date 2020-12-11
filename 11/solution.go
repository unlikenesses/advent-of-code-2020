package main

import (
	"fmt"
	"github.com/unlikenesses/utils"
)

type Vector struct {
	x int
	y int
}

var seats [][]string

func main() {
	lines := utils.ReadInput()
	seats = getSeats(lines)
	partOneValue := partOne()
	fmt.Println(partOneValue)
	
	seats = getSeats(lines)
	partTwoValue := partTwo()
	fmt.Println(partTwoValue)
}

func getSeats(lines []string) [][]string {
	var seats [][]string
	var row []string
	for _, line := range lines {
		row = nil
		for i := 0; i < len(line); i++ {
			row = append(row, string(line[i]))
		}
		seats = append(seats, row)
	}

	return seats
}

func partOne() int {
	hasChanged := true
	rounds := 0
	var newSeats [][]string
	var newRow []string
	var newSeat string
	for hasChanged {
		rounds++
		hasChanged = false
		for y, row := range seats {
			for x, col := range row {
				newSeat = col
				changedValue := getSeatNewValue(x, y, col)
				if changedValue != col {
					hasChanged = true
					newSeat = changedValue
				}
				newRow = append(newRow, newSeat)
			}
			newSeats = append(newSeats, newRow)
			newRow = nil
		}
		seats = newSeats
		newSeats = nil
	}
	return getNumOccupied()
}

func getSeatNewValue(x int, y int, currentValue string) string {
	numAdjacentOccupied := getNumAdjacentOccupied(x, y)
	if currentValue == "L" && numAdjacentOccupied == 0 {
		return "#"
	} else if currentValue == "#" && numAdjacentOccupied >= 4 {
		return "L"
	}
	return currentValue
}

func getNumAdjacentOccupied(x int, y int) int {
	var num int
	if isOccupied(x - 1, y - 1) {
		num++
	}
	if isOccupied(x - 1, y) {
		num++
	}
	if isOccupied(x - 1, y + 1) {
		num++
	}
	if isOccupied(x, y - 1) {
		num++
	}
	if isOccupied(x, y + 1) {
		num++
	}
	if isOccupied(x + 1, y - 1) {
		num++
	}
	if isOccupied(x + 1, y) {
		num++
	}
	if isOccupied(x + 1, y + 1) {
		num++
	}
	
	return num
}

func partTwo() int {
	hasChanged := true
	rounds := 0
	var newSeats [][]string
	var newRow []string
	var newSeat string
	for hasChanged {
		rounds++
		hasChanged = false
		for y, row := range seats {
			for x, col := range row {
				newSeat = col
				changedValue := getSeatNewValuePartTwo(x, y, col)
				if changedValue != col {
					hasChanged = true
					newSeat = changedValue
				}
				newRow = append(newRow, newSeat)
			}
			newSeats = append(newSeats, newRow)
			newRow = nil
		}
		seats = newSeats
		newSeats = nil
	}

	return getNumOccupied()
}

func getSeatNewValuePartTwo(x int, y int, currentValue string) string {
	numAdjacentOccupied := getNumAdjacentOccupiedPartTwo(x, y)
	if currentValue == "L" && numAdjacentOccupied == 0 {
		return "#"
	} else if currentValue == "#" && numAdjacentOccupied >= 5 {
		return "L"
	}
	return currentValue
}

func getNumAdjacentOccupiedPartTwo(x int, y int) int {
	var num int
	vectors := []Vector{
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, vector := range vectors {
		if nearestSeatInDirection(x, y, vector) == "#" {
			num++
		}
	} 

	return num
}

func nearestSeatInDirection(x, y int, vector Vector) string {
	foundSeat := false
	var newX, newY int
	var seat string
	for ! foundSeat && x >= 0 && x < len(seats[0]) && y >= 0 && y < len(seats) {
		newX = x + vector.x
		newY = y + vector.y
		seat = seatAt(newX, newY)
		if seat == "L" || seat == "#" {
			foundSeat = true
		}
		x = newX
		y = newY
	} 
	return seat
}

func isOccupied(x int, y int) bool {
	return seatAt(x, y) == "#"
}

func seatAt(x int, y int) string {
	if (x < 0 || x >= len(seats[0]) || y < 0 || y >= len(seats)) {
		return ""
	}
	return seats[y][x]
}

func getNumOccupied() int {
	var occupied int
	for _, row := range seats {
		for _, col := range row {
			if col == "#" {
				occupied++
			}
		}
	}

	return occupied
}