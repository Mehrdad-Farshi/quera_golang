package main

import "fmt"

func main() {
        var target int = 5
        sum := 0
        for i := 1; i <= target; i++ {
        fmt.Println("i is  =  ",i)
        fmt.Println("sum before sum+=i =  ",sum)
            sum += i
        fmt.Println("sum after sum+=i =  ",sum)
        }
        fmt.Println("sum eqauls to ",sum)
}
