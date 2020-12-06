package main

import (
    "strings"
    "fmt"
    "inputloader"
    "unicode"
)

func CalculateAnyone(inpt string) int {
    inpt = SpaceMap(inpt)
    positives := make(map[rune]bool)
    for _, c := range inpt {
        positives[c] = true
    }
    return len(positives)
}

func SpaceMap(str string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsSpace(r) {
            return -1
        }
        return r
    }, str)
}

func main() {
    inputlines := inputloader.ReadInputByEmptyLine("../data/input06.txt")
    // inputlines := []string{"abc", "c c c c a", "ace abce"}
    sum := 0
    for _, s := range inputlines {
        sum += CalculateAnyone(s)
        // fmt.Println(Calculateyes(s))
    }
    fmt.Println(sum)
}
