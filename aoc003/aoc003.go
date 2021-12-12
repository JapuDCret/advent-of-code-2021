package aoc003

import (
	"adventofcode/util"
	"fmt"
)

func challengeA(input []string) int {
	bitCount := make([]int8, len(input[0]))

	for i := 0; i < len(input); i++ {
		for k := 0; k < len(bitCount); k++ {
			if input[i][k] == '0' {
				bitCount[k]--
			} else {
				bitCount[k]++
			}
		}
	}

	gamma := 0
	epsilon := 0

	fmt.Printf("bitCount = %v\n", bitCount)

	factor := 1
	for i := len(bitCount) - 1; i >= 0; i-- {
		if bitCount[i] > 0 {
			gamma += factor
		} else {
			epsilon += factor
		}

		factor *= 2
	}

	fmt.Printf("gamma = %v\n", gamma)
	fmt.Printf("epsilon = %v\n", epsilon)

	return gamma * epsilon
}

func challengeB(input []string) int {
	bits := len(input[0])

	// oxygen generator rating
	ogr := input
	// CO2 scrubber rating
	co2sr := input

	filter := func(input []string, index int, filterOutZeroes bool) []string {
		result := []string{}

		fmt.Printf("filter(): input = %v\n", input)

		if len(input) == 1 {
			return input
		}

		for i := range input {
			if input[i][index] == '1' && filterOutZeroes {
				result = append(result, input[i])
			} else if input[i][index] == '0' && !filterOutZeroes {
				result = append(result, input[i])
			}
		}

		fmt.Printf("filter(): result = %v\n", result)

		return result
	}

	bitCount := func(input []string, index int) int {
		result := 0
		for i := 0; i < len(input); i++ {
			if input[i][index] == '0' {
				result--
			} else {
				result++
			}
		}
		return result
	}

	for i := 0; i < bits; i++ {
		ogr = filter(ogr, i, bitCount(ogr, i) >= 0)
		co2sr = filter(co2sr, i, bitCount(co2sr, i) < 0)
	}

	fmt.Printf("challengeB(): ogr = %v\n", ogr)
	fmt.Printf("challengeB(): co2sr = %v\n", co2sr)

	o := util.BitStringToInt(ogr[0])
	c := util.BitStringToInt(co2sr[0])

	fmt.Printf("challengeB(): o = %v\n", o)
	fmt.Printf("challengeB(): c = %v\n", c)

	return o * c
}
