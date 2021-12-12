package aoc001

func challengeA(input []int) int {
	counter := -1
	previous := -1
	for i := 0; i < len(input); i++ {
		if input[i] > previous {
			counter++
		}
		previous = input[i]
	}
	return counter
}
