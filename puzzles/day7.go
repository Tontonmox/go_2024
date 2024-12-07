package puzzles

import (
	"fmt"
	"go_2024/utils"
	"strconv"
	"strings"
)

func Day7(isTest bool) {
	inputs := utils.ReadInputToSlice("7", isTest)

	res1 := 0
	res2 := 0

	for _, line := range inputs {
		parts := strings.Split(line, ":")
		//Number which has to be reached
		goal, _ := strconv.Atoi(parts[0])
		var elem []int

		//Setting up a slice of the numbers we can use
		for _, n := range strings.Split(parts[1][1:], " ") {
			num, _ := strconv.Atoi(n)
			elem = append(elem, num)
		}

		//Part 1
		combList := generateComb(len(elem)-1, 1)
		res1 += testComb(combList, elem, goal)

		//Part 2
		combList = generateComb(len(elem)-1, 2)
		res2 += testComb(combList, elem, goal)

	}
	fmt.Println(res1)
	fmt.Println(res2)

}

func generateComb(n int, part int) []string {
	i := 0
	var combList []string
	combList = append(combList, "")

	for i < n {
		var newList []string
		for _, s := range combList {
			newList = append(newList, s+"+")
			newList = append(newList, s+"*")
			if part == 2 {
				newList = append(newList, s+"|")
			}
		}
		combList = newList
		i++
	}

	return combList
}

func testComb(combList []string, elem []int, goal int) int {
	//Loop on all combinations until ok
	for _, c := range combList {
		total := 0
		for i := range len(elem) {
			if i == 0 {
				total = elem[i]
			} else {
				if c[i-1] == '+' {
					total += elem[i]
				} else if c[i-1] == '*' {
					total *= elem[i]
				} else {
					strTotal := strconv.Itoa(total)
					strElem := strconv.Itoa(elem[i])
					newStr := strTotal + strElem
					total, _ = strconv.Atoi(newStr)

				}
			}
		}
		if total == goal {
			return goal
		}
	}
	return 0
}
