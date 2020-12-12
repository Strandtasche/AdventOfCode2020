package main

import (
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
    inputlines := inputloader.ReadInput("../data/input12.txt")
    locX := 0
    locY := 0
    dir := 90

    for _, val := range inputlines {
        increase, err := strconv.Atoi(val[1:])
        if err != nil {
            fmt.Println("parse error!")
            return
        }
        switch instr := val[0]; instr {
        case 'N':
            locY += increase
        case 'S':
            locY -= increase
        case 'E':
            locX += increase
        case 'W':
            locX -= increase
        case 'R':
            dir += increase
            dir %= 360
        case 'L':
            dir -= increase
            dir += 360
            dir %= 360
        case 'F':
            switch k := dir; k{
            case 0:
                locY += increase
            case 90:
                locX += increase
            case 180:
                locY -= increase
            case 270:
                locX -= increase
            default:
                fmt.Println("error, F number wrong")
                return
            }
        default:
            fmt.Println("error! letter wrong")
            return
        }
        fmt.Printf("X: %d, Y: %d, dir: %d\n", locX, locY, dir )
    }

    fmt.Println(abs(locX) + abs(locY))


}
