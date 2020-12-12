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

func rotate(degrees int, x int, y int) (int, int) {
    // radian := float64(degrees) * (180 / math.Pi)
    // tmpWaypointDiffX := float64(x) * math.Cos(radian) - float64(y) * math.Sin(radian)
    // tmpWaypointDiffY := float64(x) * math.Sin(radian) + float64(y) * math.Cos(radian)
    // xNew := int(math.Round(tmpWaypointDiffX))
    // yNew := int(math.Round(tmpWaypointDiffY))
    switch rot := (degrees + 360) % 360; rot{
    case 0:
        return x, y
    case 90:
        return y, -x
    case 180:
        return -x, -y
    case 270:
        return -y, x
    default:
        return 0, 0
    }
}

func main() {
    inputlines := inputloader.ReadInput("../data/input12.txt")
    shiplocX := 0
    shiplocY := 0
    waypointDiffX := 10
    waypointDiffY := 1

    for _, val := range inputlines {
        increase, err := strconv.Atoi(val[1:])
        if err != nil {
            fmt.Println("parse error!")
            return
        }
        switch instr := val[0]; instr {
        case 'N':
            waypointDiffY += increase
        case 'S':
            waypointDiffY -= increase
        case 'E':
            waypointDiffX += increase
        case 'W':
            waypointDiffX -= increase
        case 'R':
            waypointDiffX, waypointDiffY = rotate(increase, waypointDiffX, waypointDiffY)
        case 'L':
            waypointDiffX, waypointDiffY = rotate(-increase, waypointDiffX, waypointDiffY)
        case 'F':
            shiplocX += increase * waypointDiffX
            shiplocY += increase * waypointDiffY
        default:
            fmt.Println("error! letter wrong")
            return
        }
        fmt.Printf("X: %d, Y: %d, dir: %d, %d\n", shiplocX, shiplocY, waypointDiffX, waypointDiffY)
    }

    fmt.Println(abs(shiplocX) + abs(shiplocY))


}
