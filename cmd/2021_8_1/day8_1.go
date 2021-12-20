package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func main () {
	opSys := runtime.GOOS
	var filepath string
	if opSys == "darwin" {
		filepath = "/Users/hamishmacdonald/Learning/go/adventofcode/assets/input_day8"
	} else {
		filepath = "/home/elusivenode/study/go_projects/adventofcode/assets/input_day8"
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count1478 int
	var outputValues []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		io := strings.Split(scanner.Text(), "|")
		output := strings.Split(io[1], " ")
		for _, s := range output {
			if s != "" {
				outputValues = append(outputValues, s)
			}
			if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
				count1478++
			}
		}
	}
	fmt.Printf("The digits of 1,4,7 or 8 appear %v times\n", count1478)
}