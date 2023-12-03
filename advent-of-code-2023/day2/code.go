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

func analyzeResults(gameResults []string) map[string]int {

	results := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	for _, gameResult := range gameResults {
		color := strings.Split(strings.TrimSpace(gameResult), " ")

		value, _ := strconv.Atoi(strings.TrimSpace(color[0]))

		results[color[1]] = value
	}

	return results
}

func processLine(line string) (int, bool, int) {

	//Get Game id
	gameIdIndex := strings.Index(line, ":")
	gameWordIndex := strings.Index(line, "Game")
	gameId, _ := strconv.Atoi(string(line[gameWordIndex+5 : gameIdIndex]))

	fmt.Println(gameId)
	games := strings.Split(line[gameIdIndex+1:], ";")

	groupResultMax := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	validGame := true

	for _, game := range games {

		gameResult := analyzeResults(strings.Split(game, ","))

		if gameResult["red"] > 12 || gameResult["blue"] > 14 || gameResult["green"] > 13 {
			validGame = false
		}

		for k, v := range gameResult {
			if v > groupResultMax[k] {
				groupResultMax[k] = v
			}
		}

	}

	powerValue := utils.MultiplyMapValues(groupResultMax)

	return gameId, validGame, powerValue
}

func sumValidIds(validMap map[int]bool) int {
	var sum int

	for k, v := range validMap {
		if v {
			sum += k
		}
	}

	return sum
}

func sumPowerValues(powerMap map[int]int) int {
	var sum int

	for _, v := range powerMap {
		sum += v
	}

	return sum
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

	validMap := make(map[int]bool)
	powerMap := make(map[int]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		gameId, ok, powerValue := processLine(scanner.Text())
		validMap[gameId] = ok
		powerMap[gameId] = powerValue
	}

	fmt.Println("Part one: ", sumValidIds(validMap))
	fmt.Println("Part two: ", sumPowerValues(powerMap))
}
