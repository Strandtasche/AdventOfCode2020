package main

import (
	"strings"
	"strconv"
    "fmt"
    "inputloader"
)


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
    inputlines := inputloader.ReadInput("../data/input13.txt")
    depart, _ := strconv.Atoi(inputlines[0])
    busses := strings.Split(inputlines[1], ",")

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
    fmt.Println(min * bus)

}
