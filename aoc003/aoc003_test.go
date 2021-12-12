package aoc003

import (
	util "adventofcode/util"
	"fmt"
	"testing"
)

func TestChallengeA_WithPrecursorInput(t *testing.T) {
	input := [...]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	result := challengeA(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 198 {
		t.Errorf("Expected %d, received %d", 198, result)
	}
}

func TestChallengeA_WithRealInput(t *testing.T) {
	input := util.ReadLinesFromFile("challengeAInput.txt")

	result := challengeA(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 845186 {
		t.Errorf("Expected %d, received %d", 845186, result)
	}
}

func TestChallengeB_WithPrecursorInput(t *testing.T) {
	input := [...]string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

	result := challengeB(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 230 {
		t.Errorf("Expected %d, received %d", 230, result)
	}
}

func TestChallengeB_WithRealInput(t *testing.T) {
	input := util.ReadLinesFromFile("challengeAInput.txt")

	result := challengeB(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 4636702 {
		t.Errorf("Expected %d, received %d", 4636702, result)
	}
}
