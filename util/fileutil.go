package util

import (
	"bufio"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLinesFromFile(path string) []string {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		lines = append(lines, string(line))
	}

	return lines
}

func ReadIntegersFromFile(path string) []int {
	lines := ReadLinesFromFile(path)

	integers := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		integer, err := strconv.Atoi(lines[i])
		Check(err)
		integers[i] = integer
	}

	return integers
}
