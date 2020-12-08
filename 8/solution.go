package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Command struct {
	operation string
	argument int
}

var accumulator int

func main() {
	lines := readInput()
	commands := parseCommands(lines)
	// partOne(commands)
	partTwo(commands)
	fmt.Println(accumulator)
}

func readInput() []string {
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

	return lines
}

func parseCommands(lines []string) []Command {
	var commands []Command
	for _, line := range lines {
		commands = append(commands, parseCommand(line))
	}
	return commands
}

func parseCommand(line string) Command {
	re := regexp.MustCompile(`(jmp|acc|nop) ([+|-][\d]+)`)
	matches := re.FindStringSubmatch(line)
	argument, _ := strconv.Atoi(matches[2]) 
	return Command{matches[1], argument}
}

func partOne(commands []Command) {
	var indicesOfCompletedCommands []int
	var i = 0
	for ! inArray(i, indicesOfCompletedCommands) {
		indicesOfCompletedCommands = append(indicesOfCompletedCommands, i)
		i = processCommand(commands[i], i)
	}
}

func partTwo(commands []Command) {
	var i int
	var j = 0
	for i != len(commands) {
		accumulator = 0
		newCommands := swapNopJmp(j, commands)
		i = 0
		var indicesOfCompletedCommands []int
		for ! inArray(i, indicesOfCompletedCommands) {
			if i == len(commands) {
				break
			}
			indicesOfCompletedCommands = append(indicesOfCompletedCommands, i)
			i = processCommand(newCommands[i], i)
		}
		j++
	}
}

func swapNopJmp(iteration int, commands []Command) []Command {
	var i int
	var newCommands []Command
	for _, command := range commands {
		newCommand := Command{command.operation, command.argument}
		if command.operation == "nop" {
			if i == iteration {
				newCommand.operation = "jmp"
			}
			i++
		} else if command.operation == "jmp" {
			if i == iteration {
				newCommand.operation = "nop"
			}
			i++
		}
		newCommands = append(newCommands, newCommand)
	}

	return newCommands
}

func processCommand(command Command, pos int) int {
	switch command.operation {
	case "acc":
		accumulator += command.argument
		pos += 1
	case "jmp":
		pos += command.argument
	case "nop":
		pos += 1
	}
	return pos
}

func inArray(needle int, haystack []int) bool {
	for _, el := range haystack {
		if (el == needle) {
			return true
		}
	}

	return false
}