package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	container string
	containees []Containee
}

type Containee struct {
	qty int
	container string
}

var rules []Rule

func main() {
	lines := readInput()
	rules = parseRules(lines)
	parents := getParentBags("shiny gold")
	fmt.Println(len(parents))
	numBags := getNumBagsInside("shiny gold")
	fmt.Println(numBags)
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

func parseRules(lines []string) []Rule {
	var rules []Rule
	for _, line := range lines {
		rules = append(rules, parseRule(line))
	}

	return rules
}

func parseRule(line string) Rule {
	container, containees := splitRule(line)

	return Rule{container, containees}
}

func splitRule(line string) (string, []Containee) {
	re := regexp.MustCompile(`([\w\s]+)(bags contain) ([\w\s,]+)`)
	matches := re.FindStringSubmatch(line)

	containees := splitContainees(matches[3])

	return strings.TrimSpace(matches[1]), containees
}

func splitContainees(containeeString string) []Containee {
	re := regexp.MustCompile(`(\d+)([\w\s]+) bag`)
	matches := re.FindAllStringSubmatch(containeeString, -1)
	var containees []Containee 
	if len(matches) > 0 {
		for _, match := range matches {
			qty, _ := strconv.Atoi(match[1])
			containee := Containee{qty, strings.TrimSpace(match[2])}
			containees = append(containees, containee)
		}
	}

	return containees
}

func getParentBags(bag string) []string {
	var parentBags []string
	for _, rule := range rules {
		if inArray(bag, allContents(rule.containees)) {
			parentBags = append(parentBags, rule.container)
		}
	}
	return parentBags
}

func allContents(containees []Containee) []string {
	var total []string
	for _, containee := range containees {
		total = append(total, containee.container)
		total = append(total, allContents(getContainees(containee.container))...)
	}
	return total
}

func getContainees(bag string) []Containee {
	for _, rule := range rules {
		if rule.container == bag {
			return rule.containees
		}
	}

	return nil
}

func getNumBagsInside(bag string) int {
	containees := getContainees(bag)
	sum := 0
	for _, containee := range containees {
		sum += containee.qty + containee.qty * getNumBagsInside(containee.container)
	}
	return sum
}

func inArray(needle string, haystack []string) bool {
	for _, el := range haystack {
		if (el == needle) {
			return true
		}
	}

	return false
}