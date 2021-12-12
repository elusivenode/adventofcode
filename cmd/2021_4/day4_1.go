package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/home/elusivenode/study/go_projects/adventofcode/assets/input_day4")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers string
	var game []string
	var games = map[int][]string{}
	var i, lineCtr int
	var prevLineBlank bool = false

	for scanner.Scan() {
		line := scanner.Text()
		//first line is long
		if len(line) > 20 {
			numbers = line
		} else if len(line) == 0 {
			prevLineBlank = true
		} else {
			if prevLineBlank == true {
				game = nil
				lineCtr = 1
				i++
				prevLineBlank = false
			}
			game = append(game, line)
			lineCtr++
			if lineCtr == 6 {
				games[i] = game
			}
		}
	}
	_, test := numbers, games
	gameTest := test[100]
	for _, l := range gameTest {
		fmt.Println(l)
	}
}