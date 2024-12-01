package puzzles

import (
	"fmt"
	"go_2024/utils"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	part1()
}

func part1() {
	fmt.Println("Day 1 part 1")

	inputs := utils.ReadInputInSlice("1")

	var sliceA []int
	var sliceB []int

	for _, line := range inputs {
		//Get each of the 2 values as an int and store them into a slice
		values := strings.Split(line, "   ")

		//Valeur 1
		valA, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Println("Valeur non numérique suite au slice de  " + line)
			panic(err)
		}
		sliceA = append(sliceA, valA)

		//Valeur 2
		valB, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("Valeur non numérique suite au slice de  " + line)
			panic(err)
		}
		sliceB = append(sliceB, valB)
	}

	//Tri des deux slice pour faire la comparaison
	sort.Ints(sliceA)
	sort.Ints(sliceB)

	res_part1 := 0
	//Parcours des 2 slices en simultané et calcul du résultat
	for i, a := range sliceA {
		res_part1 += utils.AbsInt(a - sliceB[i])
	}
	fmt.Printf("Le résultat pour la partie 1 est : %d\n", res_part1)
	part2(sliceA, sliceB)
}

func part2(a []int, b []int) {
	fmt.Println("Day 1 part 2")

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
	fmt.Printf("Le résultat pour la partie 2 est : %d\n", res_part2)
}
