package main

import (
	"strconv"
	"strings"
    "fmt"
    "inputloader"
    // "util"
)


func main() {
    inputlines := inputloader.ReadInput("../data/input08.txt")
    // inputlines := inputloader.ReadInput("../data/input07.txt")
    sum := 0
    i := 0
    for  inputlines[i] != "visited" {
        fields := strings.Fields(inputlines[i])
        inputlines[i] = "visited"
        if fields[0] == "acc" {
            acc, _ := strconv.Atoi(fields[1])
            sum += acc
            i++
        } else if fields[0] == "nop" {
            i++
        } else if fields[0] == "jmp" {
            jump, _ := strconv.Atoi(fields[1])
            i += jump
        }
    }


    fmt.Println("total:")
    fmt.Println(sum)
}
