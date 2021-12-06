package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("/Users/hamishmacdonald/Learning/go/adventofcode/assets/input_day3")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var counts = map[string]map[string]int{}
	counts["pos0"] = map[string]int{}
	counts["pos1"] = map[string]int{}
	counts["pos2"] = map[string]int{}
	counts["pos3"] = map[string]int{}
	counts["pos4"] = map[string]int{}
	counts["pos5"] = map[string]int{}
	counts["pos6"] = map[string]int{}
	counts["pos7"] = map[string]int{}
	counts["pos8"] = map[string]int{}
	counts["pos9"] = map[string]int{}
	counts["pos10"] = map[string]int{}
	counts["pos11"] = map[string]int{}

	setUpMaps(counts)

	for _, l := range lines {
		processBitSequence(l, counts)
	}

	gamma, epsilon := getGammaEpsilon(counts)
	fmt.Printf("Gamma: %v - Epsilon %v - Power %v\n", gamma, epsilon, gamma * epsilon)
}

func getGammaEpsilon(masterMap map[string]map[string]int) (int64, int64) {
	var gamma, epsilon string
	for i := 0; i < len(masterMap); i++ {
		v0 := masterMap["pos" + strconv.Itoa(i)]["0"]
		v1 := masterMap["pos" + strconv.Itoa(i)]["1"]

		if v0 > v1 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	return gammaInt, epsilonInt
}

func processBitSequence(bitSeq string, masterMap map[string]map[string]int) {
	for i, c := range bitSeq {
		masterMap["pos" + strconv.Itoa(i)][string(c)] = masterMap["pos" + strconv.Itoa(i)][string(c)] + 1
	}
}

func setUpMaps(masterMap map[string]map[string]int) {
	for _,v := range masterMap {
		v["0"] = 0
		v["1"] = 0
	}
}