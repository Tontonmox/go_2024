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
