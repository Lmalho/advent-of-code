package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getDigitSum(line string) int {
	spelledNumbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	// get the first and last digit and add them
	if len(line) == 0 {
		return 0
	}

	var digitSlice []string

	for i, r := range line {
		if unicode.IsNumber(r) {
			digitSlice = append(digitSlice, string(r))
		} else {
			for j, v := range spelledNumbers {
				// check if the substring from i to end of line starts with the spelled number
				if strings.HasPrefix(line[i:], string(v)) {
					digitSlice = append(digitSlice, strconv.Itoa(j+1))
				}
			}
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

	fmt.Println("sumSlice ", utils.SumSliceValues(sumSlice))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
