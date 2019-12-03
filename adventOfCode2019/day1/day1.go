package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	// Uses the parameter path of type string and os to open the file if it exists
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// closes file after checking it's valid
	defer file.Close()

	// Defines variable lines as a string slice
	var lines []string

	// Creates a scanner variable using the bufio package to "scan" the file
	scanner := bufio.NewScanner(file)
	// For loop to grab all the lines in the file
	for scanner.Scan() {
		// Adding the text to the slice lines defined above
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Takes in int x and returns an int
func math(x int) int {
	// Defines variable z that is equal to the solution
	z := (x / 3) - 2
	return z
}

func main() {
	// Defines the variable used to get the final number
	num := 0

	// Calls the readlines func and gets back the lines and errors
	lines, err := readLines("day1Input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Loops over the index and values of lines
	for _, line := range lines {
		// Converts the strings in lines to integers
		i, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}
		// Defines var K and calls the math func
		k := math(i)
		// Uses the num var defined above to add up all the numbers from the math func
		num = num + k
	}

	// Returns the solution for advent of code day 1 2019
	fmt.Println("Final?", num)
}
