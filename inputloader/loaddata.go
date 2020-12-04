package inputloader

import (
    "bufio"
    "log"
    "os"
)

func ReadInput(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var lines []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return lines
}

func ReadInputByEmptyLine(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var lines []string
    var templine string
    var passport string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        passport = scanner.Text()
        if len(passport) == 0 {
            lines = append(lines, templine)
            templine = ""
        } else {
            templine += " " + passport
        }
    }
    lines = append(lines, templine)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return lines
}
