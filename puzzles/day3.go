package puzzles

import (
	"fmt"
	"go_2024/utils"
	"strconv"
	"strings"
)

func Day3(isTest bool) {
	inputs := utils.ReadInputToSlice("3", isTest)
	result1 := 0
	result2 := 0

	//Part 1
	for _, input := range inputs {
		result1 += countMul(input)
	}
	fmt.Println(result1)

	//Part 2
	for _, input := range inputs {
		//fmt.Println("Debut input")
		//fmt.Println(input)
		for {
			//Search for first dont't()
			posDont := strings.Index(input, "don't()")
			if posDont == -1 {
				break
			}
			//Search for first do(), if not, remove everything after don't()
			posDo := strings.Index(input[posDont:], "do()")
			if posDo == -1 {
				input = input[:posDont]
				//fmt.Println(input)
				break
			}
			posDo += posDont
			//Replace everything between don't() and do(), including the instructions, by a *
			input = input[:posDont] + "*" + input[posDo+4:]
			//fmt.Println(input)
		}
		result2 += countMul(input)
	}
	fmt.Println(result2)

}

func countMul(input string) int {
	result := 0
	//Split on mul, and only strings starting with ( andcontaining a ), keep the part between parenthesis
	split1 := strings.Split(input, "mul")
	var input2 []string
	for _, part := range split1 {
		if part[0] == '(' {
			posEnd := strings.Index(part, ")")
			if posEnd != -1 {
				input2 = append(input2, part[1:posEnd])
			}
		}
	}

	//For each result, split on a ",", check that there is only two parts and that each part is numeric
	for _, part := range input2 {
		splitted := strings.Split(part, ",")
		if len(splitted) == 2 {
			num1, err1 := strconv.Atoi(splitted[0])
			num2, err2 := strconv.Atoi(splitted[1])
			if err1 == nil && err2 == nil {
				result += num1 * num2
			}
		}
	}
	return result
}
