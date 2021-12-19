package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {

	opSys := runtime.GOOS
	var filepath string
	if opSys == "darwin" {
		filepath = "/Users/hamishmacdonald/Learning/go/adventofcode/assets/input_day6"
	} else {
		filepath = "/home/elusivenode/study/go_projects/adventofcode/assets/input_day6"
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	population := make(map[int]int64)
	for i := 0; i <= 8; i++ {
		population[i] = 0
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		initialSchool := scanner.Text()
		for _, f := range strings.Split(initialSchool, ",") {
			i,_ := strconv.Atoi(f)
			population[i]++
		}
	}

	days := 256
	for i := 1; i <= days; i++ {
		currentPopulation := make(map[int]int64)
		for k,v := range population {
			currentPopulation[k] = v
		}
		for k := 8; k >= 1; k-- {
			population[k-1] = currentPopulation[k]
		}
		population[8] = currentPopulation[0]
		population[6] += currentPopulation[0]
	}

	totalFish := int64(0)
	for _,v := range population {
		totalFish += v
	}
	fmt.Printf("Total fish after %v days is %v.\n", days, totalFish)
}