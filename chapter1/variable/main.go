package main

import "fmt"

func main() {
name := "ahmad"
    switch name {
        case "ahmad":
        fmt.Println("Hello Ahmad")
        fallthrough
        case "mohammad":
        fmt.Println("Greetings")
        default:
        fmt.Println("How are you?")
    }
}
