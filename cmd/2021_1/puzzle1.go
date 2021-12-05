package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("/Users/hamishmacdonald/Learning/go/adventofcode/assets/input_day1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var curr, prev, ctr, sumCurr3, sumPrev3, ctr3 int
	var curr3, prev3 []string
	for i,s := range lines {
		prev = curr
		curr, err = strconv.Atoi(s)
		if i > 0 && curr > prev {
			ctr++
		}
		if i > 2 {
			prev3 = lines[i-3:i]
			curr3 = lines[i-2:i+1]
			sumPrev3 = sumSlice(prev3)
			sumCurr3 = sumSlice(curr3)
			if sumCurr3 > sumPrev3 {
				ctr3++
			}
		}
	}
	fmt.Printf("Number of measurements greater than their predecessor is %v\n", ctr)
	fmt.Printf("Number of sliding 3 block windows greater than their predecessor is %v", ctr3)
}

func sumSlice(subslice []string) int {
	var v, sum int
	for _,s := range subslice {
		v,_ = strconv.Atoi(s)
		sum += v
	}
	return sum
}