package main

import (
	// "sort"
	"strings"
	"strconv"
    "fmt"
    "inputloader"
)

func part1 (depart int, busses []string) int {
    waittime := make([]int, len(busses))

    for it, val := range busses {
        number, err := strconv.Atoi(val)
        if err != nil {
            waittime[it] = -1
            continue
        }
        waittime[it] = number - (depart % number)
    }

    var loc int
    min := waittime[0] // assumption: first element ist not x
    for it := range waittime {
        if waittime[it] > 0 && waittime[it] < min {
            min = waittime[it]
            loc = it
        }
    }

    bus, _ := strconv.Atoi(busses[loc])
    return min * bus

}

func stupid(busses []string) int {
    i := 0
    solved := false
    iteration, _ := strconv.Atoi(busses[0])
    for {
        solved = true
        for it, val := range busses {
            number, err := strconv.Atoi(val)
            if err != nil {
                continue
            }
            if (i+it) % number != 0 {
                i += iteration
                solved = false
                break
            }

        }
        if solved {
            return i
        }
    }
}

func main() {
    inputlines := inputloader.ReadInput("../data/input13.txt")
    // inputlines := []string{"", "1789,37,47,1889"}
    // depart, _ := strconv.Atoi(inputlines[0])
    busses := strings.Split(inputlines[1], ",")
    var limitedBusses []int
    var indices []int

    for it, val := range busses {
        number, err := strconv.Atoi(val)
        if err != nil {
            continue
        }
        limitedBusses = append(limitedBusses, number)
        indices = append(indices, it)
    }

    fmt.Println(limitedBusses)
    fmt.Println(indices)
    return

    i := 0
    solved := false
    iteration, _ := strconv.Atoi(busses[0])
    for {
        solved = true
        for it, val := range busses {
            number, err := strconv.Atoi(val)
            if err != nil {
                continue
            }
            if (i+it) % number != 0 {
                i += iteration
                solved = false
                break
            }

        }
        if solved {
            fmt.Println(i)
            return
        }
    }

}
