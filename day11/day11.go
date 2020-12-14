package main

import (
	"strings"
    "fmt"
    "inputloader"
    // "util"
    // "sort"
)

var field string
var width int
var height int

func prettyPrint(field []byte) {
    for i:= 0; i < height; i++ {
        fmt.Println(string(field[i*width:(i+1)*width]))
    }
}

func charValid(row int, col int) bool {
    return row < width && col < height && row >= 0 && col >= 0
}


func main() {
    inputlines := inputloader.ReadInput("../data/input11.txt")
    // inputlines := []string{"LLL","LLL","LLL"}
    seatingString := []byte(strings.Join(inputlines, ""))
    newSeating := make([]byte, len(seatingString))
    width = len(inputlines[0])
    height = len(inputlines)
    matrix := []int{-1, 0, 1}
    var nocc int

    for {
        for i := 0; i < width; i++ {
            for j := 0; j < height; j++ {
                if seatingString[i + j*width] == '.' {
                    newSeating[i + j * width] = '.'
                    continue
                }
                nocc = 0
                for _, up := range matrix {
                    for _, side := range matrix {
                        if !(up == 0 && side == 0) && charValid(i + up, j + side) {
                            if seatingString[i+up + (j + side) * width] == '#' {
                                nocc++
                            }
                        }
                    }
                }
                if seatingString[i + j * width] == 'L' {
                    if nocc == 0 {
                        newSeating[i + j * width] = '#'
                    } else {
                        newSeating[i + j * width] = 'L'
                    }

                } else if seatingString[i + j * width] == '#' {
                    if nocc >= 4 {
                        newSeating[i + j * width] = 'L'
                    } else {
                        newSeating[i + j * width] = '#'
                    }
                }
            }
        }

        if string(seatingString) == string(newSeating) {
            fmt.Println("found it!")
            prettyPrint(newSeating)
            break
        } else {
            copy(seatingString, newSeating)
            newSeating = make([]byte, len(seatingString))
        }
    }


    count := 0
    for _, val := range newSeating{
        if val == '#'{
            count++
        }
    }
    fmt.Println(count)
}
