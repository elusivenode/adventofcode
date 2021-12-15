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

type winningGame struct {
	gameNo int
	roundAchievedBingo int
	finalScore int
	positionNo int
}

var position int = 1
var round int = 1
var games = map[int][][]int{}
var scores = map[int][][]int{}
var progress = map[int]map[string]int{}
var winners = make(map[int]*winningGame)
var firstWinner *winningGame
var lastWinner *winningGame
var setWinningGames = make(map[int]bool)

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
		processNumber(keys, no)

		fmt.Printf("Processing complete on no. %v.\n\n", n)
		round += 1
		gameToPrint := 4
		wantToPrint := false
		if wantToPrint {
			printProgress(gameToPrint, games[gameToPrint], scores[gameToPrint], no)
		}
	}
	for k, v := range winners {
		fmt.Printf("Game: %v Stats: %v\n", k, v)
	}
	fmt.Printf("First winner: %v\n", firstWinner)
	fmt.Printf("Last winner: %v\n", lastWinner)
}

func printProgress(gameNo int, game[][]int, score[][]int, No int) {
	fmt.Printf("Score evolution for game %v:\n", gameNo)
	fmt.Printf("Number drawn: %v\n", No)

	ct := 5
	for i := 0; i < ct; i++ {
		fmt.Println("")
		for j := 0; j < ct; j++ {
			fmt.Printf("%v-%v\t", game[i][j], score[i][j])
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

func processNumber(keys []int, number int) {
	for _, k := range keys {
		alreadyWinner := setWinningGames[k]
		if !alreadyWinner {
			processScores(number, k)
		}
	}
}

func processScores(number int, game int) {
	ct := len(games[game])
	for i := 0; i < ct; i++ {
		for j := 0; j < ct; j++ {
			if games[game][i][j] == int(number) {
				//fmt.Printf("No. %v found in game %v\n", number, game)
				scores[game][i][j] = 1
				processProgress(i, j, game, number)
			}
		}
	}
}

func processProgress(rowNo int, colNo int, game int, number int) {
	progress[game]["r" + strconv.Itoa(rowNo)]  += 1
	progress[game]["c" + strconv.Itoa(colNo)]  += 1
	if progress[game]["r" + strconv.Itoa(rowNo)] == 5 || progress[game]["c" + strconv.Itoa(colNo)] == 5 {
		updateWinners(game, round, number)
	}
}

func updateWinners(game int, round int, number int) {
	winner := &winningGame {
		gameNo: game,
		roundAchievedBingo: round,
		finalScore: calcWinningScore(games[game], scores[game], number),
		positionNo: position,
	}
	winners[game] = winner
	if position == 1 {
		firstWinner = winner
	}
	setWinningGames[game] = true
	lastWinner = winner
	position += 1
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

