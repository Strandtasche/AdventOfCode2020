package main

import (
	"strconv"
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
        valid[splitted2[0]] = dataValidation(splitted2[0], splitted2[1])
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

func dataValidation(key string, value string) bool {
    if key == "byr" {
        year, _ := strconv.Atoi(value)
        return year >= 1920 && year <= 2002
    }
    if key == "iyr" {
        year, _ := strconv.Atoi(value)
        return year >= 2010 && year <= 2020
    }
    if key == "eyr" {
        year, _ := strconv.Atoi(value)
        return year >= 2020 && year <= 2030
    }
    if key == "hgt" {
        if value[len(value)-2:] == "cm" {
            height, _ := strconv.Atoi(value[:len(value)-2])
            return height >= 150 && height <= 193
        } else if value[len(value)-2:] == "in" {
            height, _ := strconv.Atoi(value[:len(value)-2])
            return height >= 59 && height <= 76
        } else {
            return false
        }
    }
    if key == "hcl" {
        _, err := strconv.ParseUint(value[1:], 16, 64)
        if err != nil {
            return false
        }
        return string(value[0]) == "#" && len(value) == 7
    }
    if key == "ecl" {
        return stringInSlice(value, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"})
    }
    if key == "pid" {
        _, err := strconv.Atoi(value)
        if err != nil {
            return false
        }
        return len(value) == 9
    }
    return true
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
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
    fmt.Println(validatePassport("ecl:gry pid:860033327 eyr:2020 hcl:#fffffdi byr:2003 iyr:2017 hgt:183cm"))

}
