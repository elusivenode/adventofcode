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

	var school []*fish
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		initialSchool := scanner.Text()
		for _,f := range strings.Split(initialSchool, ",") {
			i,_ := strconv.Atoi(f)
			fish := fish {
				daysUntilSpawn: i,
				isNew: false,
			}
			school = append(school, &fish)
		}
	}

	noOfDays := 80
	for days := noOfDays; days > 0; days-- {
		school = elapseOneDay(school)
		fmt.Printf("After %v days, there are %v fish in the school.\n", noOfDays - days + 1, len(school))
	}
}

func elapseOneDay(school []*fish) []*fish {
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
	return school
}
