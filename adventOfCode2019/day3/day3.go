package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "strconv"
)

func readLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}

func main2(wire string, x int, y int) []string {
	var nums []string
    fuck := strings.Split(wire, ",")
	for i, _ := range fuck {
        yes := string(fuck[i])
        yes = string(yes[0])
        sub := false
		if yes == "R" {
			nums, x, y = add(fuck[i], nums, x, y, sub)
        } else if yes == "L" {
            sub = true
            nums, x, y = add(fuck[i], nums, x, y, sub)
        } else if yes == "U" {
            nums, x, y = subtract(fuck[i], nums, x, y, sub)
        } else if yes == "D" {
            sub = true
            nums, x, y = subtract(fuck[i], nums, x, y, sub)
        } else {
            continue
        }
	}

	return nums
}

func add(i string, nums []string, x int, y int, sub bool) ([]string, int, int) {
    intVal, err := strconv.Atoi(i[1:])
    if err != nil {
        fmt.Println(err) 
    }
    for b := 0; b <= intVal; b++ {
        if sub == true {
            x = x - 1
            nums = append(nums, x, y)
        } else {
            x++
            nums = append(nums, x, y)
        }
    }

    return nums, x, y
}

func subtract(i string, nums []string, x int, y int, sub bool) ([]string, int, int) {
        return nums, x, y
}


func main() {
	lines, err := readLines("f.txt")
	if err != nil {
		fmt.Println(err)
	}
	wire1 := string(lines[0])
	wire2 := string(lines[1])

	a, b := 0, 0

	x := main2(wire1, a, b)
	y := main2(wire2, a, b)

	_, _ = x, y

}

