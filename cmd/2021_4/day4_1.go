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
	var scores = map[int][][]int{}
	var progress = map[int]map[string]int{}
	game := make([][]int, 5)
	score := make([][]int, 5)
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
			score[lineCtr] = make([]int, 5)
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
				scores[gameNo] = score
			}
		}
	}

	keys := make([]int, len(games))
	i := 0
	for k := range games {
		keys[i] = k
		i++
	}

	setUpProgressMap(progress, keys)

	for _, n := range numbers {
		processNumber(games, scores, progress, keys, n)
	}

	fmt.Println(len(numbers))
}

func processNumber(gamesMap map[int][][]int, scoresMap map[int][][]int, progressMap map[int]map[string]int,
	keys []int, number int32) {
	for _, k := range keys {
		processScores(gamesMap, scoresMap, progressMap, number, k)
	}
	fmt.Printf("%v %v %v %v %v", gamesMap, scoresMap, progressMap, keys, number)
}

func processScores(gamesMap map[int][][]int, scoresMap map[int][][]int, progressMap map[int]map[string]int,
	number int32, game int) {
	ct := len(gamesMap[game])
	for i := 0; i < ct; i++ {
		for j := 0; j < ct; j++ {
			if gamesMap[game][i][j] == int(number) {
				scoresMap[game][i][j] = 1
				processProgress(progressMap, i, j, game)
			}
		}
	}
	fmt.Printf("%v %v %v %v", gamesMap, scoresMap, number, game)
}

func processProgress(progressMap map[int]map[string]int, rowNo int, colNo int, game int) {
	progressMap[game]["r" + strconv.Itoa(rowNo)]  += 1
	progressMap[game]["c" + strconv.Itoa(colNo)]  += 1
}

func setUpProgressMap(progressMap map[int]map[string]int, keys []int) {
		labels := []string{"r", "c"}
		var progressKey string
		for _, k := range keys {
			progressMap[k] = make(map[string]int)
			for _, l := range labels {
				for i := 1; i <=5; i++ {
					ctr := strconv.Itoa(i)
					progressKey = l + ctr
					progressMap[k][progressKey] = 0
				}
			}
		}
	}

