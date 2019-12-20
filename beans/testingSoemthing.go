package main

import ("fmt")

func main() {
    x := "beans"
    fmt.Println("x", x)
    _, err := fmt.Println("x[0:len-1]", string(x[0:int(len(x)-1)]))
    y := x[0:int(len(x)-1)]
    fmt.Printf("%s", y)
    if err != nil {
        fmt.Println(err)
    }
}
