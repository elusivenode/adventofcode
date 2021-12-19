package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type fish struct {
	daysUntilSpawn int
	isNew bool
}

func main() {

	opSys := runtime.GOOS
	var filepath string
	if opSys == "darwin" {
		filepath = "/Users/hamishmacdonald/Learning/go/adventofcode/assets/input_day6"
	} else {
		filepath = "/home/elusivenode/study/go_projects/adventofcode/assets/input_day6"
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var parentFish []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		initialSchool := scanner.Text()
		for _, f := range strings.Split(initialSchool, ",") {
			i,_ := strconv.Atoi(f)
			parentFish = append(parentFish, i)
		}
	}

	var wg sync.WaitGroup
	var schoolChannel chan *[]*fish

	schoolChannel = make(chan *[]*fish, len(parentFish))
	daysToTrack := 256
	for _, f := range parentFish {
		wg.Add(1)
		go processFish(f, daysToTrack, &wg, schoolChannel)
	}
	wg.Wait()
	close(schoolChannel)

	totalFish := 0
	for school := range schoolChannel {
		totalFish += len(*school)
	}

	fmt.Printf("Total count: %v fish", totalFish)
}

func processFish(dayUntilSpawn int, days int, wg *sync.WaitGroup, ch chan *[]*fish) {
	defer wg.Done()

	parentFish := fish{
		daysUntilSpawn: dayUntilSpawn,
		isNew: false,
	}
	var school []*fish
	school = append(school , &parentFish)

	for d := days; d > 0 ; d-- {
		for _, f := range school {
			if f.isNew == true && f.daysUntilSpawn == 8 {
				f.isNew = false
			}
			if f.isNew == false {
				if f.daysUntilSpawn == 0 {
					f.daysUntilSpawn = 6
					newFish := fish {
						daysUntilSpawn: 8,
						isNew: true,
					}
					school = append(school, &newFish)
				} else {
					f.daysUntilSpawn--
				}
			}
		}
		fmt.Printf("Day: %v processed\n", days - d + 1)
	}
	ch <- &school
}