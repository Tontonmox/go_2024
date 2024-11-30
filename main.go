package main

import (
	"fmt"
	"go_2024/puzzles"
	"os"
)

func main() {
	day := os.Args[1]

	fmt.Println("This is day :" + day)

	puzzles.Day1()
}
