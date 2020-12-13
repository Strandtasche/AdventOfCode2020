package main

import (
	"strings"
    "fmt"
    "inputloader"
    // "util"
    // "sort"
)

type Field struct {
    val []byte
    len_row int
    len_col int
}

func iterateGame(field Field) Field {
    newfieldbyte := make([]byte, len(field.val))
    copy(newfieldbyte, field.val)
    newField := Field{newfieldbyte, field.len_row, field.len_col}

    for it := 0; it < field.len_row; it++ {
        for it2 := 0; it2 < field.len_col; it2++ {
            newField.val[it * newField.len_col + it2] = singleChair(getArea(field, it, it2))
        }
    }
    return newField
}

func singleChair(inpt []byte) byte {
    if inpt[4] == '.' {
        return '.'
    }
    occ := 0
    for i := range inpt {
        if i != 4 {
            if inpt[i] == '#' {
                occ++
            }
        }
    }

    if inpt[4] == 'L' && occ == 0 {
        return '#'
    }
    if inpt[4] == '#' && occ >= 4 {
        return 'L'
    }
    return inpt[4]
}

func getArea(field Field, row int, col int) []byte {
    area := []byte{getChar(field, row-1, col-1), getChar(field, row-1, col), getChar(field, row-1, col+1),
                   getChar(field, row, col-1), getChar(field, row, col), getChar(field, row, col+1),
                   getChar(field, row+1, col-1), getChar(field, row+1, col), getChar(field, row+1, col+1)}
    return area
}

func getChar(field Field, row int, col int) byte {
    if row >= field.len_row || col >= field.len_col || row < 0 || col < 0 {
        return '.'
    }
    return field.val[row * field.len_col + col]
}

func prettyPrint(field Field) {
    for i:= 0; i < field.len_row; i++ {
        fmt.Println(string(field.val[i*field.len_col:(i+1)*field.len_col]))
    }

}



func main() {
    inputlines := inputloader.ReadInput("../data/input11_alt.txt")
    // inputlines := []string{"LLL","LLL","LLL"}
    seatingString := []byte(strings.Join(inputlines, ""))
    width := len(inputlines[0])
    height := len(inputlines)


    for i := 0; i < width; i++ {
        for j := 0; i < height; j++ {
            for
        }
    }

    fmt.Println("done")

    count := 0
    fmt.Println(count)
}
