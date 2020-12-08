package main

import (
	"strconv"
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
    numbers []int
}

func parseBag(inpt string) (string, []string, []int) {
    regex := regexp.MustCompile(`\w*\s\w*\sbag`)
    regex2 := regexp.MustCompile(`\d\s\w*\s\w*\sbag`)
    fields := strings.Fields(inpt)
    target := fields[0] + " " + fields[1]
    if fields[3] != "contains"{
        errors.New("failed parse")
    }
    parsed := regex.FindAllString(inpt, -1)
    parsed2 := regex2.FindAllString(inpt, -1)
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
    numbers := make([]int, len(parsed))
    for i, s := range parsed2 {
        numbers[i], _ = strconv.Atoi(string(s[0]))
    }
    if len(numbers) + 1 != len(parsed) {
        errors.New("numbers parsed wrong")
    }

    return target, parsed[1:], numbers

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


func countBags(graph map[string]*Node, key string) int {
    if len(graph[key].vertices) == 0 {
        return 1
    } else {
        sum := 0
        for i, s := range graph[key].vertices {
            sum += graph[key].numbers[i] * countBags(graph, s)
        }
        fmt.Printf("key %s, number %d\n", key, sum)
        return sum + 1
    }
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
    // inputlines := []string{
    //     "shiny gold bags contain 2 dark red bags.",
    //     "dark red bags contain 2 dark orange bags.",
    //     "dark orange bags contain 2 dark yellow bags.",
    //     "dark yellow bags contain 2 dark green bags.",
    //     "dark green bags contain 2 dark blue bags.",
    //     "dark blue bags contain 2 dark violet bags.",
    //     "dark violet bags contain no other bags.",
    // }


    graph := make(map[string]*Node)

    for _, s := range inputlines {
        target, vertices, numbers := parseBag(s)
        graph[target] = &Node{false, vertices, numbers}
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
    fmt.Println(countBags(graph, "shiny gold") - 1)
    // for k, v := range graph {
    //     fmt.Println("key: " + k)
    //     fmt.Println(*v)
    // }
}
