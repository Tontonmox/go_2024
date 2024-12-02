package puzzles

import (
	"fmt"
	"go_2024/utils"
	"slices"
	"strconv"
	"strings"
)

func Day2(isTest bool) {
	inputs := utils.ReadInputToSlice("2", isTest)

	nbSafe1 := 0
	nbSafe2 := 0

	for _, report := range inputs {
		//Convert each report into a slice of ints
		levelString := strings.Split(report, " ")
		levels := make([]int, len(levelString))

		for i, s := range levelString {
			levels[i], _ = strconv.Atoi(s)
		}

		// Testing each list either in increase or decrease direction
		errorIndex := checkForError(levels, 1)
		errorIndex2 := checkForError(levels, -1)

		if errorIndex == -1 || errorIndex2 == -1 {
			nbSafe1++
			nbSafe2++
		} else {
			//Part 2 CRADO : if there was an error, make 4 different lists, by removing each time one of the two numbers creating the errors, on each direction
			//If one on these lists is safe then result is ok for part 2
			var newList1 []int
			newList1 = append(newList1, levels[:errorIndex]...)
			newList1 = append(newList1, levels[errorIndex+1:]...)

			var newList2 []int
			newList2 = append(newList2, levels[:errorIndex+1]...)
			newList2 = append(newList2, levels[errorIndex+2:]...)

			var newList3 []int
			newList3 = append(newList3, levels[:errorIndex2]...)
			newList3 = append(newList3, levels[errorIndex2+1:]...)

			var newList4 []int
			newList4 = append(newList4, levels[:errorIndex2+1]...)
			newList4 = append(newList4, levels[errorIndex2+2:]...)

			if (checkForError(newList1, 1) == -1) || (checkForError(newList2, 1) == -1) ||
				(checkForError(newList3, 1) == -1) || (checkForError(newList4, 1) == -1) ||
				(checkForError(newList1, -1) == -1) || (checkForError(newList2, -1) == -1) ||
				(checkForError(newList3, -1) == -1) || (checkForError(newList4, -1) == -1) {
				nbSafe2++
			}
		}
	}

	fmt.Println(nbSafe1)
	fmt.Println(nbSafe2)
}

// Check a list of levels for "safeness" and return the index of the first error
func checkForError(levels []int, mod int) int {
	errorIndex := -1

	for i := range len(levels) - 1 {
		//List the 3 ok values for next value and check if it's in
		okValues := []int{levels[i] + mod, levels[i] + 2*mod, levels[i] + 3*mod}
		if !slices.Contains(okValues, levels[i+1]) {
			errorIndex = i
			break
		}
	}

	return errorIndex
}
