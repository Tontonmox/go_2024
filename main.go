package main

import (
	"fmt"
	"go_2024/puzzles"
	"os"
)

func main() {

	testLaunch := false
	var day string

	for i, arg := range os.Args {
		if i == 1 {
			day = arg
		}
		if arg == "-test" {
			testLaunch = true
		}

	}

	switch day {
	case "1":
		puzzles.Day1(testLaunch)
	case "2":
		puzzles.Day2(testLaunch)
	case "3":
		puzzles.Day3(testLaunch)
	case "4":
		puzzles.Day4(testLaunch)
	case "5":
		puzzles.Day5(testLaunch)
	default:
		fmt.Println("Valeur incorrecte entrée pour le jour")
	}
}
