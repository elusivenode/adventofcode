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
		filepath = "/Users/hamishmacdonald/Learning/go/adventofcode/assets/input_day4"
	} else {
		filepath = "/home/elusivenode/study/go_projects/adventofcode/assets/input_day4"
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers string
	var games = map[int][][]int{}
	game := make([][]int, 5)
	var gameNo, lineCtr int
	var prevLineBlank bool = false

	for scanner.Scan() {
		line := scanner.Text()
		//first line is long
		if len(line) > 20 {
			numbers = line
		} else if len(line) == 0 {
			prevLineBlank = true
		} else {
			lineNos := strings.Split(line, " ")
			if prevLineBlank == true {
				lineCtr = 0
				gameNo++
				prevLineBlank = false
				for i := 0; i < 5; i++ {
					game[i] = nil
				}

			}
			game[lineCtr] = make([]int, 5)
			i := 0
			for _, n := range lineNos {
				if n != "" {
					game[lineCtr][i], _ = strconv.Atoi(n)
					i++
				}
			}
			lineCtr++
			if lineCtr == 5 {
				games[gameNo] = game
			}
		}
	}
	_, test := numbers, games
	gameTest := test[100]
	for _, l := range gameTest {
		fmt.Println(l)
	}
}
