package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func processLine(counter *int, line string, cardMap *map[int]int) int {
	*counter++
	//Get Game id

	gameIdIndex := strings.Index(line, ":")
	numbers := strings.Split(line[gameIdIndex+1:], "|")

	winningNumbers := utils.ConvertStringSliceToIntSlice(strings.Split(strings.TrimSpace(numbers[0]), " "))
	cardNumbers := utils.ConvertStringSliceToIntSlice(strings.Split(strings.TrimSpace(numbers[1]), " "))

	slices.Sort(winningNumbers)
	slices.Sort(cardNumbers)

	countWinningNumbers := 0

	(*cardMap)[*counter] += 1

	for _, number := range cardNumbers {
		_, found := slices.BinarySearch(winningNumbers, number)
		if found {
			countWinningNumbers++
		}
	}
	for i := 1; i <= countWinningNumbers; i++ {
		(*cardMap)[*counter+i] += (*cardMap)[*counter]
	}

	return int(math.Pow(2, float64(countWinningNumbers-1)))
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

	var results []int
	cardMap := make(map[int]int)
	scanner := bufio.NewScanner(file)
	points := 0
	counter := 0

	for scanner.Scan() {
		points = processLine(&counter, scanner.Text(), &cardMap)

		results = append(results, points)
	}

	fmt.Println("Result part one: ", utils.SumSliceValues(results))
	fmt.Println("Result part two: ", utils.SumMapValues(cardMap))

}
