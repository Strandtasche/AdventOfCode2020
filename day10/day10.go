package main

import (
    "fmt"
    "inputloader"
    "util"
    "sort"
)

func checkSums (num int, preamble []int) bool {
    permutations := make([]int, 25*24)
    pos := 0
    for m, i1 := range preamble {
        for n, i2 := range preamble {
            if m != n {
                permutations[pos] = i1 + i2
                pos++
            }
        }
    }
    return util.IntInSlice(num, permutations)

}


func main() {
    inputlines := inputloader.ReadInput("../data/input10.txt")
    inputlinesInt, _ := util.SliceAtoi(inputlines)
    inputlinesInt = append(inputlinesInt, 0)
    diffs := make([]int, len(inputlinesInt))
    sort.Ints(inputlinesInt)

    for it := range inputlinesInt[1:] {
        diffs[it] = inputlinesInt[1+ it] - inputlinesInt[it]
    }
    distribution := make(map[int]int)

    fmt.Println("printing diffs")
    for _, row := range diffs {
        distribution[row]++
        fmt.Println(row)
    }

    fmt.Println(distribution)

    fmt.Println(distribution[1] * (distribution[3] + 1) )

}
