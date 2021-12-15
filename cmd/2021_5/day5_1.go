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

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
	isHor bool
	isVert bool
}

func main() {

	opSys := runtime.GOOS
	var filepath string
	if opSys == "darwin" {
		filepath = "/Users/hamishmacdonald/Learning/go/adventofcode/assets/input_day5"
	} else {
		filepath = "/home/elusivenode/study/go_projects/adventofcode/assets/input_day5"
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []*line
	scanner := bufio.NewScanner(file)
	var maxX, maxY int

	for scanner.Scan() {
		lineOfText := scanner.Text()
		points := strings.Split(strings.ReplaceAll(lineOfText, " ",""), "->")
		var x1, x2, y1, y2 int
		var isHor, isVert bool
		for i, p := range points {
			splitPoints := strings.Split(p, ",")
			if i == 0 {
				x1, _ = strconv.Atoi(splitPoints[0])
				y1, _ = strconv.Atoi(splitPoints[1])
				if x1 > maxX {
					maxX = x1
				}
				if y1 > maxY {
					maxY = y1
				}
			} else {
				x2, _ = strconv.Atoi(splitPoints[0])
				y2, _ = strconv.Atoi(splitPoints[1])

				if x2 > maxX {
					maxX = x2
				}
				if y2 > maxY {
					maxY = y2
				}

				if y1 == y2 {
					isHor = true
				}
				if x1 == x2 {
					isVert = true
				}
				newLine := newLine(x1, y1, x2, y2, isHor, isVert)
				lines = append(lines, newLine)
			}
		}
	}

	mapOfPoints := getMapOfPoints(maxX, maxY)
	for _, l := range lines {
		if l.isHor || l.isVert {
			processLine(mapOfPoints, l)
		}
	}

	ctOverlappingLines := 0
	for _,v := range mapOfPoints {
		if v >= 2 {
			ctOverlappingLines++
		}
	}
	fmt.Printf("There are %v points at which at least 2 lines overlap", ctOverlappingLines)
}

func processLine(pointsMap map[string]int, line *line) {
	if line.isHor {
		y := line.y1
		var xMin, xMax int
		if line.x1 < line.x2 {
			xMin = line.x1
			xMax = line.x2
		} else {
			xMin = line.x2
			xMax = line.x1
		}
		for i := xMin; i <= xMax; i++ {
			point := strconv.Itoa(i) + "-" + strconv.Itoa(y)
			pointsMap[point]++
		}
	} else {
		x := line.x1
		var yMin, yMax int
		if line.y1 < line.y2 {
			yMin = line.y1
			yMax = line.y2
		} else {
			yMin = line.y2
			yMax = line.y1
		}
		for i := yMin; i <= yMax; i++ {
			point := strconv.Itoa(x) + "-" + strconv.Itoa(i)
			pointsMap[point]++
		}
	}
}

func getMapOfPoints(maxX int, maxY int) map[string]int {
	points := make(map[string]int)
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			pt := strconv.Itoa(i) + "-" + strconv.Itoa(j)
			points[pt] = 0
		}
	}
	return points
}

func newLine(x1 int, y1 int, x2 int, y2 int, isHor bool, isVert bool) *line {
	newLine := line {
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
		isHor: isHor,
		isVert: isVert,
	}
	return &newLine
}