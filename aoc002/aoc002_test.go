package aoc002

import (
	util "adventofcode/util"
	"fmt"
	"testing"
)

func TestChallengeA_WithPrecursorInput(t *testing.T) {
	input := [...]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	result := challengeA(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 150 {
		t.Errorf("Expected %d, received %d", 150, result)
	}
}

func TestChallengeA_WithRealInput(t *testing.T) {
	input := util.ReadLinesFromFile("challengeAInput.txt")

	result := challengeA(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 1727835 {
		t.Errorf("Expected %d, received %d", 1727835, result)
	}
}

func TestChallengeB_WithPrecursorInput(t *testing.T) {
	input := [...]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	result := challengeB(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 900 {
		t.Errorf("Expected %d, received %d", 900, result)
	}
}

func TestChallengeB_WithRealInput(t *testing.T) {
	input := util.ReadLinesFromFile("challengeAInput.txt")

	result := challengeB(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 1544000595 {
		t.Errorf("Expected %d, received %d", 1544000595, result)
	}
}
