package main

import (
    "strings"
    "fmt"
    "inputloader"
)

func validatePassport(inpt string) bool {
    splitted := strings.Fields(inpt)
    valid  := make(map[string]bool)
    var splitted2 []string
    for _, s := range splitted {
        splitted2 = strings.Split(s, ":")
        valid[splitted2[0]] = true
    }
    // fmt.Printf("len map: %d \n", len(valid))
    return allTrue(valid, []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"})
}

func allTrue(inptmap map[string]bool, keys []string) bool {
    for _, s := range keys {
        if !inptmap[s] {
            return false
        }
    }
    return true
}

func main() {
    inputlines := inputloader.ReadInputByEmptyLine("../data/input04.txt")
    // fmt.Println(inputlines[0)
    sum := 0
    for _, s := range inputlines{
        if validatePassport(s) {
            sum++
        }
    }
    fmt.Println(sum)
    // fmt.Println(inputlines[0])
    // fmt.Println(validatePassport(inputlines[3]))
    // fmt.Println(validatePassport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffdi byr:1937 iyr:2017 hgt:183cm"))

}
