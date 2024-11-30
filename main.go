package main

import (
	"fmt"
	"go_2024/utils"
	"os"
)

func main() {
	fmt.Println("Hello World !!")

	day := os.Args[1]

	inputs := utils.ReadInput(day)

	for _, line := range inputs {
		fmt.Println(line)
	}

}
