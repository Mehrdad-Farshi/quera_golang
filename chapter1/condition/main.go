package main

import "fmt"

func main() {
name := 10
    switch  {
        case name >12:
        fmt.Println("Hello Ahmad")
        fallthrough
        case name >11:
        fmt.Println("Greetings")
        default:
        fmt.Println("How are you?")
    }
}
