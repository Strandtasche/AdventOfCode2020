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
    // inputlines := inputloader.ReadInput("../data/input08_alt.txt")
    inputlinesAlt := make([]string, len(inputlines))
    copy(inputlinesAlt, inputlines)

    for it := range inputlinesAlt {
        copy(inputlinesAlt, inputlines)
        if string(inputlinesAlt[it][0]) == "j" {
            // fmt.Println("replaced jump")
            inputlinesAlt[it] = strings.Replace(inputlinesAlt[it], "jmp", "nop", 1)
        } else if string(inputlinesAlt[it][0]) == "n"{
            // fmt.Println("replaced nop")
            inputlinesAlt[it] = strings.Replace(inputlinesAlt[it], "nop", "jmp", 1)
        } else {
            continue
        }

        sum := 0
        i := 0
        for  inputlinesAlt[i] != "visited" {
            fields := strings.Fields(inputlinesAlt[i])
            inputlinesAlt[i] = "visited"
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
            if i >= len(inputlines){
                fmt.Printf("finished: %d", sum)
                return
            }
        }
        // fmt.Printf("canceled due to visited with %d\n", sum)
        // fmt.Println(inputlinesAlt)
    }



    fmt.Println("total:")
}
