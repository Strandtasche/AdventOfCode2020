package main

import (
	"fmt"
	"inputloader"
)

func treefinder(row int, code string) bool {
	return code[(row*3)%len(code)] == '#'
}

func treefinderPro(row int, code string, x int) bool {
	return code[(row*x)%len(code)] == '#'
}

func main() {
	inputlines := inputloader.ReadInput("../data/input03.txt")
	// inputlines := []string{"....", "....", "....", "...."}

	sums := make([]int, 5)
	for i, s := range inputlines {
		if treefinderPro(i, s, 1) {
			sums[0]++
		}
		if treefinderPro(i, s, 3) {
			sums[1]++
		}
		if treefinderPro(i, s, 5) {
			sums[2]++
		}
		if treefinderPro(i, s, 7) {
			sums[3]++
		}
		if i%2 == 0 && treefinderPro(i/2, s, 1) {
			sums[4]++
		}
	}
	prod := 1
	for _, x := range sums {
		prod *= x
	}
	fmt.Println(sums)
	fmt.Println(prod)
}
