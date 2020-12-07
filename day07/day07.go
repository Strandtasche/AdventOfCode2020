package main

import (
	"strings"
    "fmt"
    "inputloader"
    // "util"
    "errors"
    "regexp"
)

// Node represents node in graph
type Node struct {
    // this is a comment
    containsMine bool
    vertices []string
}

func parseBag(inpt string) (string, []string) {
    regex := regexp.MustCompile(`\w*\s\w*\sbag`)
    fields := strings.Fields(inpt)
    target := fields[0] + " " + fields[1]
    if fields[3] != "contains"{
        errors.New("failed parse")
    }
    parsed := regex.FindAllString(inpt, -1)
    if target != parsed[0][:len(parsed[0])-3] {
        errors.New("failed parse")
    }

    // remove "no other bag" from parsed
    for i, v := range parsed {
        if v == "no other bag" {
            parsed = append(parsed[:i], parsed[i+1:]...)
            break
        }
    }

    for i, s := range parsed {
        parsed[i] = s[:len(s)-4]
    }
    return target, parsed[1:]

}

func breadthFirstSearch(graph map[string]*Node, children []string) bool {
    for _, k := range children {
        if graph[k].containsMine {
            return true
        }
    }
    return false
}

func countColors(graph map[string]*Node) int {
    counter := 0
    for _, v := range graph {
        if v.containsMine {
            counter++
        }
    }
    return counter
}

func main() {
    inputlines := inputloader.ReadInput("../data/input07.txt")
    // inputlines := []string{"light red bags contain 1 bright white bag, 2 muted yellow bags.",
    //                         "dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
    //                         "bright white bags contain 1 shiny gold bag.",
    //                         "muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
    //                         "shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
    //                         "dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
    //                         "vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
    //                         "faded blue bags contain no other bags.",
    //                         "dotted black bags contain no other bags."}

    graph := make(map[string]*Node)

    for _, s := range inputlines {
        target, vertices := parseBag(s)
        graph[target] = &Node{false, vertices}
    }

    graph["shiny gold"].containsMine = true

    changed := 1
    for changed > 0 {
        changed = 0
        for _, v := range graph {
            if !v.containsMine && breadthFirstSearch(graph, v.vertices) {
                v.containsMine = true
                changed++
            }
        }
    }


    fmt.Println("total:")
    fmt.Println(countColors(graph) - 1)
    // for k, v := range graph {
    //     fmt.Println("key: " + k)
    //     fmt.Println(*v)
    // }
}
