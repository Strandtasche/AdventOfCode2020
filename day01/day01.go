package main

import (
	"fmt"
	"inputloader"
	"strconv"
)

func main() {
	inputlines := inputloader.ReadInput("../data/input01.txt")
	inputints := []int{}
	for _, s := range inputlines {
		cost, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			return
		}
		inputints = append(inputints, cost)
	}

	for _, j := range inputints {
		for _, l := range inputints {
			for _, m := range inputints {
				if j+l+m == 2020 {
					fmt.Printf("j: %d, l: %d, m: %d, product: %d.", j, l, m, j*l*m)
					return
				}
			}
		}
	}
}
