package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

func main() {

	opSys := runtime.GOOS
	var filepath string
	if opSys == "darwin" {
		filepath = "/Users/hamishmacdonald/Learning/go/adventofcode/assets/input_day3"
	} else {
		filepath = "/home/elusivenode/study/go_projects/adventofcode/assets/input_day3"
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	oxygen := make([]string, len(lines))
	copy(oxygen, lines)
	co2 := make([]string, len(lines))
	copy(co2, lines)

	var foundOxygen, foundCO2 bool
	foundOxygen = false
	foundCO2 = false

	i := 0
	for !foundOxygen {
		mostCommon, _ := getCounts(oxygen, i)
		oxygen = findElement(oxygen, mostCommon, i)
		if len(oxygen) == 1 {
			foundOxygen = true
		}
		i++
	}

	i = 0
	for !foundCO2 {
		_, leastCommon := getCounts(co2, i)
		co2 = findElement(co2, leastCommon, i)
		if len(co2) == 1 {
			foundCO2 = true
		}
		i++
	}
	oxygenRating, _ := strconv.ParseInt(oxygen[0], 2, 64)
	co2ScrubberRating, _ := strconv.ParseInt(co2[0], 2, 64)
	fmt.Printf("Oxygen rating is: %v; CO2 scrubber rating: %v; Life support rating is: %v", oxygenRating, co2ScrubberRating, oxygenRating*co2ScrubberRating)
}

func getCounts(slice []string, idx int) (mostCommon uint8, leastCommon uint8) {
	zeros := 0
	ones := 0
	for _, x := range slice {
		if x[idx] == '0' {
			zeros++
		} else {
			ones++
		}
	}
	if zeros > ones {
		mostCommon = '0'
		leastCommon = '1'
	} else {
		leastCommon = '0'
		mostCommon = '1'
	}
	return mostCommon, leastCommon
}

func findElement(slice []string, criteria uint8, idx int) []string {
	i := 0
	for _, x := range slice {
		if x[idx] == criteria {
			slice[i] = x
			i++
		}
	}
	return slice[:i]
}
