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

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func main() {
	lines := readInput()
	passports := readPassports(lines)
	// valid := getNumValidPassports(passports, 1)
	valid := getNumValidPassports(passports, 2)
	fmt.Println(valid)
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

func readPassports(lines []string) []Passport {
	var passports []Passport
	passport := Passport{"", "", "", "", "", "", "", ""}
	for _, line := range lines {
		if line == "" {
			passports = append(passports, passport)
			passport = Passport{"", "", "", "", "", "", "", ""}
		}
		pairs := strings.Fields(line)
		for _, pair := range pairs {
			key, value := parseField(pair)
			switch key {
			case "byr":
				passport.byr = value
			case "iyr":
				passport.iyr = value
			case "eyr":
				passport.eyr = value
			case "hgt":
				passport.hgt = value
			case "hcl":
				passport.hcl = value
			case "ecl":
				passport.ecl = value
			case "pid":
				passport.pid = value
			case "cid":
				passport.cid = value
			}
		}
	}
	passports = append(passports, passport)

	return passports
}

func parseField(pair string) (string, string) {
	re := regexp.MustCompile(`(cid|iyr|pid|eyr|ecl|hcl|byr|hgt):([^\s]+)`)
	matches := re.FindStringSubmatch(pair)

	return matches[1], matches[2]
}

func getNumValidPassports(passports []Passport, puzzlePart int) int {
	valid := 0
	for _, passport := range passports {
		if isValid(passport, puzzlePart) {
			valid++
		}
	}

	return valid
}

func isValid(passport Passport, puzzlePart int) bool {
	if puzzlePart == 1 {
		return passport.byr != "" && passport.ecl != "" && passport.eyr != "" && passport.hcl != "" && passport.hgt != "" && passport.iyr != "" && passport.pid != ""
	}

	if passport.byr == "" || passport.ecl == "" || passport.eyr == "" || passport.hcl == "" || passport.hgt == "" || passport.iyr == "" || passport.pid == "" {
		return false
	}

	byr, _ := strconv.Atoi(passport.byr)
	if byr < 1920 || byr > 2002 {
		return false
	}
	iyr, _ := strconv.Atoi(passport.iyr)
	if iyr < 2010 || iyr > 2020 {
		return false
	}
	eyr, _ := strconv.Atoi(passport.eyr)
	if eyr < 2020 || eyr > 2030 {
		return false
	}
	if !isValidHeight(passport.hgt) {
		return false
	}
	if !isValidHairColour(passport.hcl) {
		return false
	}
	if !isValidEyeColour(passport.ecl) {
		return false
	}
	if !isValidPassportID(passport.pid) {
		return false
	}

	return true
}

func isValidHeight(hgt string) bool {
	re := regexp.MustCompile(`([\d]+)(cm|in)`)
	matches := re.FindStringSubmatch(hgt)

	if len(matches) == 0 {
		return false
	}
	height, _ := strconv.Atoi(matches[1])
	if matches[2] == "cm" {
		return height >= 150 && height <= 193
	}

	return height >= 59 && height <= 76
}

func isValidHairColour(hcl string) bool {
	re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	matches := re.FindStringSubmatch(hcl)

	return len(matches) > 0
}

func isValidEyeColour(ecl string) bool {
	re := regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
	matches := re.FindStringSubmatch(ecl)

	return len(matches) > 0
}

func isValidPassportID(pid string) bool {
	re := regexp.MustCompile(`^\d{9}$`)
	matches := re.FindStringSubmatch(pid)

	return len(matches) > 0
}
