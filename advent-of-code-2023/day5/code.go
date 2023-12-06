package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getAlmanacMap(startIndex int, grid []string) [][]int {

	var almanacMap [][]int

	for i := startIndex; i < len(grid); i++ {
		if grid[i] == "" {
			break
		}

		values := utils.ConvertStringSliceToIntSlice(strings.Split(grid[i], " "))
		almanacMap = append(almanacMap, values)

	}
	return almanacMap
}

func crossNumbers(originMap [][]int, origin int, crossNumbers *[]int) {

	added := false
	for _, v := range originMap {

		if v[1] <= origin && origin < v[1]+v[2] {

			*crossNumbers = append(*crossNumbers, v[0]+origin-v[1])
			added = true
			break
		}

	}
	if !added {
		*crossNumbers = append(*crossNumbers, origin)
	}
}

func almanacMaps(grid []string, seedToSoil, soilToFertilezer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation *[][]int) {

	for index, row := range grid {
		switch row {
		case "seed-to-soil map:":
			*seedToSoil = getAlmanacMap(index+1, grid)

		case "soil-to-fertilizer map:":
			*soilToFertilezer = getAlmanacMap(index+1, grid)

		case "fertilizer-to-water map:":

			*fertilizerToWater = getAlmanacMap(index+1, grid)

		case "water-to-light map:":

			*waterToLight = getAlmanacMap(index+1, grid)

		case "light-to-temperature map:":

			*lightToTemperature = getAlmanacMap(index+1, grid)

		case "temperature-to-humidity map:":

			*temperatureToHumidity = getAlmanacMap(index+1, grid)

		case "humidity-to-location map:":

			*humidityToLocation = getAlmanacMap(index+1, grid)

		}
	}

}

func processGridPartOne(grid []string) ([][]int, int) {

	// get the seeds (numbers) from the first row

	seeds := utils.ConvertStringSliceToIntSlice(strings.Split(grid[0][strings.Index(grid[0], ":")+1:], " "))

	var seedToSoil, soilToFertilezer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation [][]int

	almanacMaps(grid, &seedToSoil, &soilToFertilezer, &fertilizerToWater, &waterToLight, &lightToTemperature, &temperatureToHumidity, &humidityToLocation)

	var finalMap [][]int
	lowestLocation := 0

	for index, seed := range seeds {
		seedMap := []int{seed}

		crossNumbers(seedToSoil, seed, &seedMap)
		crossNumbers(soilToFertilezer, seedMap[1], &seedMap)
		crossNumbers(fertilizerToWater, seedMap[2], &seedMap)
		crossNumbers(waterToLight, seedMap[3], &seedMap)
		crossNumbers(lightToTemperature, seedMap[4], &seedMap)
		crossNumbers(temperatureToHumidity, seedMap[5], &seedMap)
		crossNumbers(humidityToLocation, seedMap[6], &seedMap)
		finalMap = append(finalMap, seedMap)

		if index == 0 {
			lowestLocation = seedMap[7]
		} else if lowestLocation > seedMap[7] {
			lowestLocation = seedMap[7]
		}

	}

	return finalMap, lowestLocation

}

func processGridPartTwo(grid []string) ([][]int, int) {

	// get the seeds (numbers) from the first row

	seeds := utils.ConvertStringSliceToIntSlice(strings.Split(grid[0][strings.Index(grid[0], ":")+1:], " "))

	// divide seeds in array of 2 elements start, range

	seedranges := [][]int{}

	for i := 0; i < len(seeds); i += 2 {
		seedranges = append(seedranges, []int{seeds[i], seeds[i+1]})
	}

	var seedToSoil, soilToFertilezer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation [][]int

	almanacMaps(grid, &seedToSoil, &soilToFertilezer, &fertilizerToWater, &waterToLight, &lightToTemperature, &temperatureToHumidity, &humidityToLocation)

	for _, seedrange := range seedranges {

		seedMapMin := []int{seedrange[0]}
		seedMapMax := []int{seedrange[0] + seedrange[1] - 1}

		crossNumbers(seedToSoil, seedrange[0], &seedMapMin)
		crossNumbers(seedToSoil, seedrange[0]+seedrange[1]-1, &seedMapMax)

		crossNumbers(soilToFertilezer, seedMapMin[1], &seedMapMin)
		crossNumbers(soilToFertilezer, seedMapMax[1], &seedMapMax)

		crossNumbers(fertilizerToWater, seedMapMin[2], &seedMapMin)
		crossNumbers(fertilizerToWater, seedMapMax[2], &seedMapMax)

		crossNumbers(waterToLight, seedMapMin[3], &seedMapMin)
		crossNumbers(waterToLight, seedMapMax[3], &seedMapMax)

		crossNumbers(lightToTemperature, seedMapMin[4], &seedMapMin)
		crossNumbers(lightToTemperature, seedMapMax[4], &seedMapMax)

		crossNumbers(temperatureToHumidity, seedMapMin[5], &seedMapMin)
		crossNumbers(temperatureToHumidity, seedMapMax[5], &seedMapMax)

		crossNumbers(humidityToLocation, seedMapMin[6], &seedMapMin)
		crossNumbers(humidityToLocation, seedMapMax[6], &seedMapMax)

		fmt.Println("seedMapMin", seedMapMin)
		fmt.Println("seedMapMax", seedMapMax)

	}

	return nil, 0
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

	finalMapPartOne, lowestLocationPartOne := processGridPartOne(grid)

	finalMapPartTwo, lowestLocationPartTwo := processGridPartTwo(grid)

	fmt.Println("Part One")
	fmt.Println("	finalMap", finalMapPartOne)
	fmt.Println("	lowestLocation", lowestLocationPartOne)
	fmt.Println("Part Two")
	fmt.Println("	finalMap", finalMapPartTwo)
	fmt.Println("	lowestLocation", lowestLocationPartTwo)

}
