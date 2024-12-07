package puzzles

import (
	"fmt"
	"go_2024/utils"
	"slices"
	"strings"
)

type coord struct {
	x int
	y int
}

type state struct {
	pos coord
	dir coord
}

func Day6(isTest bool) {
	grid := utils.ReadInputToSlice("6", isTest)
	xMax := len(grid[0]) - 1
	yMax := len(grid) - 1
	curDir := coord{0, -1}

	x, y := -1, -1
	//Find starting position
	for i, line := range grid {
		x = strings.Index(line, "^")
		if x != -1 {
			y = i
			break
		}
	}

	var initPath []state
	//Replace currentPos by X in the grid
	grid[y] = utils.ReplaceAtIndex(grid[y], 'X', x)

	//Initialise path with first value
	initPath = append(initPath, state{coord{x, y}, curDir})

	testGrid := make([]string, len(grid))
	copy(testGrid, grid)

	testPath := make([]state, len(initPath))
	copy(testPath, initPath)

	testPath, isLoop := simulate(testGrid, testPath, x, y, xMax, yMax, curDir)
	_ = isLoop

	//Once out of the loop, count the X in the grid
	fmt.Println(countX(testGrid))

	//Test every possible new blocker and check if it makes a loop
	loopCount := 0
	var tested []coord

	for n, p := range testPath {
		if n == 0 {
			continue
		}
		i := p.pos.y
		j := p.pos.x
		toTest := coord{j, i}
		if slices.Contains(tested, toTest) {
			continue
		} else {
			tested = append(tested, toTest)
		}

		newGrid := make([]string, len(grid))
		copy(newGrid, grid)

		newPath := make([]state, len(initPath))
		copy(newPath, initPath)

		newGrid[i] = utils.ReplaceAtIndex(newGrid[i], '#', j)

		_, isLoop := simulate(newGrid, newPath, x, y, xMax, yMax, curDir)

		if isLoop {
			loopCount++
		}
	}

	fmt.Println(loopCount)

}

func simulate(grid []string, path []state, x int, y int, xMax int, yMax int, curDir coord) ([]state, bool) {
	//Movement loop
	isLoop := false
	for {
		//Compute next position
		nX := x + curDir.x
		nY := y + curDir.y

		//If next position outside of grid, break the loop
		if nX < 0 || nX > xMax || nY < 0 || nY > yMax {
			break
		} else {
			//Check if next pos is an obstacle. If so, rotate and calculate new next position
			if grid[nY][nX] == '#' {
				curDir = nextDir(curDir)
				nX = x + curDir.x
				nY = y + curDir.y

				//Check (again) if it leads outside of the grid
				if nX < 0 || nX > xMax || nY < 0 || nY > yMax {
					break
				}
			}
			//Same thing in case U turn necessary
			if grid[nY][nX] == '#' {
				curDir = nextDir(curDir)
				nX = x + curDir.x
				nY = y + curDir.y

				//Check (again) if it leads outside of the grid
				if nX < 0 || nX > xMax || nY < 0 || nY > yMax {
					break
				}
			}

			//If not, make the move and mark the new spot with X
			x = nX
			y = nY
			grid[y] = utils.ReplaceAtIndex(grid[y], 'X', x)
			newState := state{coord{x, y}, curDir}

			if slices.Contains(path, newState) {
				isLoop = true
				break
			} else {
				path = append(path, newState)
			}
		}
	}
	return path, isLoop
}

func nextDir(curDir coord) coord {
	directions := map[coord]coord{
		{0, -1}: {1, 0},
		{1, 0}:  {0, 1},
		{0, 1}:  {-1, 0},
		{-1, 0}: {0, -1}}

	return directions[curDir]
}

func countX(grid []string) int {
	res := 0
	for _, line := range grid {
		for _, char := range line {
			if char == 'X' {
				res++
			}
		}
	}

	return res
}
