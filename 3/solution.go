package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Grid struct {
	width  int
	height int
	rows   [][]string
}

type Toboggan struct {
	x, y int
}

type Slope struct {
	right, down int
}

func main() {
	lines := readInput()
	grid := buildGrid(lines)
	// trees := partOneTrees(grid)
	trees := partTwoTrees(grid)
	fmt.Println(trees)
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

func buildGrid(lines []string) Grid {
	var rows [][]string
	var width = len(lines[0])
	var height = len(lines)
	for _, line := range lines {
		var row []string
		for i := 0; i < len(line); i++ {
			row = append(row, string(line[i]))
		}
		rows = append(rows, row)
	}

	return Grid{width, height, rows}
}

func partOneTrees(grid Grid) int {
	slope := Slope{3, 1}
	return calculateTreesHit(grid, slope)
}

func partTwoTrees(grid Grid) int {
	slopes := [5]Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	var treesHit [5]int
	for i, slope := range slopes {
		trees := calculateTreesHit(grid, slope)
		treesHit[i] = trees
	}
	sum := 1
	for _, tree := range treesHit {
		sum = sum * tree
	}

	return sum
}

func calculateTreesHit(grid Grid, slope Slope) int {
	toboggan := Toboggan{}
	trees := 0
	for toboggan.y <= grid.height {
		toboggan.x, toboggan.y = moveToboggan(toboggan.x, toboggan.y, slope)
		square := getSquareAt(grid, toboggan.x, toboggan.y)
		if square == "#" {
			trees++
		}
	}

	return trees
}

func moveToboggan(tobogganX, tobogganY int, slope Slope) (int, int) {
	return tobogganX + slope.right, tobogganY + slope.down
}

func getSquareAt(grid Grid, x, y int) string {
	if y >= grid.height {
		return ""
	}
	xPos := x % grid.width

	return grid.rows[y][xPos]
}
