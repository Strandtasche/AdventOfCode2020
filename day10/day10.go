package main

import (
    "fmt"
    "inputloader"
    "util"
    "sort"
)

func isValidStep(inpt int) bool {
    return inpt == 1 || inpt == 2 || inpt == 3
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
    }

    // fmt.Println(distribution)

    // fmt.Println(distribution[1] * (distribution[3] + 1) )

    diffs = append(diffs, 3)
    diffs = append(diffs, 100)
    paths := make([]int, len(inputlinesInt))
    paths[0] = 1

    // fmt.Println(inputlinesInt)
    // fmt.Println(diffs)

    for it := range inputlinesInt {
        left := len(paths) - it
        if isValidStep(diffs[it]) && left > 1{
            paths[it+1] += paths[it]
        }
        if isValidStep(diffs[it] + diffs[it+1]) && left > 2{
            paths[it+2] += paths[it]
        }
        if isValidStep(diffs[it] + diffs[it+1] + diffs[it+2]) && left > 3{
            paths[it+3] += paths[it]
        }
        // fmt.Printf("run %d\n", it)
        // fmt.Println(paths)
    }

    fmt.Println(paths[len(paths)-1])


}
