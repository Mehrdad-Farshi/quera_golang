package main

import "fmt"

func main() {
    x := 1
    if x > 2 {
    fmt.Println(x,"is greater than 2")   
    }else if x < 0 {
    fmt.Println(x,"is negative")   
    }else{
    fmt.Println(x,"is between 0 and 2")   
    }
}
