package puzzles

import (
	"fmt"
	"go_2024/utils"
	"slices"
)

func Day8(isTest bool) {
	inputs := utils.ReadInputToSlice("8", isTest)

	xMax := len(inputs[0]) - 1
	yMax := len(inputs) - 1

	fmt.Printf("Coordonnées amx %d-%d \n", xMax, yMax)

	antList := make(map[rune][]coord)

	//For each type on antenna make a list of existing coords
	for cY, line := range inputs {
		for cX, char := range line {
			if char != '.' {
				antList[char] = append(antList[char], coord{cX, cY})
			}
		}
	}

	var antinodes1 []coord
	var antinodes2 []coord

	//Loop on antenna types, then for each type check every possible couple
	for _, antList := range antList {
		antNum := len(antList)
		for i, a1 := range antList {
			numTest := antNum - (i + 1)
			for mod := range numTest {
				a2 := antList[i+mod+1]

				for _, an := range testAntennas(a1, a2, xMax, yMax) {
					if !slices.Contains(antinodes1, an) {
						antinodes1 = append(antinodes1, an)
					}
				}

				for _, an := range testAntennasLooped(a1, a2, xMax, yMax) {
					if !slices.Contains(antinodes2, an) {
						antinodes2 = append(antinodes2, an)
					}
				}

			}
		}
	}

	fmt.Println(len(antinodes1))
	fmt.Println(len(antinodes2))

}

func testAntennas(a1 coord, a2 coord, xMax int, yMax int) []coord {
	var validAntinodes []coord

	//First antinode
	x1 := a2.x + (a2.x - a1.x)
	y1 := a2.y + (a2.y - a1.y)

	if x1 >= 0 && x1 <= xMax && y1 >= 0 && y1 <= yMax {
		validAntinodes = append(validAntinodes, coord{x1, y1})
	}

	//Second antinode
	x2 := a1.x + (a1.x - a2.x)
	y2 := a1.y + (a1.y - a2.y)

	if x2 >= 0 && x2 <= xMax && y2 >= 0 && y2 <= yMax {
		validAntinodes = append(validAntinodes, coord{x2, y2})
	}

	//fmt.Printf("Antennas %v - %v have antinodes %d-%d and %d-%d\n", a1, a2, x1, y1, x2, y2)

	return validAntinodes
}

func testAntennasLooped(a1 coord, a2 coord, xMax int, yMax int) []coord {
	var validAntinodes []coord

	//First antinode
	mod := 0
	for {
		x1 := a2.x + (a2.x-a1.x)*mod
		y1 := a2.y + (a2.y-a1.y)*mod

		if x1 >= 0 && x1 <= xMax && y1 >= 0 && y1 <= yMax {
			validAntinodes = append(validAntinodes, coord{x1, y1})
		} else {
			break
		}
		mod++
	}

	//Second antinode
	mod = 0
	for {
		x2 := a1.x + (a1.x-a2.x)*mod
		y2 := a1.y + (a1.y-a2.y)*mod

		if x2 >= 0 && x2 <= xMax && y2 >= 0 && y2 <= yMax {
			validAntinodes = append(validAntinodes, coord{x2, y2})
		} else {
			break
		}
		mod++
	}
	//fmt.Printf("Antennas %v - %v have antinodes %d-%d and %d-%d\n", a1, a2, x1, y1, x2, y2)

	return validAntinodes
}
