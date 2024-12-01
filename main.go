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

	fmt.Println("This is day :" + day)

	puzzles.Day1(testLaunch)
}
