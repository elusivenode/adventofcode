package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("/Users/hamishmacdonald/Learning/go/adventofcode/assets/input_day1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var curr, prev, ctr int
	for i,s := range lines {
		prev = curr
		curr, err = strconv.Atoi(s)
		if i > 0 && curr > prev {
			ctr++
		}
	}
	fmt.Printf("Number of measurements greater than their predecessor is %v", ctr)
}
