package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/unlikenesses/utils"
)

type Instruction struct {
	command string
	arg     int
}

var x, y int
var waypointX, waypointY int
var dir string

func main() {
	lines := utils.ReadInput()
	instructions := parseInstructions(lines)
	dir = "E"
	waypointX = 10
	waypointY = -1
	// val := partOne(instructions)
	val := partTwo(instructions)
	fmt.Println(val)
}

func parseInstructions(lines []string) []Instruction {
	var instructions []Instruction
	for _, line := range lines {
		instructions = append(instructions, parseInstruction(line))
	}
	return instructions
}

func parseInstruction(line string) Instruction {
	re := regexp.MustCompile(`(N|S|E|W|L|R|F)(\d+)`)
	matches := re.FindStringSubmatch(line)
	arg, _ := strconv.Atoi(matches[2])
	return Instruction{matches[1], arg}
}

func partOne(instructions []Instruction) float64 {
	for _, instruction := range instructions {
		performInstruction(instruction)
	}
	return manhattanDistance()
}

func performInstruction(instruction Instruction) {
	switch instruction.command {
	case "N":
		y -= instruction.arg
	case "S":
		y += instruction.arg
	case "E":
		x += instruction.arg
	case "W":
		x -= instruction.arg
	case "L":
		rotate("L", instruction.arg)
	case "R":
		rotate("R", instruction.arg)
	case "F":
		advance(instruction.arg)
	}
}

func rotate(direction string, amount int) {
	dirs := []string{"N", "E", "S", "W"}
	i := amount / 90
	if direction == "L" {
		i = i * -1
	}
	currentDirIndex := getIndex(dirs, dir)
	newDirIndex := currentDirIndex + i
	if newDirIndex >= len(dirs) {
		newDirIndex = newDirIndex % (len(dirs))
	} else if newDirIndex < 0 {
		newDirIndex = len(dirs) + newDirIndex
	}
	fmt.Println("Rotating", amount, "degrees", direction, "from", dir)
	dir = dirs[newDirIndex]
	fmt.Println("gives", dir)
}

func getIndex(dirs []string, dir string) int {
	for i, d := range dirs {
		if d == dir {
			return i
		}
	}
	return -1
}

func advance(distance int) {
	switch dir {
	case "N":
		y -= distance
	case "S":
		y += distance
	case "E":
		x += distance
	case "W":
		x -= distance
	}
}

func partTwo(instructions []Instruction) float64 {
	for _, instruction := range instructions {
		performInstructionPartTwo(instruction)
	}
	return manhattanDistance()
}

func performInstructionPartTwo(instruction Instruction) {
	switch instruction.command {
	case "N":
		waypointY -= instruction.arg
	case "S":
		waypointY += instruction.arg
	case "E":
		waypointX += instruction.arg
	case "W":
		waypointX -= instruction.arg
	case "L":
		rotateWaypoint("L", instruction.arg)
	case "R":
		rotateWaypoint("R", instruction.arg)
	case "F":
		advanceToWaypoint(instruction.arg)
	}
}

func rotateWaypoint(direction string, amount int) {
	relativeX, relativeY := getWaypointRelativeToShip()
	var newRelativeX, newRelativeY int
	if direction == "L" {
		amount = 360 - amount
	}
	numRotations := amount / 90
	for r := 0; r < numRotations; r++ {
		newRelativeX = relativeY * -1
		newRelativeY = relativeX
		relativeX = newRelativeX
		relativeY = newRelativeY
	}
	setNewWaypointPos(relativeX, relativeY)
}

func getWaypointRelativeToShip() (int, int) {
	return waypointX - x, waypointY - y
}

func setNewWaypointPos(relativeX, relativeY int) {
	waypointX = x + relativeX
	waypointY = y + relativeY
}

func advanceToWaypoint(distance int) {
	relativeX, relativeY := getWaypointRelativeToShip()
	for i := 0; i < distance; i++ {
		x += relativeX
		waypointX += relativeX
		y += relativeY
		waypointY += relativeY
	}
}

func manhattanDistance() float64 {
	return math.Abs(float64(x)) + math.Abs(float64(y))
}
