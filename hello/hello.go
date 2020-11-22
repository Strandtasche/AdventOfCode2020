package main

import "fmt"
import "example.com/greetings"

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}