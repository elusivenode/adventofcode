package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("/home/elusivenode/study/go_projects/adventofcode/assets/input_day2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var hori, depth int
	for _, l := range lines {
		direction, length := getDirections(l)
		if direction == "up" {
			depth -= length
		} else if direction == "down" {
			depth += length
		} else if direction == "back" {
			hori -= length
		} else if direction == "forward" {
			hori += length
		} else {
			fmt.Println("Unknown direction.")
		}
	}
	fmt.Printf("Depth is %v and horizonal is %v\n", depth, hori)
	fmt.Printf("Answer is %v\n", depth*hori)
}

func getDirections(dirs string) (string, int) {
	dirsParse := strings.Fields(dirs)
	direction := dirsParse[0]
	length, _ := strconv.Atoi(dirsParse[1])
	return direction, length
}
