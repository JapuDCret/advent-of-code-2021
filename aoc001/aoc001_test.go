package aoc001

import (
	util "adventofcode/util"
	"fmt"
	"testing"
)

func TestChallengeA_WithPrecursorInput(t *testing.T) {
	input := [...]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	result := challengeA(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 7 {
		t.Errorf("Expected %d, received %d", 7, result)
	}
}

func TestChallengeA_WithRealInput(t *testing.T) {
	input := util.ReadIntegersFromFile("challengeAInput.txt")

	result := challengeA(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 1342 {
		t.Errorf("Expected %d, received %d", 1342, result)
	}
}
