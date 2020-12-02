package main

import (
	"fmt"
	"inputloader"
	"strconv"
	"strings"
)

func validpw(inpt string) bool {
	parsed := strings.Fields(inpt)
	rangeval := strings.Split(parsed[0], "-")
	lowVal, _ := strconv.Atoi(rangeval[0])
	highVal, _ := strconv.Atoi(rangeval[1])
	numberOcs := strings.Count(parsed[2], string(parsed[1][0]))
	returnval := (numberOcs >= lowVal) && (numberOcs <= highVal)
	return returnval
}

func main() {
	inputlines := inputloader.ReadInput("../data/input02.txt")
	sum := 0
	for _, s := range inputlines {
		if validpw(s) {
			sum++
		}
	}
	fmt.Println(sum)
}
