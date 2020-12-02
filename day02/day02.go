package main

import (
	"fmt"
	"inputloader"
	"strconv"
	"strings"
)

func validpwPart1(inpt string) bool {
	parsed := strings.Fields(inpt)
	rangeval := strings.Split(parsed[0], "-")
	lowVal, _ := strconv.Atoi(rangeval[0])
	highVal, _ := strconv.Atoi(rangeval[1])
	numberOcs := strings.Count(parsed[2], string(parsed[1][0]))
	returnval := (numberOcs >= lowVal) && (numberOcs <= highVal)
	return returnval
}

func validpwPart2(inpt string) bool {
	parsed := strings.Fields(inpt)
	rangeval := strings.Split(parsed[0], "-")
	lowVal, _ := strconv.Atoi(rangeval[0])
	highVal, _ := strconv.Atoi(rangeval[1])
	// if lowVal+1 >= len(parsed[2]) {
	// 	return false
	// }
	// if highVal+1 >= len(parsed[2]) {
	// 	fmt.Println("debug!")
	// 	return parsed[2][lowVal+1] == parsed[1][0]
	// }
	lowValLow := parsed[2][lowVal-1] == parsed[1][0]
	highValLow := parsed[2][highVal-1] == parsed[1][0]
	returnval := (lowValLow || highValLow) && !(lowValLow && highValLow)
	return returnval
}

func main() {
	inputlines := inputloader.ReadInput("../data/input02.txt")
	sum := 0
	for _, s := range inputlines {
		if validpwPart2(s) {
			sum++
		}
	}
	fmt.Println(sum)
}
