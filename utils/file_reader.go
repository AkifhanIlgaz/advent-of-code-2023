package utils

import (
	"bufio"
	"os"
)

func ReadLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
