package main

import (
	"fmt"
	"regexp"
	"strconv"
	"github.com/unlikenesses/utils"
)

type Policy struct {
	min    int
	max    int
	letter string
}

type Password struct {
	policy   Policy
	password string
}

func main() {
	lines := utils.ReadInput()
	passwords := parsePasswords(lines)
	// validPasswords := getValidPasswordsPartOne(passwords)
	validPasswords := getValidPasswordsPartTwo(passwords)
	fmt.Println(len(validPasswords))
}

func parsePasswords(lines []string) []Password {
	var passwords []Password
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			min, _ := strconv.Atoi(matches[1])
			max, _ := strconv.Atoi(matches[2])
			policy := Policy{min, max, matches[3]}
			password := Password{policy, matches[4]}
			passwords = append(passwords, password)
		} else {
			fmt.Println("No matches for line ", line)
		}
	}

	return passwords
}

func getValidPasswordsPartOne(passwords []Password) []Password {
	var valid []Password
	for _, password := range passwords {
		policy := password.policy
		re := regexp.MustCompile(policy.letter)
		matches := re.FindAllStringIndex(password.password, -1)
		letterCount := len(matches)
		if letterCount >= policy.min && letterCount <= policy.max {
			valid = append(valid, password)
		}
	}

	return valid
}

func getValidPasswordsPartTwo(passwords []Password) []Password {
	var valid []Password
	for _, password := range passwords {
		policy := password.policy
		if (string(password.password[policy.min-1]) == policy.letter || string(password.password[policy.max-1]) == policy.letter) && password.password[policy.min-1] != password.password[policy.max-1] {
			valid = append(valid, password)
		}
	}

	return valid
}
