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

func ReadInputInSlice(day string) []string {
	path := "./inputs/" + day
	file, err := os.Open(path)

	Check(err)

	var fileText []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileText = append(fileText, scanner.Text())
	}

	return fileText
}

// Fonction valeur absolue pour des int
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Fonction qui compte les éléments dans une slice répondant à une condition
func CountInSlice[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}
