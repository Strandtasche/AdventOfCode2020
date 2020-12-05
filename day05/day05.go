package main

import (
	"strconv"
    "strings"
    "fmt"
    "inputloader"
    "sort"
)

var replacer = strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

func calculateBin(inpt string) int {
    number, _ := strconv.ParseInt(replacer.Replace(inpt), 2, 64)
    return int(number)
}

func main() {
    inputlines := inputloader.ReadInput("../data/input05.txt")
    // inputlines := []string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
    var seats []int
    for _, s := range inputlines {
        row := calculateBin(s[:7])
        column := calculateBin(s[7:])
        id := row * 8 + column
        seats = append(seats, id)
    }
    sort.Ints(seats)
    for i := range seats {
			if seats[i] + 1 != seats[i+1] {
                fmt.Println(seats[i] + 1)
                return
            }
    }
}
