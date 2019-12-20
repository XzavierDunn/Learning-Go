package main

import (
    "fmt"
    "strings"
)

var rev []string

func main() {
    reverse("beans")
    end := strings.Join(rev, "")
    fmt.Println(end)
}

func reverse(x string) string {
    if len(x) > 0 {
        rev = append(rev, string(x[int(len(x)-1)]))
        x = x[0:int(len(x)-1)]
        return reverse(x)
    } 
    return x
}
