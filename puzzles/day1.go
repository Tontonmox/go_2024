package puzzles

import (
	"fmt"
	"go_2024/utils"
	"sort"
	"strconv"
	"strings"
)

func Day1(isTest bool) {
	d1part1(isTest)
}

func d1part1(isTest bool) {
	inputs := utils.ReadInputToSlice("1", isTest)

	var sliceA []int
	var sliceB []int

	for _, line := range inputs {
		//Get each of the 2 values as an int and store them into a slice
		values := strings.Split(line, "   ")

		//First list
		valA, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Println("Non numeric value when slicing  " + line)
			panic(err)
		}
		sliceA = append(sliceA, valA)

		//Second list
		valB, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("Non numeric value when slicing   " + line)
			panic(err)
		}
		sliceB = append(sliceB, valB)
	}

	//Sort the two slices to compare each elements
	sort.Ints(sliceA)
	sort.Ints(sliceB)

	res_part1 := 0
	//Parse the 2 slices together and calculate difference
	for i, a := range sliceA {
		res_part1 += utils.AbsInt(a - sliceB[i])
	}
	fmt.Printf("Result day 1 part 1 : %d\n", res_part1)
	d1part2(sliceA, sliceB)
}

func d1part2(a []int, b []int) {
	res_part2 := 0
	for _, a := range a {
		nbOcc := utils.CountInSlice(
			b,
			func(x int) bool {
				return x == a
			},
		)
		res_part2 += a * nbOcc
	}
	fmt.Printf("Result day 1 part 2 : %d\n", res_part2)
}
