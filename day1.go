package main 
import (
    "fmt"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    final := make([]int, 0)

    dat, err := ioutil.ReadFile("day1Input.txt")
    check(err)

    wtf := string(dat).split("\n")

    for _, i := range wtf {
        fmt.Println(string(i))
        final = append(final, math(int(i)))
    }


}

func math(x int) int {
    z := (x / 3) - 2
    return z
}
