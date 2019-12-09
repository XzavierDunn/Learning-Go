package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, err := readLines("f.txt")
	if err != nil {
		fmt.Println(err)
	}

	wire1 := lines[0]
	wire2 := lines[1]

	a, b := 0, 0

	x := main2(wire1, a, b)
	y := main2(wire2, a, b)

	_, _ = x, y

}

func main2(wire string, x int, y int) []string {
	var nums []string
	for i := in wire {
		if wire[i] == 'R' {
			nums, x, y = add(i, nums, x, y)
		}
	}

	return nums
}

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
