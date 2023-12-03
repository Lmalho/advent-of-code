package utils

func SumSliceValues(slice []int) int {
	var sum int

	for _, v := range slice {
		sum += v
	}

	return sum
}

func MultiplyMapValues(validMap map[string]int) int {
	var mult int = 1

	for _, v := range validMap {
		mult *= v
	}
	return mult
}
