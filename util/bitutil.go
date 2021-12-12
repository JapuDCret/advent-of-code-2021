package util

func BitStringToInt(bitString string) int {
	result := 0

	factor := 1
	for i := len(bitString) - 1; i >= 0; i-- {
		if bitString[i] == '1' {
			result += factor
		}

		factor *= 2
	}

	return result
}
