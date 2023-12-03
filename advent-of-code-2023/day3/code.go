package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
)

type coordinate struct {
	row, col int
}

func processGridPartOne(grid []string) {
	// find all symbols in a grid
	// ignore numbers and dots
	// record the coordinates of each symbol
	// find adjacent numbers to each symbol in a 3x3 grid

	numberCoordinates := make(map[coordinate]bool)

	for row, line := range grid {
		for col, r := range line {
			if unicode.IsNumber(r) || string(r) == "." {
				continue
			} else {
				// fmt.Println("symbol: ", string(r), "row: ", row, "col: ", col)
				for adjRow := row - 1; adjRow <= row+1; adjRow++ {
					for adjCol := col - 1; adjCol <= col+1; adjCol++ {
						if adjRow < 0 || adjCol < 0 || adjRow > len(grid)-1 || adjCol > len(line)-1 || !unicode.IsNumber(rune(grid[adjRow][adjCol])) {
							continue
						}
						// found a number adjacent to the symbol, run through the column in the row to find the first digit
						n := adjCol
						for n > 0 && unicode.IsNumber(rune(grid[adjRow][n-1])) {
							n--
						}
						numberCoordinates[coordinate{adjRow, n}] = true
					}
				}
			}
		}
	}

	// fmt.Println(numberCoordinates)

	var numberList []int

	// find the full number in the grid with the coordinates

	for k := range numberCoordinates {
		var number string
		for i := k.col; i < len(grid[k.row]); i++ {
			if unicode.IsNumber(rune(grid[k.row][i])) {
				number += string(grid[k.row][i])
			} else {
				break
			}
		}
		// fmt.Println(number)
		n, _ := strconv.Atoi(number)
		numberList = append(numberList, n)
	}

	fmt.Println(numberList)
	fmt.Println("Result part one: ", utils.SumSliceValues(numberList))
}

func processGridPartTwo(grid []string) {
	var gearList []int
	for row, line := range grid {
		for col, r := range line {
			if string(r) != "*" {
				continue
			} else {
				gearCoordinates := make(map[coordinate]bool)
				for adjRow := row - 1; adjRow <= row+1; adjRow++ {
					for adjCol := col - 1; adjCol <= col+1; adjCol++ {
						if adjRow < 0 || adjCol < 0 || adjRow > len(grid)-1 || adjCol > len(line)-1 || !unicode.IsNumber(rune(grid[adjRow][adjCol])) {
							continue
						}
						// found a number adjacent to the symbol, run through the column in the row to find the first digit
						n := adjCol
						for n > 0 && unicode.IsNumber(rune(grid[adjRow][n-1])) {
							n--
						}
						gearCoordinates[coordinate{adjRow, n}] = true
					}
				}
				if len(gearCoordinates) != 2 {
					continue
				} else {
					var numberList []int
					for k := range gearCoordinates {
						var number string
						for i := k.col; i < len(grid[k.row]); i++ {
							if unicode.IsNumber(rune(grid[k.row][i])) {
								number += string(grid[k.row][i])
							} else {
								break
							}
						}
						// fmt.Println(number)
						n, _ := strconv.Atoi(number)
						numberList = append(numberList, n)
					}
					gearList = append(gearList, numberList[0]*numberList[1])
				}
			}
		}
	}

	fmt.Println("Result part two: ", utils.SumSliceValues(gearList))
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

	processGridPartOne(grid)
	processGridPartTwo(grid)

}
