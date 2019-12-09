package main

import "fmt"

func juice(x int) {
    for x <= 200 {
        fmt.Println(x)
        x++
    }
}

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
    }
}

func main() {
    go f(0)
    fmt.Println("yes")
}
