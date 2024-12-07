package utils

import (
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInputToSlice(day string, isTest bool) []string {
	var path string
	if isTest {
		path = "./samples/" + day
	} else {
		path = "./inputs/" + day
	}

	file, err := os.Open(path)

	Check(err)

	var fileText []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileText = append(fileText, scanner.Text())
	}

	return fileText
}

// Abs but with int
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Count elements in a slice meeting a condition
// Param 1 : slice
// Param 2 : func returning a bool
// For each element, if the function returns true, +1 in the counter
func CountInSlice[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
