package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/unlikenesses/utils"
)

var timestamp int
var busIds []string

func main() {
	lines := utils.ReadInput()
	timestamp, _ = strconv.Atoi(lines[0])
	busIds = strings.Split(lines[1], ",")
	// val := partOne()
	val := partTwo()
	fmt.Println(val)
}

func partOne() int {
	nextBuses := make(map[int]int)
	for _, busId := range busIds {
		if busId != "x" {
			id, _ := strconv.Atoi(busId)
			mod := timestamp % id
			nextBus := timestamp - mod + id
			nextBuses[nextBus] = id
		}
	}
	ts, busId := getFirstEl(nextBuses)
	return (ts - timestamp) * busId
}

func getFirstEl(buses map[int]int) (int, int) {
	// https://stackoverflow.com/a/23332089
	keys := make([]int, len(buses))
	i := 0
	for k := range buses {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	var ts, busId int
	for _, k := range keys {
		ts = k
		busId = buses[k]
		break
	}
	return ts, busId
}

func partTwo() int {
	// Brute force: after 20 mins still no result...
	var found bool
	ts := 100000000000000
	for !found {
		ts++
		found = isValid(ts)
	}

	return ts
}

func isValid(ts int) bool {
	for i, busId := range busIds {
		if busId != "x" {
			id, _ := strconv.Atoi(busId)
			if (ts+i)%id != 0 {
				return false
			}
		}
	}
	return true
}
