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
    inputlines := inputloader.ReadInput("../data/input09.txt")
    inputlinesInt, _ := util.SliceAtoi(inputlines)
    var target int
    // target = 127

    for it, val := range inputlinesInt[25:] {
        if !checkSums(val, inputlinesInt[it:it+25]){
            fmt.Println(val)
            target = val
            break
        }
    }
    var tmpSafe []int

    for it, val := range inputlinesInt {
        sum := val
        tmpSafe = []int{val}
        j := 1
        for sum < target {
            sum += inputlinesInt[it + j]
            tmpSafe = append(tmpSafe, inputlinesInt[it + j])
            if sum == target {
                sort.Ints(tmpSafe)
                fmt.Println(tmpSafe[0] + tmpSafe[len(tmpSafe) - 1])
                return
            }
            j++
        }
    }
}
