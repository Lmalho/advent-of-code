package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getDigitSum(line string) int {
	// get the first and last digit and add them
	if len(line) == 0 {
		return 0
	}

	var digitSlice []string

	for _, r := range line {
		if unicode.IsNumber(r) {
			digitSlice = append(digitSlice, string(r))
		}
	}

	if len(digitSlice) == 0 {
		return 0
	}

	firstDigit, _ := strconv.Atoi(digitSlice[0])
	lastDigit, _ := strconv.Atoi(digitSlice[len(digitSlice)-1])
	result, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
	return result
}

func sumSliceValues(slice []int) int {
	var sum int

	for _, v := range slice {
		sum += v
	}

	return sum
}

func main() {
	// read input.txt file
	// for each line, process the line
	// print the result

	file, err := os.Open("input/input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sumSlice []int

	for scanner.Scan() {
		fmt.Println("sum", getDigitSum(scanner.Text()))
		sumSlice = append(sumSlice, getDigitSum(scanner.Text()))
	}

	fmt.Println("sumSlice ", sumSliceValues(sumSlice))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
