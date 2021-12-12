package aoc012

import (
	util "adventofcode/util"
	"fmt"
	"testing"
)

func TestChallengeA_WithPrecursorInput(t *testing.T) {
	input := [...]string{"start-A", "start-b", "A-c", "A-b", "b-d", "A-end", "b-end"}

	result := challengeA(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 10 {
		t.Errorf("Expected %d, received %d", 10, result)
	}
}

func TestChallengeA_WithSecondPrecursorInput(t *testing.T) {
	input := [...]string{"dc-end", "HN-start", "start-kj", "dc-start", "dc-HN", "LN-dc", "HN-end", "kj-sa", "kj-HN", "kj-dc"}

	result := challengeA(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 19 {
		t.Errorf("Expected %d, received %d", 19, result)
	}
}

func TestChallengeA_WithRealInput(t *testing.T) {
	input := util.ReadLinesFromFile("challengeAInput.txt")

	result := challengeA(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 4775 {
		t.Errorf("Expected %d, received %d", 4775, result)
	}
}

func TestChallengeB_WithPrecursorInput(t *testing.T) {
	input := [...]string{"start-A", "start-b", "A-c", "A-b", "b-d", "A-end", "b-end"}

	result := challengeB(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 36 {
		t.Errorf("Expected %d, received %d", 36, result)
	}
}

func TestChallengeB_WithRealInput(t *testing.T) {
	input := util.ReadLinesFromFile("challengeAInput.txt")

	result := challengeB(input[:])

	fmt.Printf("result = %v\n", result)

	if result != 152480 {
		t.Errorf("Expected %d, received %d", 152480, result)
	}
}
