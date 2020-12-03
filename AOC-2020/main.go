package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to AdventOfCode 2020!\n" +
		"Please select a day (1-3):")

	var day int

	_, _ = fmt.Scanln(&day)


	switch day {
	case 1:
		ExpenseReport()
	case 2:
		CountValidPasswords()
	case 3:
		CountTrees()
	}
}
