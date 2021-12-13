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
	var numbers []string
	var games = map[int][][]int{}
	var scores = map[int][][]int{}
	var progress = map[int]map[string]int{}
	var gameNo, lineCtr int
	var prevLineBlank bool = false

	for scanner.Scan() {
		line := scanner.Text()
		//first line is long
		if len(line) > 20 {
			numbers = strings.Split(line, ",")
		} else if len(line) == 0 {
			prevLineBlank = true
		} else {
			lineNos := strings.Split(line, " ")
			if prevLineBlank == true {
				lineCtr = 0
				gameNo++
				prevLineBlank = false
				games[gameNo] = make([][]int, 5)
				scores[gameNo] = make([][]int, 5)
			}
			games[gameNo][lineCtr] = make([]int, 5)
			scores[gameNo][lineCtr] = make([]int, 5)
			i := 0
			for _, n := range lineNos {
				if n != "" {
					games[gameNo][lineCtr][i], _ = strconv.Atoi(n)
					i++
				}
			}
			lineCtr++
		}
	}

	keys := make([]int, len(games))
	i := 0
	for k, _ := range games {
		keys[i] = k
		i++
	}

	setUpProgressMap(progress, keys)

	for _, n := range numbers {
		no, _ := strconv.Atoi(n)
		winnerFound, winningGame := processNumber(games, scores, progress, keys, no)
		if winnerFound {
			fmt.Printf("Game %v bingo!\n", winningGame)
			winningScore := calcWinningScore(games[winningGame], scores[winningGame], no)
			fmt.Printf("The winning score was %v\n", winningScore)
			break
		} else {
			//fmt.Printf("Process no. %v.  No winner found yet.\n", n)
		}
	}
}

func calcWinningScore(game [][]int, score [][]int, lastNo int) int {
	ct := len(game)
	sumUnmarked := 0
	for i := 0; i < ct; i++ {
		for j := 0; j < ct; j++ {
			if score[i][j] == 0 {
				sumUnmarked += game[i][j]
			}
		}
	}
	return sumUnmarked * lastNo
}

func processNumber(gamesMap map[int][][]int, scoresMap map[int][][]int, progressMap map[int]map[string]int,
	keys []int, number int)  (winnerFound bool, winningGame int){
	for _, k := range keys {
		winnerFound, winningGame := processScores(gamesMap, scoresMap, progressMap, number, k)
		if winnerFound {
			return true, winningGame
		}
	}
	return false, -1
}

func processScores(gamesMap map[int][][]int, scoresMap map[int][][]int, progressMap map[int]map[string]int,
	number int, game int) (winnerFound bool, winningGame int){
	ct := len(gamesMap[game])
	for i := 0; i < ct; i++ {
		for j := 0; j < ct; j++ {
			if gamesMap[game][i][j] == int(number) {
				//fmt.Printf("No. %v found in game %v\n", number, game)
				scoresMap[game][i][j] = 1
				winnerFound, winningGame := processProgress(progressMap, i, j, game)
				if winnerFound {
					return true, winningGame
				}
			}
		}
	}
	return false, -1
}

func processProgress(progressMap map[int]map[string]int, rowNo int, colNo int, game int) (winnerFound bool, winningGame int){
	progressMap[game]["r" + strconv.Itoa(rowNo)]  += 1
	if progressMap[game]["r" + strconv.Itoa(rowNo)] == 5 {
		return true, game
	}
	progressMap[game]["c" + strconv.Itoa(colNo)]  += 1
	if progressMap[game]["c" + strconv.Itoa(colNo)] == 5 {
		return true, game
	}
	return false, -1
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

