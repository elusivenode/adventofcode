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


	var schools = map[int]*[]*fish{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		initialSchool := scanner.Text()
		for ct, f := range strings.Split(initialSchool, ",") {
			var school []*fish
			i,_ := strconv.Atoi(f)
			fish := fish {
				daysUntilSpawn: i,
				isNew: false,
			}
			school = append(school , &fish)
			schools[ct] = &school
		}
	}

	noOfDays := 80
	totalFish := 0

	for _, s := range schools {
		for days := noOfDays; days > 0; days-- {
			elapseOneDay(s)
		}
		totalFish += len(*s)
	}
	fmt.Printf("After %v days there are %v fish in the school", noOfDays, totalFish)
}
func elapseOneDay(school *[]*fish) {
	for _, f := range *school {
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
				*school = append(*school, &newFish)
			} else {
				f.daysUntilSpawn--
			}
		}
	}
}
