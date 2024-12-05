package puzzles

import (
	"fmt"
	"go_2024/utils"
	"reflect"
	"strconv"
	"strings"
)

func Day5(isTest bool) {
	inputs := utils.ReadInputToSlice("5", isTest)

	var rules [101][101]bool
	var pages []string

	ruleEnd := false
	for _, line := range inputs {
		if line == "" {
			ruleEnd = true
			continue
		}

		//Loading rules into array
		if !ruleEnd {

			r := strings.Split(line, "|")
			r1, _ := strconv.Atoi(r[0])
			r2, _ := strconv.Atoi(r[1])
			rules[r1][r2] = true

		} else {
			//Loading pages into array
			pages = append(pages, line)
		}
	}

	res1 := 0
	res2 := 0
	for _, page := range pages {
		nums := strings.Split(page, ",")
		midIndex := len(nums) / 2

		newNums := sortPage(nums, rules)

		if reflect.DeepEqual(nums, newNums) {
			r, _ := strconv.Atoi(nums[midIndex])
			res1 += r

		} else {
			//For part 2 only : sort correctly the incorrect page and add the new middle number
			r, _ := strconv.Atoi(newNums[midIndex])
			res2 += r
		}
	}
	fmt.Println(res1)
	fmt.Println(res2)
}

func sortPage(page []string, rules [101][101]bool) []string {
	sorted := make([]string, len(page))
	copy(sorted, page)

	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {

			num1, _ := strconv.Atoi(sorted[j])
			num2, _ := strconv.Atoi(sorted[j+1])

			if rules[num2][num1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	return sorted
}
