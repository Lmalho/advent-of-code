package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func processGridPartOne(grid []string) int {

	var times []int = utils.ConvertStringSliceToIntSlice(strings.Split(strings.TrimSpace(strings.Split(grid[0], ":")[1]), " "))
	var distances []int = utils.ConvertStringSliceToIntSlice(strings.Split(strings.TrimSpace(strings.Split(grid[1], ":")[1]), " "))

	results := make([]int, len(times))

	for i, v := range times {
		winCount := 0

		for j := 0; j < v; j++ {
			d := (v - j) * j

			if d > distances[i] {
				winCount++
			}
		}
		results[i] = winCount

	}
	return utils.MultiplySliceValues(results)

}

func processGridPartTwo(grid []string) int {

	time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(grid[0], ":")[1], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(grid[1], ":")[1], " ", ""))

	winCount := 0

	for j := 0; j < time; j++ {
		d := (time - j) * j
		if d > distance {
			winCount++
		}
	}

	return winCount

}

func main() {
	// read input.txt file

	// for each line, process the line
	// print the result

	p := filepath.Join("input", "input.txt")

	file, err := os.Open(p)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	resultPartOne := processGridPartOne(grid)

	resultPartTwo := processGridPartTwo(grid)

	fmt.Println("resultPartOne", resultPartOne)

	fmt.Println("resultPartTwo", resultPartTwo)

}
