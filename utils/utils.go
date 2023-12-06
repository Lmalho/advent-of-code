package utils

import "strconv"

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

func MultiplySliceValues(slice []int) int {
	var mult int = 1

	for _, v := range slice {
		mult *= v
	}

	return mult
}

func ConvertStringSliceToIntSlice(stringSlice []string) []int {
	var intSlice []int

	for _, v := range stringSlice {
		if v == "" {
			continue
		}
		intElement, _ := strconv.Atoi(v)
		intSlice = append(intSlice, intElement)
	}

	return intSlice
}

func SumMapValues(validMap map[int]int) int {
	var sum int

	for _, v := range validMap {
		sum += v
	}
	return sum
}
