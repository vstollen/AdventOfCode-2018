package main

import "fmt"

type group []string

func CountYesAnsweredQuestions() {
	data := ReadData("day_6.txt")
	groups := parseGroups(data)

	sumOfCounts := 0

	for _, group := range groups {
		answered := map[rune]bool{}

		for _, line := range group {
			for _, answer := range line {
				answered[answer] = true
			}
		}

		sumOfCounts += len(answered)
	}

	fmt.Printf("Sum of answered questions: %v\n", sumOfCounts)
}

func parseGroups(data []string) []group {
	var groups []group

	for _, line := range data {
		if len(groups) == 0 || line == "" {
			groups = append(groups, group{})

			if line == "" {
				continue
			}
		}

		groups[len(groups)-1] = append(groups[len(groups)-1], line)
	}

	return groups
}
