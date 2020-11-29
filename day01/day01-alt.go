package main

import (
	"fmt"
	"inputloader"
	"strconv"
)

func getLaunchMass (mass int) int {
	return mass/3 - 2
}

func main() {
	inputlines := inputloader.ReadInput("../data/input_01-alt")
	// var inputlines [1]string
	// inputlines[0] = "100756"
	sum := 0
	for _, s := range inputlines {
		mass, _ := strconv.Atoi(s)
		for getLaunchMass(mass) > 0 {
			sum += getLaunchMass(mass)
			mass = getLaunchMass(mass)
		}
	}

	fmt.Println(sum)
}