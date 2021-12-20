package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {

	opSys := runtime.GOOS
	var filepath string
	if opSys == "darwin" {
		filepath = "/Users/hamishmacdonald/Learning/go/adventofcode/assets/input_day7"
	} else {
		filepath = "/home/elusivenode/study/go_projects/adventofcode/assets/input_day7"
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var crabPositions []int
	maxPos := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		positions := strings.Split(scanner.Text(), ",")
		for _, pos := range positions {
			p, _ := strconv.Atoi(pos)
			crabPositions = append(crabPositions, p)
			if p > maxPos {
				maxPos = p
			}
		}
	}

	fuelCostMap := make(map[int]int)
	minCost := 1000000
	minPos := 0
	for i := 0; i <= maxPos; i++ {
		calculateFuel(i, &crabPositions, &fuelCostMap)
		if fuelCostMap[i] < minCost {
			minCost = fuelCostMap[i]
			minPos = i
		}
	}
	fmt.Printf("Aligning at position %v has the lowest cost of %v", minPos, minCost)
}

func calculateFuel(alignmentPoint int, startingPositions *[]int, fuelCostMap *map[int]int) {
	fuelCost := 0
	for _, c := range *startingPositions {
		cost := int(math.Abs(float64(alignmentPoint - c)))
		fuelCost += cost
	}
	fcm := *fuelCostMap
	fcm[alignmentPoint] = fuelCost
}