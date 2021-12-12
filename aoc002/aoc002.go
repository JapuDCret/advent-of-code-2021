package aoc002

import (
	util "adventofcode/util"
	"strconv"
	"strings"
)

func challengeA(input []string) int {
	steer := func(i string) (int, int) {
		inputs := strings.Split(i, " ")
		instruction := inputs[0]
		value, err := strconv.Atoi(inputs[1])
		util.Check(err)
		switch instruction {
		case "forward":
			return value, 0
		case "up":
			return 0, -value
		case "down":
			return 0, value
		default:
			return 0, 0
		}
	}

	horizontal := 0
	depth := 0
	for i := 0; i < len(input); i++ {
		h, d := steer(input[i])
		horizontal += h
		depth += d
	}
	return horizontal * depth
}

func challengeB(input []string) int {
	steer := func(i string) (int, int) {
		inputs := strings.Split(i, " ")
		instruction := inputs[0]
		value, err := strconv.Atoi(inputs[1])
		util.Check(err)
		switch instruction {
		case "forward":
			return value, 0
		case "up":
			return 0, -value
		case "down":
			return 0, value
		default:
			return 0, 0
		}
	}

	horizontal := 0
	depth := 0
	aim := 0
	for i := 0; i < len(input); i++ {
		h, a := steer(input[i])
		horizontal += h
		aim += a
		depth += aim * h
	}
	return horizontal * depth
}
