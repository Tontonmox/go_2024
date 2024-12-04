package puzzles

import (
	"fmt"
	"go_2024/utils"
)

func Day4(isTest bool) {
	inputs := utils.ReadInputToSlice("4", isTest)

	res1 := 0
	l := len(inputs[0])
	h := len(inputs)

	//PART 1
	//First search : vertical
	for i := range l {
		for j := range h - 3 {
			word := string(inputs[j][i]) + string(inputs[j+1][i]) + string(inputs[j+2][i]) + string(inputs[j+3][i])
			if word == "XMAS" || word == "SAMX" {
				res1++
			}
		}
	}
	//Second search : Horizontal
	for i := range l - 3 {
		for j := range h {
			word := string(inputs[j][i]) + string(inputs[j][i+1]) + string(inputs[j][i+2]) + string(inputs[j][i+3])
			if word == "XMAS" || word == "SAMX" {
				res1++
			}
		}
	}
	//3rd search : Diagonal (from top left)
	for i := range l - 3 {
		for j := range h - 3 {
			word := string(inputs[j][i]) + string(inputs[j+1][i+1]) + string(inputs[j+2][i+2]) + string(inputs[j+3][i+3])
			if word == "XMAS" || word == "SAMX" {
				res1++
			}
		}
	}
	//4th search : Diagonal (from top left)
	for i := 3; i < l; i++ {
		for j := range h - 3 {
			word := string(inputs[j][i]) + string(inputs[j+1][i-1]) + string(inputs[j+2][i-2]) + string(inputs[j+3][i-3])
			if word == "XMAS" || word == "SAMX" {
				res1++
			}
		}
	}
	fmt.Println(res1)

	//PART 2
	//Search all "A" characters, excluding first/last rows/columns
	//For each A check the two top corners : they must be M or S, and their opposite corner must be the other value
	res2 := 0
	for i := 1; i < l-1; i++ {
		for j := 1; j < h-1; j++ {
			if inputs[j][i] == 'A' &&
				((inputs[j-1][i-1] == 'M' && inputs[j+1][i+1] == 'S') || (inputs[j-1][i-1] == 'S' && inputs[j+1][i+1] == 'M')) &&
				((inputs[j-1][i+1] == 'M' && inputs[j+1][i-1] == 'S') || (inputs[j-1][i+1] == 'S' && inputs[j+1][i-1] == 'M')) {
				res2++
			}
		}
	}
	fmt.Println(res2)
}
