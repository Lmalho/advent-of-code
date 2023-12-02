package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func analyzeResults(gameResults []string) bool {

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

	if results["red"] > 12 || results["blue"] > 14 || results["green"] > 13 {
		return false
	}

	return true
}

func processLine(line string) (int, bool) {

	//Get Game id
	gameIdIndex := strings.Index(line, ":")
	gameWordIndex := strings.Index(line, "Game")
	gameId, _ := strconv.Atoi(string(line[gameWordIndex+5 : gameIdIndex]))

	fmt.Println(gameId)
	games := strings.Split(line[gameIdIndex+1:], ";")

	validGame := true

	for _, game := range games {

		validGame = analyzeResults(strings.Split(game, ","))
		if !validGame {
			break
		}

	}

	return gameId, validGame
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

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		gameId, ok := processLine(scanner.Text())
		validMap[gameId] = ok
	}

	fmt.Println(sumValidIds(validMap))
}
