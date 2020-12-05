package main

import (
	"strconv"
    "strings"
    "fmt"
    "inputloader"
)

var replacer = strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

func calculateBin(inpt string) int {
    number, _ := strconv.ParseInt(replacer.Replace(inpt), 2, 64)
    return int(number)
}

func main() {
    inputlines := inputloader.ReadInput("../data/input05.txt")
    // inputlines := []string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
    max := 0
    for _, s := range inputlines {
        row := calculateBin(s[:7])
        column := calculateBin(s[7:])
        id := row * 8 + column
        if id > max {
            max = id
        }
    }
    fmt.Println(max)
}
