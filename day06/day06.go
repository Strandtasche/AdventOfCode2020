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

func CalculateEveryone(inpt string) int {
    individuals := strings.Fields(inpt)
    positives := make(map[rune]bool)
    for _, c := range SpaceMap(individuals[0]) {
        positives[c] = true
    }

    for _, indv := range individuals {
        tmpPos := make(map[rune]bool)
        for _, c := range indv {
            tmpPos[c] = true
        }
        intersection := make(map[rune]bool)
        for k := range positives {
            if tmpPos[k] {
                intersection[k] = true
            }
        }
        positives = intersection
    }
    return len(positives)
}

func main() {
    inputlines := inputloader.ReadInputByEmptyLine("../data/input06.txt")
    // inputlines := []string{"abc", "c c c c a", "ace abce"}
    sum := 0
    for _, s := range inputlines {
        sum += CalculateEveryone(s)
        // fmt.Println(CalculateEveryone(s))
    }
    fmt.Println("total:")
    fmt.Println(sum)
}
